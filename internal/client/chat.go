package client

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

func AskGigaChat(file string, client *http.Client, token models.Token) (*http.Response, error) {
	jsonRequest, err := json.Marshal(newChatRequest(file, viper.GetString("language")))
	if err != nil {
		return &http.Response{}, err
	}

	apiResponse, err := client.Do(apiRequest(jsonRequest, token))
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, nil
}

func newChatRequest(file string, language string) *models.ChatRequest {
	return &models.ChatRequest{
		Model:          viper.GetString("model"),
		Stream:         false,
		UpdateInterval: 0,
		Messages: []models.Message{
			{Role: "system", Content: fmt.Sprintf("Ты - утилита для анализа текстовых файлов. Проанализируй содержимое переданного файла и ответь на вопрос «что в нём?». Ответь одним коротким, но ясным предложением. Язык ответа - %s. Если в файле просто какой-то бессмысленный набор символов, так и скажи.\n\n Файл для анализа: %s", language, file)},
		},
	}
}

func GetAnswer(apiResponse *http.Response) (models.ChatResponse, error) {
	jsonChatResponse, err := io.ReadAll(apiResponse.Body)
	apiResponse.Body.Close()
	if err != nil {
		logger.LogFatal("", err)
	}

	var chatResponse models.ChatResponse
	if err := json.Unmarshal(jsonChatResponse, &chatResponse); err != nil {
		return chatResponse, err
	}
	return chatResponse, nil
}

func apiRequest(jsonRequest []byte, token models.Token) *http.Request {
	request, err := http.NewRequest("POST", viper.GetString("chat_endpoint"), bytes.NewBuffer(jsonRequest))
	if err != nil {
		logger.LogFatal("", err)
	}

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	request.Header.Set("Content-Type", "application/json")
	return request
}
