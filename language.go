package main

import (
	"fmt"

	"github.com/jeandeaual/go-locale"
)

var languages = map[string]map[string]string{
	"About": {
		"en": "About",
		"zh": "关于",
	},
	"InputGPT": {
		"en": "InputGPT",
		"zh": "InputGPT",
	},
	"InputGPT a Helpful input Assistant": {
		"en": "InputGPT",
		"zh": "InputGPT",
	},
	"Exit": {
		"en": "Exit",
		"zh": "退出",
	},
	"About the App": {

		"en": "About the App",
		"zh": "退出APP",
	},
	"Quit the whole app": {
		"zh": "退出APP",
	},
	"Click to active GPT": {
		"en": "Click to active GPT",
		"zh": "点击激活 GPT",
	},
	"Clear Context": {
		"zh": "清除所有上下文",
	},
	"Default": {
		"zh": "默认",
	},
	"Copy the question then click \"%s\" to query GPT": {
		"zh": "复制提问的文本，再按 \"%s\" 发送给GPT",
	},
	"Clear Context %d/%d": {
		"zh": "清除所有上下文 %d/%d",
	},
}

var LANG = "en"

func setLang() {
	userLanguage, err := locale.GetLanguage()
	if err == nil {
		fmt.Println("Language:", userLanguage)
		LANG = userLanguage
	}
}

// give the right lanuage with query text when no match return the langText itself
func UText(langText string) string {
	if lang, ok := languages[langText]; ok {
		if text, ok := lang[LANG]; ok {
			return text
		}
	}
	if LANG != "en" {
		fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, LANG)
	}

	return langText
}
