package client

import (
	"log"
	"net/http"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/client/gigachat"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/client/yandex"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
	"github.com/spf13/viper"
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
	currentModel := viper.GetString("current_model")
	switch currentModel {
	case "giga_chat":
		return &ChatClient{GPT: gigachat.NewGigaChat()}
	case "yandex_gpt":
		return &ChatClient{GPT: yandex.NewYandexGPT()}
	default:
		log.Fatalf("unknown GPT model: %s", currentModel)
		return nil
	}
}
