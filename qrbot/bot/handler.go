package bot

//Обработчики команд. Если команд станет больше, их можно будет легко добавить.

import (
	"log"
	"main.go/qrcode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "help", "start":
		b.sendMessage(update.Message.Chat.ID, msgHelp)
	case "info":
		b.sendMessage(update.Message.Chat.ID, msgInfo)
	case "prank":
		b.sendMessage(update.Message.Chat.ID, msgPrank)
	default:
		filePath, err := qrcode.GenerateQRCode(update.Message.Text)
		if err != nil {
			log.Println("Ошибка при генерации QR-кода:", err)
			return
		}
		b.sendPhoto(update.Message.Chat.ID, filePath)
	}
}
