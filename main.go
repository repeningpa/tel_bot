package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

	////////////////////
	message := map[string]interface{}{
		"chat_id": "496818745",
		"text":    "Меня пересобрали...",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api.telegram.org/bot1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM/sendMessage", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		log.Fatalln(resp)
	}

	res, err := http.Get("http://img10.joyreactor.cc/pics/post/full/%D0%9A%D0%9F%D0%91%D0%A1%D0%90-Anime-%D1%84%D1%8D%D0%BD%D0%B4%D0%BE%D0%BC%D1%8B-%D0%B0%D0%B3%D0%B8%D1%82%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B5-%D0%BF%D0%BB%D0%B0%D0%BA%D0%B0%D1%82%D1%8B-5580394.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	fileByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	image := tgbotapi.FileBytes{Name: "build", Bytes: fileByte}
	_, err = bot.Send(tgbotapi.NewPhotoUpload(int64(496818745), image))
	if err != nil {
		log.Fatalln(err)
	}
	////////////////////

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := bot.GetUpdatesChan(u)

	var reply = ""
	for update := range updates {

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

// MakeRequest is ...
// func MakeRequest() {

// 	message := map[string]interface{}{
// 		"chat_id": "496818745",
// 		"text":    "Меня пересобрали...",
// 	}

// 	bytesRepresentation, err := json.Marshal(message)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	resp, err := http.Post("https://api.telegram.org/bot1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM/sendMessage", "application/json", bytes.NewBuffer(bytesRepresentation))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	res, err := http.Get("http://img10.joyreactor.cc/pics/post/full/%D0%9A%D0%9F%D0%91%D0%A1%D0%90-Anime-%D1%84%D1%8D%D0%BD%D0%B4%D0%BE%D0%BC%D1%8B-%D0%B0%D0%B3%D0%B8%D1%82%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B5-%D0%BF%D0%BB%D0%B0%D0%BA%D0%B0%D1%82%D1%8B-5580394.jpeg")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	fileByte, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	image := tgbotapi.FileBytes{Name: "build", Bytes: fileByte}
// 	_, err = bot.Send(tgbotapi.NewPhotoUpload(int64(496818745), image))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var result map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&result)

// 	log.Println(result)
// 	log.Println(result["data"])
// }
