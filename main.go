package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"
	"github.com/joho/godotenv"
	hook "github.com/robotn/gohook"
)

type UserSetting struct {
	mask         string
	model        string
	maxConext    int
	headMessages []gpt.ChatCompletionRequestMessage
	histMessages []gpt.ChatCompletionRequestMessage
}

var g_userSetting UserSetting

func initUserSetting() {
	g_userSetting.mask = "Default"
	g_userSetting.model = "gpt-3.5-turbo-0613"
	g_userSetting.maxConext = getMaxContext()
	g_userSetting.headMessages = []gpt.ChatCompletionRequestMessage{
		{
			Role:    "system",
			Content: "Just complete the text I give you, do not explain.",
		},
	}
}

func registerHotKeys() {
	var txtChan chan string
	ctx, cancel := context.WithCancel(context.Background())

	gptHotkeys := getGPTHotkeys()
	lastHit := time.Now()
	fmt.Printf("--- Please press %s to auto generate text --- \n", gptHotkeys)
	initUserSetting()

	hook.Register(hook.KeyDown, gptHotkeys, func(e hook.Event) {
		go func() {
			fmt.Println(gptHotkeys)
			if time.Now().Sub(lastHit).Seconds() > 1.0 {
				lastHit = time.Now()
			} else {
				return
			}

			clipboardContent, err := clipboard.ReadAll()
			if err != nil {
				fmt.Println("Failed to read clipboard content:", err)
				return
			}

			if len(clipboardContent) < 1 {
				fmt.Println("Empty question")
				return
			}

			fmt.Println("### prompt:", g_userSetting.mask)
			fmt.Println("### user:")
			fmt.Println(clipboardContent)
			messages := []gpt.ChatCompletionRequestMessage{}

			msgIdx := 0
			if len(g_userSetting.histMessages)-g_userSetting.maxConext*2 > 0 {
				msgIdx = len(g_userSetting.histMessages) - g_userSetting.maxConext*2
			}

			txtChan = make(chan string, 100)
			messages = append(messages, g_userSetting.headMessages...)
			messages = append(messages, g_userSetting.histMessages[msgIdx:]...)
			messages = append(messages, gpt.ChatCompletionRequestMessage{
				Role:    "user",
				Content: clipboardContent,
			})

			ctx, cancel = context.WithCancel(context.Background())
			go queryGTP(ctx, txtChan, messages)

			//isCancel := false
			assistantAns := ""
			fmt.Print("### Assistant:\n")
			for {
				select {
				case txt, ok := <-txtChan:
					if !ok {
						// txtChan is closed, exit the loop
						//fmt.Println("complete")
						fmt.Print("\n")
						g_userSetting.histMessages = append(g_userSetting.histMessages, gpt.ChatCompletionRequestMessage{
							Role:    "user",
							Content: clipboardContent,
						})

						g_userSetting.histMessages = append(g_userSetting.histMessages, gpt.ChatCompletionRequestMessage{
							Role:    "assistant",
							Content: assistantAns,
						})
						updateClearContextTitle(len(g_userSetting.histMessages) / 2)
						return
					}
					fmt.Print(txt)
					for i, t := range strings.Split(txt, "\n") {
						if i > 0 {
							robotgo.KeyTap("enter")
						}
						if len(t) > 0 {
							robotgo.TypeStr(t)
						}
					}
					assistantAns += txt
				case <-ctx.Done():
					// ctx is done, exit the loop
					return
				}
			}
		}()
	})

	escCnt := 0
	lastEscHit := time.Now()
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("esc")

		if time.Now().Sub(lastEscHit).Milliseconds() < 500 {
			escCnt++
			fmt.Println("increase escCnt to", escCnt)
			if escCnt == 2 { //triple 'esc' click for quick clean context
				clearContext()
				escCnt = 0
			}
		} else {
			escCnt = 0
		}
		lastEscHit = time.Now()
		fmt.Println("esc")
		go func() {
			cancel()
		}()
	})

	s := hook.Start()
	<-hook.Process(s)
}

func main() {
	godotenv.Load("env.txt")
	setLang()
	OSDepCheck()
	go registerHotKeys()
	systray.Run(onReady, onExit)
}
