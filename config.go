package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/hanyuancheung/gpt-go"
)

func isDebug() bool {
	return os.Getenv("DEBUG") != ""
}

func getGPTHotkeys() []string {
	hotkeys := os.Getenv("GPT_HOTKEYS")
	if hotkeys == "" {
		if runtime.GOOS == "windows" {
			hotkeys = "space+shift"
		} else if runtime.GOOS == "darwin" {
			hotkeys = "space+shift"
		} else {
			hotkeys = "space+shift"
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
		return 0
	} else {
		return maxContextInt
	}
	return 0
}

type Prompt struct {
	Name         string                             `json:"name"`
	Model        string                             `json:"model"`
	HeadMessages []gpt.ChatCompletionRequestMessage `json:"headMessages"`
	MaxContext   int                                `json:"maxContext"`
}

func loadPrompt(filepath string) (Prompt, error) {
	var prompt Prompt

	// Read the file contents
	data, err := os.ReadFile(filepath)
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

func parsePrompt(content string) (Prompt, error) {
	var prompt Prompt

	// Unmarshal the JSON data into the Prompt struct
	err := json.Unmarshal([]byte(content), &prompt)
	if err != nil {
		return prompt, err
	}

	return prompt, nil
}

func savePrompt(prompt Prompt, filepath string) error {
	// Marshal the Prompt struct into JSON data
	data, err := json.Marshal(prompt)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
