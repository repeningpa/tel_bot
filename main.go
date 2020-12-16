package main

import (
	"bytes"
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	MakeRequest()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		reply := "Ты эт, не шуми тут..."
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "привет":
			reply = "Пароль не верный."
		case "hello":
			reply = "world"
		case "fendi":
			reply = "FENDI GUCHI FLIP-FLOP"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}

func MakeRequest() {
	message := map[string]interface{}{
		"chat_id": "496818745",
		"text": "Меня пересобрали..."
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api.telegram.org/bot1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM/sendMessage", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
}
