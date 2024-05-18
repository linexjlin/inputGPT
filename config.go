package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

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
}

func getModeList() []string {
	modesEnv := os.Getenv("MODES")
	if modesEnv == "" {
		modesEnv = "gpt-3.5-turbo"
	}

	modes := strings.Split(modesEnv, ",")
	if len(modes) == 0 {
		modes = []string{"gpt-3.5-turbo"}
	}
	fmt.Println(len(modes))
	return modes
}

type ModePrompt struct {
	Name         string                             `json:"name"`
	Model        string                             `json:"model"`
	HeadMessages []gpt.ChatCompletionRequestMessage `json:"headMessages"`
	MaxContext   int                                `json:"maxContext"`
}

func loadModePrompt(filepath string) (ModePrompt, error) {
	var prompt ModePrompt

	// Read the file contents
	data, err := os.ReadFile(filepath)
	if err != nil {
		return prompt, err
	}

	// Unmarshal the JSON data into the ModePrompt struct
	err = json.Unmarshal(data, &prompt)
	if err != nil {
		return prompt, err
	}

	return prompt, nil
}

func parseModePrompt(content string) (ModePrompt, error) {
	var prompt ModePrompt

	// Unmarshal the JSON data into the ModePrompt struct
	err := json.Unmarshal([]byte(content), &prompt)
	if err != nil {
		return prompt, err
	}

	return prompt, nil
}

func saveModePrompt(prompt ModePrompt, filepath string) error {
	// Marshal the ModePrompt struct into JSON data
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
