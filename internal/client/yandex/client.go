package yandex

import (
	"net/http"
)

type YandexGPT struct {
	httpClient *http.Client
}

func NewYandexGPT() *YandexGPT {
	return &YandexGPT{httpClient: &http.Client{}}
}
