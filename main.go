package main

import (
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("env.txt")
}

var UText func(string) string
var UMenuText func(string) string

func initUText(l *Language) {
	UText = l.UText
	return
}

func initUMenuText(l *Language) {
	UMenuText = func(s string) string {
		return l.UTextWithLangCode(s, "emoji") + l.UText(s)
	}
}

func main() {
	var uc UserCore
	var l *Language
	l = NewLanguage()
	OSDepCheck()
	initUText(l)
	initUMenuText(l)
	st := SysTray{userCore: &uc}
	go registerHotKeys(&uc, &st)
	st.Run()
}
