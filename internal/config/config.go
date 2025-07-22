package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
