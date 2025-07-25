package gigachat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/logger"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func (g *GigaChat) GetToken() (models.Token, error) {
	tokenResponse, err := g.httpClient.Do(tokenRequest(viper.GetString("giga_chat.auth_key")))
	if err != nil {
		return nil, err
	}

	jsonToken, err := io.ReadAll(tokenResponse.Body)
	tokenResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	var token models.GigaChatToken
	if err := json.Unmarshal(jsonToken, &token); err != nil {
		return nil, err
	}
	return token, nil
}

func tokenRequest(authKey string) *http.Request {
	data := url.Values{}
	data.Set("scope", viper.GetString("giga_chat.scope"))

	request, err := http.NewRequest(
		"POST",
		viper.GetString("giga_chat.token_endpoint"),
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		logger.LogFatal("token request failed", err)
	}

	request.Header.Set("Authorization", "Bearer "+authKey)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("RqUID", uuid.New().String())
	return request
}
