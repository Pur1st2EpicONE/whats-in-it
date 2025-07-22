package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func checkFile() (string, error) {
	file, err := getFilePathFromArgs(os.Args)
	if err != nil {
		return "", err
	}
	data, err := readFile(file)
	if err != nil {
		return "", err
	}
	if err := checkIfEmpty(data); err != nil {
		return "", err
	}

	return data, nil
}

func getFilePathFromArgs(args []string) (string, error) {
	if len(args) != 2 {
		fmt.Println("Usage: wii [FILE]")
		return "", errors.New("incorrect number of arguments")
	}
	return args[1], nil
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("wii: %s: No such file or directory\n", file)
		return "", err
	}
	return string(data), nil
}

func checkIfEmpty(data string) error {
	if len(data) == 0 { // сами обрабатываем ситуацию с пустым файлом, чтобы попусту не тратить токены
		switch viper.GetString("language") {
		case "русский":
			fmt.Println("Файл пуст.")
		case "english":
			fmt.Println("File is empty.")
		case "deutsch":
			fmt.Println("Datei ist leer.")
		case "汉语":
			fmt.Println("文件为空.")
		default:
			fmt.Println("📁🚫📃")
		}
		return errors.New("empty file")
	}
	return nil
}
