package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hanyuancheung/gpt-go"
)

func queryGPT(ctx context.Context, txtChan chan string, messages []gpt.ChatCompletionRequestMessage) {
	fmt.Println("Query messages:")
	showAsJson(messages)
	client := gpt.NewClient(
		getOpenAIkey(),
		gpt.WithBaseURL(getOpenAIBaseUrl()),
		gpt.WithTimeout(600*time.Second),
	)

	err := client.ChatCompletionStream(ctx, &gpt.ChatCompletionRequest{
		Model:    g_userCore.model,
		Messages: messages,
	}, func(response *gpt.ChatCompletionStreamResponse) {
		//fmt.Println(response.Choices)
		//fmt.Printf("%+v\n", response)
		if len(response.Choices) > 0 {
			if response.Choices[0].Delta.Content != "" {
				txtChan <- response.Choices[0].Delta.Content
			}
			if response.Choices[0].FinishReason == "stop" {
				close(txtChan)
			}
		}
	})
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		txtChan <- fmt.Sprintf("ChatCompletionStream error: %v\n", err)
		close(txtChan)
		return
	}
}
