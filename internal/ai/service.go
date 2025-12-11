package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client *openai.Client
}

func NewService(apiKey string) *Service {
	// 1. Configure for OpenRouter
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1" // Point to OpenRouter
	
	client := openai.NewClientWithConfig(config)
	
	return &Service{
		client: client,
	}
}

func (s *Service) ProcessUserMessage(ctx context.Context, userMsg string) (string, error) {
	
	// 2. Create Request
	req := openai.ChatCompletionRequest{
		// Free Model on OpenRouter
		Model: "z-ai/glm-4.5-air:free", 
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a Vedic Astrologer. If you receive birth details, call the get_kundali_details tool.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userMsg,
			},
		},
		Tools: GetTools(),
	}

	// 3. Send Request
	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	msg := resp.Choices[0].Message

	// 4. Check if AI called the Tool
	if len(msg.ToolCalls) > 0 {
		call := msg.ToolCalls[0]
		// Return the function name and the JSON arguments the AI generated
		return fmt.Sprintf(">>> ACTION: %s | ARGS: %s", 
			call.Function.Name, 
			call.Function.Arguments), nil
	}

	// 5. Otherwise return normal text
	return msg.Content, nil
}