package models

import (
	"fmt"
	"log"
)

type YandexToken struct {
	AccessToken string `json:"iamToken"`
	ExpiresAt   string `json:"expiresAt"`
}

type YandexTokenRequest struct {
	OauthToken string `json:"yandexPassportOauthToken"`
}

func (i YandexToken) GetToken() string {
	return i.AccessToken
}

type YandexRequest struct {
	ModelURI          string            `json:"modelUri"`
	CompletionOptions CompletionOptions `json:"completionOptions"`
	Messages          []Message         `json:"messages"`
}

type CompletionOptions struct {
	Stream           bool             `json:"stream"`
	Temperature      float64          `json:"temperature"`
	MaxTokens        string           `json:"maxTokens"`
	ReasoningOptions ReasoningOptions `json:"reasoningOptions"`
}

type ReasoningOptions struct {
	Mode string `json:"mode"`
}

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type YandexResponse struct {
	Result struct {
		Alternatives []struct {
			Message YandexMessage `json:"message"`
		} `json:"alternatives"`
	} `json:"result"`
}

type YandexMessage struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

func (y YandexResponse) GiveAnswer() (string, error) {
	if len(y.Result.Alternatives) == 0 {
		log.Printf("YandexGPT response has no choices")
		return "", fmt.Errorf("no choices in response")
	}

	text := y.Result.Alternatives[0].Message.Text
	if text == "" {
		log.Printf("YandexGPT is silent")
		return "", fmt.Errorf("empty response")
	}

	return text, nil
}
