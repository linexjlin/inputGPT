package main

import (
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
)

var g_userCore UserCore
var g_languages = Language{Data: make(map[string]map[string]string)}

func main() {
	godotenv.Load("env.txt")
	g_languages.SetLang()
	g_languages.Load()
	OSDepCheck()
	go registerHotKeys(&g_userCore)
	systray.Run(onReady, onExit)
}
