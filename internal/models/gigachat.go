package models

import (
	"fmt"
	"log"
)

type GigaChatToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

func (g GigaChatToken) GetToken() string {
	return g.AccessToken
}

type GigaChatRequest struct {
	Model          string            `json:"model"`
	Stream         bool              `json:"stream"`
	UpdateInterval int               `json:"update_interval"`
	Messages       []GigaChatMessage `json:"messages"`
}

type GigaChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GigaChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (g GigaChatResponse) GetResponse() (string, error) {
	if len(g.Choices) == 0 {
		log.Printf("GigaChat response has no choices")
		return "", fmt.Errorf("no choices in response")
	}

	content := g.Choices[0].Message.Content
	if content == "" {
		log.Printf("GigaChat is silent")
		return "", fmt.Errorf("empty response")
	}

	return content, nil
}
