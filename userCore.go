package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hanyuancheung/gpt-go"
)

type UserCore struct {
	mask           string
	model          string
	maxConext      int
	msgCnt         int
	headMessages   []gpt.ChatCompletionRequestMessage
	histMessages   []gpt.ChatCompletionRequestMessage
	setContextMenu func(string)
}

func (u *UserCore) initUserCore() {
	u.mask = "Default"
	u.model = "gpt-3.5-turbo"
	u.maxConext = getMaxContext()
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

func (u *UserCore) reloadMask() {
	if u.mask == "Default" {
		return
	}
	filepath := fmt.Sprintf("prompts/%s.json", u.mask)
	if p, e := loadModePrompt(filepath); e != nil {
		fmt.Println(e)
	} else {
		u.SetModePrompt(p)
	}
}

func (u *UserCore) SetModePrompt(p ModePrompt) {
	u.initUserCore()
	u.headMessages = p.HeadMessages
	if p.Model != "" {
		u.model = p.Model
	}

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
}

func (u *UserCore) updateContextMenu() {
	if u.setContextMenu == nil {
		return
	}
	if u.maxConext == 0 {
		u.setContextMenu(UText("Clear Context"))
	} else {
		u.setContextMenu(fmt.Sprintf(UText("Clear Context %d/%d"), u.msgCnt, u.maxConext))
	}
}

func (u *UserCore) ClearContext() {
	fmt.Println("clean all context")
	u.histMessages = u.histMessages[:0]
	u.msgCnt = 0
	u.updateContextMenu()
}

func (u *UserCore) QueryGPT(ctx context.Context, txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	fmt.Println("Query messages:")
	showAsJson(messages)
	client := gpt.NewClient(
		getOpenAIkey(),
		gpt.WithBaseURL(getOpenAIBaseUrl()),
		gpt.WithTimeout(600*time.Second),
	)

	err := client.ChatCompletionStream(ctx, &gpt.ChatCompletionRequest{
		Model:    u.model,
		Messages: messages,
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
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		txtChan <- fmt.Sprintf("ChatCompletionStream error: %v\n", err)
	}
	close(txtChan)
}
