package yandex

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
	"github.com/spf13/viper"
)

func (y *YandexGPT) GetToken() (models.Token, error) {
	tokenResponse, err := y.httpClient.Do(tokenRequest(viper.GetString("yandex_verification_code")))
	if err != nil {
		return nil, err
	}

	jsonToken, err := io.ReadAll(tokenResponse.Body)
	tokenResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	var token models.YandexToken
	if err := json.Unmarshal(jsonToken, &token); err != nil {
		return nil, err
	}
	return token, nil
}

func tokenRequest(verificationCode string) *http.Request {
	jsonRequest, err := json.Marshal(newTokenRequest(verificationCode))
	if err != nil {
		return &http.Request{}
	}
	request, _ := http.NewRequest(
		"POST",
		viper.GetString("yandex_token_endpoint"),
		bytes.NewBuffer(jsonRequest))
	return request

}

func newTokenRequest(verificationCode string) *models.YandexTokenRequest {
	return &models.YandexTokenRequest{OauthToken: verificationCode}
}
