package main

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := "Ты эт, не шуми тут..."
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "привет":
			reply = "Пароль не верный."
		case "hello":
			reply = "world"
		case "fendi":
			reply = "FENDI GUCHI FLIP-FLOP"
		}

		// создаем ответное сообщение

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
