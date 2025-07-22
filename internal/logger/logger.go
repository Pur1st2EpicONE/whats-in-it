package logger

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func LogFatal(msg string, err error) {
	slog.Error(msg, slog.String("err", err.Error()))
	os.Exit(1)
}
