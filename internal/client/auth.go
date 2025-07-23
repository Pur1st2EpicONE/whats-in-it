package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GetToken(client *http.Client) (models.Token, error) {
	tokenResponse, err := client.Do(tokenRequest(viper.GetString("auth_key")))
	if err != nil {
		return models.Token{}, err
	}

	jsonToken, err := io.ReadAll(tokenResponse.Body)
	tokenResponse.Body.Close()
	if err != nil {
		return models.Token{}, err
	}

	var token models.Token
	if err := json.Unmarshal(jsonToken, &token); err != nil {
		return models.Token{}, err
	}
	return token, nil
}

func tokenRequest(authKey string) *http.Request {
	data := url.Values{}
	data.Set("scope", viper.GetString("scope"))

	request, _ := http.NewRequest(
		"POST",
		viper.GetString("token_endpoint"),
		bytes.NewBufferString(data.Encode()),
	)

	request.Header.Set("Authorization", "Bearer "+authKey)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("RqUID", uuid.New().String())
	return request
}
