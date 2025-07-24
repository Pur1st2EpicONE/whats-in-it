package client

import (
	"net/http"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/client/gigachat"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
)

type GPT interface {
	GetToken() (models.Token, error)
	AskWhatsInIt(string, models.Token) (*http.Response, error)
	InterpretAnswer(*http.Response) (models.Response, error)
}

type ChatClient struct {
	GPT
}

func NewChatClient() *ChatClient {
	return &ChatClient{GPT: gigachat.NewGigaChat()}
}
