package main

import (
	"time"

	"github.com/hanyuancheung/gpt-go"
	"github.com/jeandeaual/go-locale"
)

func renderMessages(messages []gpt.ChatCompletionRequestMessage, newMsg string) (retMsgs []gpt.ChatCompletionRequestMessage) {
	loclang, _ := locale.GetLocale()
	vars := map[string]string{
		"msg":    newMsg,
		"date":   time.Now().Format("2006-01-02"),
		"time":   time.Now().Format("15:04:05"),
		"mylang": loclang,
	}

	for _, message := range messages {
		m := message
		m.Content = applyStringTemplate(message.Content, vars)
		retMsgs = append(retMsgs, m)
	}

	return retMsgs
}
