package main

import (
	"context"
	"fmt"
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

	if time.Since(c.lastHit).Milliseconds() > 1000 {
		c.lastHit = time.Now()
	} else {
		return
	}

	if len(c.queryString) < 1 {
		fmt.Println("Empty question")
		return
	}

	fmt.Println("### prompt:", c.u.mask)
	fmt.Println("### user:")
	fmt.Println(c.queryString)

	go func() {
		fmt.Println("models:", c.u.models)
		if c.u.maskModel != "" {
			fmt.Println("using maskModel", c.u.maskModel)
		} else if len(c.u.models) > 1 {
			fmt.Println("using muti models", c.u.models)
			for i, model := range c.u.models {
				//check c.ctx status before loop
				fmt.Println("calling", model)
				TypeStr(fmt.Sprintf("<%s>\n", model))
				c.queryWithMode(model)
				if i == len(c.u.models) {
					TypeStr(fmt.Sprintf("\n</%s>", model))
				} else {
					TypeStr(fmt.Sprintf("\n</%s>\n\n", model))
				}

				select {
				case <-c.ctx.Done():
					fmt.Println("Context canceled or deadline exceeded during iteration")
					return
				default:
					// Continue with the current iteration
				}
			}
		} else if len(c.u.models) > 0 {
			fmt.Println("using one models", c.u.models)
			c.queryWithMode(c.u.models[0])
		} else {
			fmt.Println("using defaultMode:", c.u.defaultModel)
			c.queryWithMode(c.u.defaultModel)
		}
		clipboard.Write(clipboard.FmtText, []byte(c.queryString))
	}()
}

const (
	ThinkingStr = "⏳"
)

func (c *Core) queryWithMode(model string) {
	fmt.Println("### model:", model)
	prompts, new := c.u.GeneratePromptMessages(c.queryString)

	if len(prompts) == 0 {
		return
	}

	c.txtChan = make(chan string, 1024)
	fmt.Println("Generating...")
	c.cancel()
	c.ctx, c.cancel = context.WithCancel(context.Background())

	TypeStr(ThinkingStr)
	workDone := make(chan struct{}, 2)
	go c.st.ShowRunningIcon(c.ctx, workDone)
	go c.u.QueryGPT(c.ctx, model, c.txtChan, prompts)

	assistantAns := ""
	fmt.Print("### Assistant:\n")

	defer func() {
		//recover clipboard
		//clipboard.Write(clipboard.FmtText, []byte(c.queryString))
	}()

	tmpText := ""
	nextType := time.Now()
	stop := false
	for {
		if t, ok := <-c.txtChan; ok {
			fmt.Print(t)
			tmpText = tmpText + t
		} else {
			time.Sleep(time.Until(nextType))
			stop = true
		}

		if time.Now().After(nextType) {
			if assistantAns == "" {
				for i := 0; i < len([]rune(ThinkingStr)); i++ {
					TypeBackspace()
					time.Sleep(time.Millisecond * 10)
				}
			}
			assistantAns = assistantAns + tmpText
			if tmpText != "" {
				TypeStr(tmpText)
				tmpText = ""
				nextType = time.Now().Add(time.Millisecond * 200) //write interavl 100 milliseconds
			}

			if stop {
				if time.Since(nextType).Microseconds() < 0 {
					time.Sleep(time.Since(nextType) * -1)
				}

				fmt.Print("\n")
				new = append(new, gpt.ChatCompletionRequestMessage{
					Role:    "assistant",
					Content: assistantAns,
				})
				c.u.AddNewMessages(new)
				workDone <- struct{}{}
				break
			}
		}
	}
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
