package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"
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

func getOpenAIBaseUrl() string {
	if os.Getenv("OPENAI_API_BASE_URL") != "" {
		return os.Getenv("OPENAI_API_BASE_URL")
	}
	return "https://api.openai.com/v1"
}

func getMaxContext() int {
	maxContext := os.Getenv("MAX_CONTEXT")
	maxContextInt, err := strconv.Atoi(maxContext)
	if err != nil {
		fmt.Println("MAX_CONTEXT不是一个有效的数字")
		return 4
	} else {
		return maxContextInt
	}
	return 4
}

type Prompt struct {
	Model        string                             `json:"model"`
	HeadMessages []gpt.ChatCompletionRequestMessage `json:"headMessages"`
	MaxContext   int                                `json:"maxContext"`
}

func loadPrompt(filepath string) (Prompt, error) {
	var prompt Prompt

	// Read the file contents
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return prompt, err
	}

	// Unmarshal the JSON data into the Prompt struct
	err = json.Unmarshal(data, &prompt)
	if err != nil {
		return prompt, err
	}

	return prompt, nil
}
