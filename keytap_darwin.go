package main

import (
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
)

func escKey() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_ESC)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 50)
}

func paste() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.HasSuper(true)
	kb.SetKeys(keybd_event.VK_V)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	//time.Sleep(time.Millisecond * 10)
}

func backspace() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_DELETE)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	//time.Sleep(time.Millisecond * 10)
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
