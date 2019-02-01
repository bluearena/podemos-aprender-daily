package main

import (
	"bytes"
	"log"
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
			log.Printf("[%s] %s", update.Message.From.UserName, msgText)

			response := getResponse(msgText)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

	}
}

func getResponse(msgText string) string {

	// get project and hours from message text
	msgContent := strings.Split(msgText, ",")

	formatErr := "Para que el registro de horas pueda realizarse correctamente debe ingresarlo de la siguiente forma, ej: @time, banco de tiempo, 15"

	if len(msgContent) != 3 {
		return formatErr
	}

	// extract project name form array
	project := msgContent[1]

	// remove spaces and convert to integer
	hours, err := strconv.Atoi(removeSpaces(msgContent[2]))
	if err != nil {
		return formatErr
	}

	addHours(project, hours)
	hoursInvested := getTimeInvested(project)

	var response bytes.Buffer // efficient way to concatenate strings
	response.WriteString("Actualmente se llevan ")
	response.WriteString(strconv.Itoa(hoursInvested))
	response.WriteString(" invertidas en el proyecto ")
	response.WriteString(project)
	return response.String()
}

func removeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

/* it will be better split logic and access data,
but for the moment we will maintain all together in the same layer */
func addHours(project string, hours int) {

}

func getTimeInvested(project string) int {
	return 0
}
