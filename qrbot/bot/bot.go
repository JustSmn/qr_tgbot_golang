package bot

//Основная логика бота. Обрабатывает обновления, команды и отправляет сообщения.

import (
	"log"
	"main.go/logger"
	"main.go/qrcode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	msgHelp  = "Привет!!! Отправь мне свою ссылку и я моментально пришлю тебе ее QR код😵‍💫🤤😋😍🤪"
	msgInfo  = "Это мой петпроектик, созданный для личных нужд (частенько нужно переделать ссылки в qr коды, поэтому написал данного простенького бота) v.1.2 \n Создался этот проект на языке программирования Golang. Надеюсь, я буду поддерживать и дальше этот проектик (в идеале мне бы все переделать и добавить кучу крутых штук, нооо, не знаю, справлюсь или нет)👩‍💻🧑‍💻👨‍💻🫦🫂🙀😼😼😼😋😮‍💨🥵💋😍 \n JustSMN 2025 Все права защищены (неправда)"
	msgPrank = "На вас была наведена порча. Вы в опасности🙀💀👻😳!!! Чтобы снять порчу переведите мне немнога денег. 89833648367 СБП Сбер, Тинькофф, Альфа, можно наличкой в 3.15 ночи с заженными свечами протянуть в зеркало.\n Только тогда будет снята порча😈😈😈 \n Да-да-да, всего хорошего вам😈😈😈"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	api.Debug = true
	log.Printf("Авторизован как %s", api.Self.UserName)

	return &Bot{api: api}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Логируем сообщение пользователя
		logger.LogMessage(update.Message.From.UserName, update.Message.Text, update.Message.Time())

		// Обрабатываем команды
		switch update.Message.Command() {
		case "help", "start":
			b.sendMessage(update.Message.Chat.ID, msgHelp)
			continue
		case "info":
			b.sendMessage(update.Message.Chat.ID, msgInfo)
			continue
		case "prank":
			b.sendMessage(update.Message.Chat.ID, msgPrank)
			continue
		}

		// Генерируем и отправляем QR-код
		filePath, err := qrcode.GenerateQRCode(update.Message.Text)
		if err != nil {
			log.Println("Ошибка при генерации QR-кода:", err)
			continue
		}

		b.sendPhoto(update.Message.Chat.ID, filePath)
	}
}

func (b *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.api.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}

func (b *Bot) sendPhoto(chatID int64, filePath string) {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(filePath))
	photo.Caption = "Вот твой QR-код!"
	_, err := b.api.Send(photo)
	if err != nil {
		log.Println("Ошибка при отправке фото:", err)
	}
}
