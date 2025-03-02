package main

//Точка входа в приложение. Создает бота и запускает его.

import (
	"log"
	"main.go/bot"
)

func main() {
	// Замени на свой токен
	botToken := "7819453922:AAHMYUiRh1M3RwXW_aeNJOZn2xgSdP0nlM4"

	// Создаем бота
	qrBot, err := bot.NewBot(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Запускаем бота
	qrBot.Start()
}
