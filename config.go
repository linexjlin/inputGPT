package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/go-vgo/robotgo"
	_ "github.com/joho/godotenv/autoload"
)

func getGPTHotkeys() []string {
	hotkeys := os.Getenv("GPT_HOTKEYS")
	if hotkeys == "" {
		if runtime.GOOS == "windows" {
			hotkeys = "q+ctrl+shift"
		} else if runtime.GOOS == "darwin" {
			hotkeys = "z+shift+command"
		} else {
			hotkeys = "z+ctrl+alt"
		}
	}
	return strings.Split(hotkeys, "+")
}

func pressPaste() {
	if runtime.GOOS == "windows" {
		robotgo.KeyTap("v", "control")
	} else if runtime.GOOS == "darwin" {
		robotgo.KeyTap("v", "command")
	} else {
		robotgo.KeyTap("v", "control")
	}
}

func getOpenAIkey() string {
	return os.Getenv("OPENAI_API_KEY")
}
