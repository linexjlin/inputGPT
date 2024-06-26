package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func init() {
	godotenv.Load("env.txt")
}

var UText func(string) string
var UMenuText func(string) string
var gCore *Core

func initUText(l *Language) {
	UText = l.UText
}

func initUMenuText(l *Language) {
	UMenuText = func(s string) string {
		return l.UTextWithLangCode(s, "emoji") + l.UText(s)
	}
}

func queryHotkey() {
	gptHK := getGPTHotkeys()
	modKey, key := keyNameToModifier(gptHK[0]), keyNamesToKey(gptHK[1])
	hk := hotkey.New([]hotkey.Modifier{modKey}, key)
	err := hk.Register()
	if err != nil {
		fmt.Printf("hotkey: failed to register hotkey: %v", err)
		return
	}

	for {
		<-hk.Keyup()
		fmt.Printf("hotkey: %v is up\n", hk)
		gCore.queryHit()
	}
}

func escapeHotkey() {
	hk := hotkey.New([]hotkey.Modifier{}, hotkey.KeyEscape)
	err := hk.Register()
	if err != nil {
		fmt.Printf("hotkey: failed to register hotkey: %v\n", err)
		return
	}

	<-hk.Keyup()
	hk.Unregister()
	escKey()
	gCore.escapeHit()
	mainthread.Init(escapeHotkey)
}

func main() {
	var uc UserCore
	var l *Language
	l = NewLanguage()
	OSDepCheck()
	initUText(l)
	initUMenuText(l)
	uc.initUserCore()
	st := SysTray{userCore: &uc}

	gCore = NewCore(&uc, &st)
	go mainthread.Init(queryHotkey)
	go mainthread.Init(escapeHotkey)
	st.Run()
}
