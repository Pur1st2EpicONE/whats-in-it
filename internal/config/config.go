package config

import (
	"github.com/spf13/viper"
)

func InitConfig() error {

	viper.AddConfigPath("/etc/whats-in-it/")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
