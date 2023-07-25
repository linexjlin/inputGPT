package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"

	hook "github.com/robotn/gohook"
)

type UserSetting struct {
	model        string
	maxConext    int
	headMessages []gpt.ChatCompletionRequestMessage
	histMessages []gpt.ChatCompletionRequestMessage
}

var g_userSetting UserSetting

func initUserSetting() {
	g_userSetting.model = "gpt-3.5-turbo-0613"
	g_userSetting.maxConext = getMaxContext()
	g_userSetting.headMessages = []gpt.ChatCompletionRequestMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant!",
		},
	}
}

func registerHotKeys() {
	gptHotkeys := getGPTHotkeys()
	lastHit := time.Now()
	fmt.Printf("--- Please press %s to auto generate text --- \n", gptHotkeys)
	/*
		headMessages := []gpt.ChatCompletionRequestMessage{
			{
				Role:    "system",
				Content: "You are a helpful assistant!",
			},
		}*/
	initUserSetting()
	//histMessages := []gpt.ChatCompletionRequestMessage{}

	hook.Register(hook.KeyDown, gptHotkeys, func(e hook.Event) {
		fmt.Println(gptHotkeys)
		if time.Now().Sub(lastHit).Seconds() > 1.0 {
			lastHit = time.Now()
		} else {
			return
		}

		clipboardContent, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Failed to read clipboard content:", err)
		}
		fmt.Println("### user:")
		fmt.Println(clipboardContent)
		messages := []gpt.ChatCompletionRequestMessage{}

		g_userSetting.histMessages = append(g_userSetting.histMessages, gpt.ChatCompletionRequestMessage{
			Role:    "user",
			Content: clipboardContent,
		})
		msgIdx := 0
		if len(g_userSetting.histMessages)-g_userSetting.maxConext > 0 {
			msgIdx = len(g_userSetting.histMessages) - g_userSetting.maxConext
		}
		var txtChan = make(chan string, 100)
		messages = append(messages, g_userSetting.headMessages...)
		go queryGTP(txtChan, append(messages, g_userSetting.histMessages[msgIdx:]...))

		assistantAns := ""
		fmt.Print("### Assistant:\n")
		for txt := range txtChan {
			robotgo.TypeStr(txt)
			fmt.Print(txt)
			assistantAns += txt
		}
		fmt.Print("\n")
		g_userSetting.histMessages = append(g_userSetting.histMessages, gpt.ChatCompletionRequestMessage{
			Role:    "assistant",
			Content: assistantAns,
		})
	})

	s := hook.Start()
	<-hook.Process(s)
}

func main() {
	go registerHotKeys()
	systray.Run(onReady, onExit)
}
