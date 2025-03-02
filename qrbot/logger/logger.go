package logger

//Логирование сообщений пользователей. Выводит в терминал ник, текст и время отправки.

import (
	"log"
	"time"
)

func LogMessage(username, text string, timestamp time.Time) {
	t := timestamp.Format("2006-01-02 15:04:05")
	log.Printf("[%s] %s: %s\n", t, username, text)
}
