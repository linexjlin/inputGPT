package main

import (
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
)

var g_userCore UserCore
var g_languages *Language

func main() {
	godotenv.Load("env.txt")
	g_languages = NewLanguage()
	OSDepCheck()
	go registerHotKeys(&g_userCore)
	systray.Run(onReady, onExit)
}
