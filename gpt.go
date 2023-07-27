package main

import (
	"context"
	"fmt"

	"github.com/hanyuancheung/gpt-go"
)

func queryGTP(txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	var opt = gpt.WithBaseURL(getOpenAIBaseUrl())
	client := gpt.NewClient(getOpenAIkey(), opt)

	err := client.ChatCompletionStream(context.Background(), &gpt.ChatCompletionRequest{
		Model:    g_userSetting.model,
		Messages: messages,
	}, func(response *gpt.ChatCompletionStreamResponse) {
		if response.Choices[0].Delta.Content != "" {
			txtChan <- response.Choices[0].Delta.Content
		}
		if response.Choices[0].FinishReason == "stop" {
			close(txtChan)
		}
	})
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		txtChan <- fmt.Sprintf("ChatCompletionStream error: %v\n", err)
		close(txtChan)
		return
	}
}
