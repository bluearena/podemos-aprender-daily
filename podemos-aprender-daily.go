package main

import (
	"bytes"
	"log"
	"podemos-aprender-daily/dataAccessLayer"
	"strconv"
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

			response := processMsg(msgText)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

	}
}

func processMsg(msgText string) string {

	// get project and hours from message text
	msgContent := strings.Split(msgText, ",")

	formatErr := "Para que el registro de horas pueda realizarse correctamente debe ingresarlo de la siguiente forma, ej: @time, banco de tiempo, 15"

	if len(msgContent) != 3 {
		return formatErr
	}

	// extract project name from array
	project := removeSpaces(msgContent[1])

	// remove spaces and convert to integer
	hours, err := strconv.Atoi(removeSpaces(msgContent[2]))
	if err != nil {
		return formatErr
	}

	dataAccessLayer.AddHours(project, hours)
	hoursInvested := dataAccessLayer.GetTimeInvested(project)

	var response bytes.Buffer // efficient way to concatenate strings
	response.WriteString("Actualmente se llevan ")
	response.WriteString(strconv.Itoa(hoursInvested))
	response.WriteString(" horas trabajadas en el proyecto ")
	response.WriteString(project)
	return response.String()
}

func removeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
