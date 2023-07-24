package main

import (
	"context"
	"fmt"

	"github.com/hanyuancheung/gpt-go"
)

func queryGTP(txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	client := gpt.NewClient(getOpenAIkey())
	err := client.ChatCompletionStream(context.Background(), &gpt.ChatCompletionRequest{
		Model:    "gpt-3.5-turbo-0613",
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
