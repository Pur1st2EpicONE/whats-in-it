package yandex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Pur1st2EpicONE/whats-in-it/internal/logger"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/models"
	"github.com/spf13/viper"
)

func (y *YandexGPT) AskWhatsInIt(file string, token models.Token) (*http.Response, error) {
	jsonRequest, err := json.Marshal(newChatRequest(file, viper.GetString("language")))
	if err != nil {
		return &http.Response{}, err
	}

	apiResponse, err := y.httpClient.Do(apiRequest(jsonRequest, token))
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, nil
}

func newChatRequest(file string, language string) *models.YandexRequest {
	return &models.YandexRequest{
		ModelURI: viper.GetString("yandex_gpt.model_uri"),
		CompletionOptions: models.CompletionOptions{
			Stream:      false,
			Temperature: 0.3,
			MaxTokens:   "500",
			ReasoningOptions: models.ReasoningOptions{
				Mode: "DISABLED",
			},
		},
		Messages: []models.Message{
			{
				Role: "user",
				Text: fmt.Sprintf("Ты - утилита для анализа текстовых файлов. Проанализируй содержимое переданного файла и ответь на вопрос «что в нём?». Ответь одним коротким, но ясным предложением. Язык ответа - %s. Если в файле просто какой-то бессмысленный набор символов, так и скажи.\n\n Файл для анализа: %s", language, file),
			},
		},
	}
}

func (y *YandexGPT) InterpretAnswer(apiResponse *http.Response) (models.Response, error) {
	jsonChatResponse, err := io.ReadAll(apiResponse.Body)
	apiResponse.Body.Close()
	if err != nil {
		logger.LogFatal("", err)
	}

	var chatResponse models.YandexResponse
	if err := json.Unmarshal(jsonChatResponse, &chatResponse); err != nil {
		return chatResponse, err
	}
	return chatResponse, nil
}

func apiRequest(jsonRequest []byte, token models.Token) *http.Request {
	request, err := http.NewRequest("POST", viper.GetString("yandex_gpt.chat_endpoint"), bytes.NewBuffer(jsonRequest))
	if err != nil {
		logger.LogFatal("", err)
	}

	request.Header.Set("Authorization", "Bearer "+token.GetToken())
	request.Header.Set("Content-Type", "application/json")
	return request
}
