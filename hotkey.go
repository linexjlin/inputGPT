package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"
	hook "github.com/robotn/gohook"
)

func registerHotKeys(userCore *UserCore, st *SysTray) {
	var txtChan chan string
	ctx, cancel := context.WithCancel(context.Background())

	gptHotkeys := getGPTHotkeys()
	lastHit := time.Now()
	fmt.Printf("--- Please press %s to auto generate text --- \n", gptHotkeys)
	userCore.initUserCore()

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

			userCore.reloadMask()

			fmt.Println("### prompt:", userCore.mask)
			fmt.Println("### model:", userCore.model)
			fmt.Println("### user:")
			fmt.Println(clipboardContent)
			prompts, new := userCore.GeneratePromptMessages(clipboardContent)

			if len(prompts) == 0 {
				return
			}

			txtChan = make(chan string, 1024)
			cancel()
			ctx, cancel = context.WithCancel(context.Background())
			fmt.Println("Generating...")
			workDone := make(chan struct{}, 2)
			go st.ShowRunningIcon(ctx, workDone)
			go userCore.QueryGPT(ctx, txtChan, prompts)

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
						userCore.AddNewMessages(new)
						workDone <- struct{}{}
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
		if time.Now().Sub(lastEscHit).Milliseconds() < 500 {
			escCnt++
			fmt.Println("increase escCnt to", escCnt)
			if escCnt == 2 { //triple 'esc' click for quick clean context
				userCore.ClearContext()
				escCnt = 0
			}
		} else {
			escCnt = 0
		}
		lastEscHit = time.Now()
		cancel()
	})

	s := hook.Start()
	<-hook.Process(s)
}
