package main

import (
	"fmt"
	"log/slog"
	"os"

	client "github.com/Pur1st2EpicONE/whats-in-it/internal/client"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/config"
	"github.com/Pur1st2EpicONE/whats-in-it/internal/logger"
	"github.com/spf13/viper"
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

	currentModel := viper.GetString("current_model")
	chatClient := client.NewChatClient(currentModel)

	token, err := chatClient.GetToken()
	if err != nil {
		logger.LogFatal("failed to get token: ", err)
	}

	apiResponse, err := chatClient.AskWhatsInIt(file, token)
	if err != nil {
		logger.LogFatal(fmt.Sprintf("failed to ask %s: ", currentModel), err)
	}

	chatAnswer, err := chatClient.InterpretAnswer(apiResponse)
	if err != nil {
		logger.LogFatal(fmt.Sprintf("failed to get answer from %s: ", currentModel), err)
	}

	response, err := chatAnswer.GetResponse()
	if err == nil {
		fmt.Println(response)
	}
}
