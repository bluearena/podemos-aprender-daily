package main

import (
	"log"
	"podemos-aprender-daily/businessLogicLayer"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// t.me/PodemosAprenderDailyBot
	bot, err := tgbotapi.NewBotAPI("725282157:AAEswHmZOK8xJ7_vNQhrQLHu1OqBqU8w0-M")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		// ignore any non-message updates
		if update.Message == nil {
			continue
		}
		msgText := update.Message.Text

		if strings.HasPrefix(msgText, "@time") {

			response := businessLogicLayer.ProcessMsg(msgText)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

	}
}
