package main

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/hanyuancheung/gpt-go"
)

type UserCore struct {
	mask            string
	maskModel       string
	models          []string
	mutiModel       bool
	temperature     float32
	maskTemperature float32
	defaultModel    string
	maxConext       int
	msgCnt          int
	headMessages    []gpt.ChatCompletionRequestMessage
	histMessages    []gpt.ChatCompletionRequestMessage
	setContextMenu  func(string)
}

func (u *UserCore) initUserCore() {
	u.mask = "Default"
	u.maskModel = ""
	u.mutiModel = getMutiModel()
	u.maxConext = getMaxContext()
	u.temperature = getTemperature()
	u.msgCnt = 0
	u.histMessages = []gpt.ChatCompletionRequestMessage{}
	u.headMessages = []gpt.ChatCompletionRequestMessage{
		{
			Role:    "system",
			Content: "Just complete the text I give you, do not explain.",
		},
	}
	u.updateContextMenu()
}

func (u *UserCore) reloadMask_del() (*ModelPrompt, error) {
	fmt.Println("mask", u.mask)
	if u.mask == "Default" {
		return nil, nil
	}
	filepath := fmt.Sprintf("prompts/%s.json", u.mask)
	fmt.Println("file path:", filepath)
	if p, e := loadModelPrompt(filepath); e != nil {
		fmt.Println(e)
		return nil, e
	} else {
		u.SetModelPrompt(p)
		return &p, nil
	}
}

func (u *UserCore) SetModelPrompt(p ModelPrompt) {
	fmt.Println("mask", u.mask)
	if p.Model != "" {
		u.maskModel = p.Model
	}
	u.maskTemperature = p.Temperature
	u.headMessages = p.HeadMessages

	if p.MaxContext != 0 {
		u.maxConext = p.MaxContext
	}
	u.ClearContext()
}

func (u *UserCore) SetMask(m string) {
	u.mask = m
}

func (u *UserCore) AddNewMessages(msgs []gpt.ChatCompletionRequestMessage) {
	u.msgCnt++
	u.histMessages = append(u.histMessages, msgs...)
	u.updateContextMenu()
}

func (u *UserCore) generateZeroContextPromptMessages(newQuery string) []gpt.ChatCompletionRequestMessage {
	if len(u.headMessages) == 0 {
		return []gpt.ChatCompletionRequestMessage{{
			Role:    "user",
			Content: newQuery,
		}}
	} else {
		renderedMessages := renderMessages(u.headMessages, newQuery)
		if renderedMessages[len(renderedMessages)-1].Role == "user" {
			return renderedMessages
		} else {
			return append(renderedMessages, gpt.ChatCompletionRequestMessage{
				Role:    "user",
				Content: newQuery,
			})
		}
	}
}

func (u *UserCore) GeneratePromptMessages(newQuery string) (prompts, new []gpt.ChatCompletionRequestMessage) {
	if u.maxConext == 0 {
		prompts = u.generateZeroContextPromptMessages(newQuery)
		return prompts, prompts
	} else {
		if len(u.histMessages) == 0 {
			prompts = u.generateZeroContextPromptMessages(newQuery)
			return prompts, prompts
		} else {
			new = append(new, gpt.ChatCompletionRequestMessage{
				Role:    "user",
				Content: newQuery,
			})
			if u.msgCnt < u.maxConext {
				prompts = append(prompts, u.histMessages...)
				prompts = append(prompts, new...)
				return prompts, new
			} else {
				prompts = append(prompts, u.histMessages[:len(u.headMessages)]...)
				if u.histMessages[len(u.headMessages)].Role == "assistant" {
					prompts = append(prompts, u.histMessages[len(u.headMessages)])
					prompts = append(prompts, u.histMessages[(len(u.histMessages)-(u.maxConext-1)*2):len(u.histMessages)]...)
				} else {
					prompts = append(prompts, u.histMessages[(len(u.histMessages)-u.maxConext*2):len(u.histMessages)]...)
				}
				prompts = append(prompts, new...)
				return prompts, new
			}
		}
	}
}

func (u *UserCore) AddSetContextMenuFunc(f func(string)) {
	u.setContextMenu = f
	u.updateContextMenu()
}

func (u *UserCore) updateContextMenu() {
	if u.setContextMenu == nil {
		return
	}
	u.setContextMenu(fmt.Sprintf(UText("Clear Context ")+"%s (%d/%d)", strings.Join(u.models, ","), u.msgCnt, u.maxConext))
}

func (u *UserCore) ClearContext() {
	fmt.Println("clean all context")
	u.histMessages = u.histMessages[:0]
	u.msgCnt = 0
	u.updateContextMenu()
}

func (u *UserCore) SetDefaultModel(model string) {
	fmt.Println("Set default mode to", model)
	u.defaultModel = model
	u.updateContextMenu()
}

func (u *UserCore) SetModels(models []string) {
	fmt.Println("Set modes to", models)
	u.models = models
	u.updateContextMenu()
}

func (u *UserCore) QueryGPT(ctx context.Context, model string, temperature float32, txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	fmt.Println("Query messages:")
	showAsJson(messages)
	client := gpt.NewClient(
		getOpenAIkey(),
		gpt.WithBaseURL(getOpenAIBaseUrl()),
		gpt.WithTimeout(600*time.Second),
	)
	err := client.ChatCompletionStream(ctx, &gpt.ChatCompletionRequest{
		Model:       model,
		Temperature: temperature,
		Messages:    messages,
	}, func(response *gpt.ChatCompletionStreamResponse) {
		//fmt.Println(response.Choices)
		//fmt.Printf("%+v\n", response)
		if len(response.Choices) > 0 {
			if response.Choices[0].Delta.Content != "" {
				txtChan <- response.Choices[0].Delta.Content
			}
			if response.Choices[0].FinishReason == "stop" {
				fmt.Println("stop")
			}
		}
	})

	if err != nil && err != io.EOF {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		txtChan <- fmt.Sprintf("ChatCompletionStream error: %v\n", err)
	}
	close(txtChan)
}
