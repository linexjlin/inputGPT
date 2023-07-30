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
	"Import": {
		"zh": "导入",
	},
	"Manage Prompts": {
		"zh": "管理引导词",
	},
	"Manager the prompts": {
		"zh": "打开文件夹，管理引导词",
	},
	"Select this prompt": {
		"zh": "选择这个Prompt",
	},
	"Open the project page": {
		"zh": "打开项目页面",
	},
	"Set API KEY": {
		"zh": "设置 API KEY",
	},
	"Modify, Delete prompts": {
		"zh": "更新、删除引导词",
	},
	"Import a prompt from clipboard": {
		"zh": "把剪贴板上的Prompt导入",
	},
	"Set the OpenAI KEY, baseurl etc..": {
		"zh": "请设置OpenAI KEY，baseurl等参数。",
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
