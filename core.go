package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hanyuancheung/gpt-go"
	"golang.design/x/clipboard"
)

type Core struct {
	u           *UserCore
	st          *SysTray
	txtChan     chan string
	ctx         context.Context
	cancel      context.CancelFunc
	lastHit     time.Time
	lastEscHit  time.Time
	escCnt      int
	queryString string
}

func NewCore(u *UserCore, st *SysTray) *Core {
	c := Core{u: u, st: st}
	c.init()
	return &c
}

func (c *Core) init() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.lastEscHit = time.Now()
	c.escCnt = 0
}

func (c *Core) queryHit() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	c.queryString = string(clipboard.Read(clipboard.FmtText))
	fmt.Println(c.queryString)

	go func() {
		if time.Since(c.lastHit).Milliseconds() > 1000 {
			c.lastHit = time.Now()
		} else {
			return
		}

		if err != nil {
			fmt.Println("Failed to read clipboard content:", err)
			return
		}

		if len(c.queryString) < 1 {
			fmt.Println("Empty question")
			return
		}

		c.u.reloadMask()

		fmt.Println("### prompt:", c.u.mask)
		fmt.Println("### model:", c.u.model)
		fmt.Println("### user:")
		fmt.Println(c.queryString)
		prompts, new := c.u.GeneratePromptMessages(c.queryString)

		if len(prompts) == 0 {
			return
		}

		c.txtChan = make(chan string, 1024)
		fmt.Println("Generating...")
		c.cancel()
		c.ctx, c.cancel = context.WithCancel(context.Background())

		TypeStr(UText("Working..."))
		workDone := make(chan struct{}, 2)
		go c.st.ShowRunningIcon(c.ctx, workDone)
		go c.u.QueryGPT(c.ctx, c.txtChan, prompts)

		assistantAns := ""
		fmt.Print("### Assistant:\n")

		defer func() {
			//recover clipboard
			clipboard.Write(clipboard.FmtText, []byte(c.queryString))
		}()

		for {
			select {
			case txt, ok := <-c.txtChan:
				if !ok {
					// txtChan is closed, exit the loop
					fmt.Print("\n")
					new = append(new, gpt.ChatCompletionRequestMessage{
						Role:    "assistant",
						Content: assistantAns,
					})
					c.u.AddNewMessages(new)
					workDone <- struct{}{}
					return
				}

				if assistantAns == "" {
					for i := 0; i < len([]rune(UText("Working..."))); i++ {
						TypeBackspace()
					}
				}

				fmt.Print(txt)

				txt = strings.ReplaceAll(txt, "\r\n", "\n")
				for i, t := range strings.Split(txt, "\n") {
					if i > 0 {
						TypeEnter()
					}
					if len(t) > 0 {
						TypeStr(t)
					}
				}
				assistantAns += txt
			case <-c.ctx.Done():
				// ctx is done, exit the loop
				return
			}
		}
	}()
}

func (c *Core) escapeHit() {
	if time.Since(c.lastEscHit).Milliseconds() < 500 {
		c.escCnt++
		fmt.Println("increase escCnt to", c.escCnt)
		if c.escCnt == 2 { //triple 'esc' click for quick clean context
			c.u.ClearContext()
			c.escCnt = 0
		}
	} else {
		c.escCnt = 0
	}
	c.lastEscHit = time.Now()
	c.cancel()
}
