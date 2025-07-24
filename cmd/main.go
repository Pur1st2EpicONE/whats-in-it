package main

import (
	"fmt"
	"log/slog"
	"os"

	client "github.com/Pur1st2EpicONE/whats-in-it/internal/client"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/config"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/logger"
)

func main() {

	slog.SetDefault(logger.InitLogger())

	if err := config.InitConfig(); err != nil {
		logger.LogFatal("config initialization failed: ", err)
	}

	file, err := checkFile()
	if err != nil {
		os.Exit(1)
	}

	chatClient := client.NewChatClient()

	token, err := chatClient.GetToken()
	if err != nil {
		logger.LogFatal("failed to get token: ", err)
	}

	apiResponse, err := chatClient.AskWhatsInIt(file, token)
	if err != nil {
		logger.LogFatal("failed to ask gigaChat: ", err)
	}

	chatAnswer, err := chatClient.InterpretAnswer(apiResponse)
	if err != nil {
		logger.LogFatal("failed to get answer from gigaChat: ", err)
	}

	fmt.Println(chatAnswer.GetResponse())
}
