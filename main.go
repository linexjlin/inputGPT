package main

import (
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
)

var g_userCore UserCore

func main() {
	godotenv.Load("env.txt")
	setLang()
	g_languages.Load()
	OSDepCheck()
	go registerHotKeys(&g_userCore)
	systray.Run(onReady, onExit)
}
