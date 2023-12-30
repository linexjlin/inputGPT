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

func registerHotKeys() {
	var txtChan chan string
	ctx, cancel := context.WithCancel(context.Background())

	gptHotkeys := getGPTHotkeys()
	lastHit := time.Now()
	fmt.Printf("--- Please press %s to auto generate text --- \n", gptHotkeys)
	g_userCore.initUserCore()

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

			g_userCore.reloadMask()

			fmt.Println("### prompt:", g_userCore.mask)
			fmt.Println("### user:")
			fmt.Println(clipboardContent)
			prompts, new := g_userCore.GeneratePromptMessages(clipboardContent)

			if len(prompts) == 0 {
				return
			}

			txtChan = make(chan string, 1024)

			ctx, cancel = context.WithCancel(context.Background())
			go queryGPT(ctx, txtChan, prompts)

			assistantAns := ""
			fmt.Print("### Assistant:\n")
			for {
				select {
				case txt, ok := <-txtChan:
					if !ok {
						// txtChan is closed, exit the loop
						fmt.Print("\n")
						new = append(new, gpt.ChatCompletionRequestMessage{
							Role:    "assistant",
							Content: assistantAns,
						})
						g_userCore.AddNewMessages(new)
						return
					}
					fmt.Print(txt)
					txt = strings.ReplaceAll(txt, "\r\n", "\n")
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
				g_userCore.ClearContext()
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
	g_languages.Load()
	OSDepCheck()
	go registerHotKeys()
	systray.Run(onReady, onExit)
}
