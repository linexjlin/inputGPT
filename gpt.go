package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hanyuancheung/gpt-go"
)

func queryGTP(ctx context.Context, txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	client := gpt.NewClient(
		getOpenAIkey(),
		gpt.WithBaseURL(getOpenAIBaseUrl()),
		gpt.WithTimeout(600*time.Second),
	)

	err := client.ChatCompletionStream(ctx, &gpt.ChatCompletionRequest{
		Model:    g_userSetting.model,
		Messages: messages,
	}, func(response *gpt.ChatCompletionStreamResponse) {
		//fmt.Printf("%+v\n", response)
		if len(response.Choices) == 0 {
			txtChan <- fmt.Sprintf("%+v", response)
			close(txtChan)
		}
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
