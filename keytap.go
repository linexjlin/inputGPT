package main

import (
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
)

func paste() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_V)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)

}

func backspace() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_BACKSPACE)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func enter() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_ENTER)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func TypeStr(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))
	paste()
}

func TypeEnter() {
	enter()
}

func TypeBackspace() {
	backspace()
}