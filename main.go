package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"

	hook "github.com/robotn/gohook"
)

func registerHotKeys() {
	maxConext := 4
	gptHotkeys := getGPTHotkeys()
	lastHit := time.Now()
	fmt.Printf("--- Please press %s to auto generate text --- \n", gptHotkeys)
	headMessages := []gpt.ChatCompletionRequestMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant!",
		},
	}

	histMessages := []gpt.ChatCompletionRequestMessage{}

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
		fmt.Println("Clipboard content:", clipboardContent)
		//robotgo.TypeStr("Hello World")
		//robotgo.KeyTap("enter")
		//robotgo.KeyTap("v", "control")
		messages := []gpt.ChatCompletionRequestMessage{}

		histMessages = append(histMessages, gpt.ChatCompletionRequestMessage{
			Role:    "user",
			Content: clipboardContent,
		})
		msgIdx := 0
		if len(histMessages)-maxConext > 0 {
			msgIdx = len(histMessages) - maxConext
		}
		var txtChan = make(chan string, 100)
		messages = append(messages, headMessages...)
		go queryGTP(txtChan, append(messages, histMessages[msgIdx:]...))

		assistantAns := ""
		for txt := range txtChan {
			robotgo.TypeStr(txt)
			fmt.Print(txt)
			assistantAns += txt
		}
		histMessages = append(histMessages, gpt.ChatCompletionRequestMessage{
			Role:    "assistant",
			Content: assistantAns,
		})
	})

	s := hook.Start()
	<-hook.Process(s)
}

func main() {
	registerHotKeys()
}
