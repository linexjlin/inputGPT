package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
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

		tmpTxt := ""
		assistantAns := ""
		lastPaste := time.Now()
		for txt := range txtChan {
			tmpTxt += txt
			assistantAns += txt
			if time.Now().Sub(lastPaste).Seconds() > 1 && tmpTxt != "" {
				clipboard.WriteAll(tmpTxt)
				time.Sleep(time.Millisecond * 200)
				pressPaste()
				lastPaste = time.Now()
				tmpTxt = ""
			}
		}
		if tmpTxt != "" {
			clipboard.WriteAll(tmpTxt)
			time.Sleep(time.Millisecond * 200)
			pressPaste()
		}
		fmt.Println(assistantAns)
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
