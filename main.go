package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

//Person ...
type Person struct {
	chatID int
}

var token string

var database *sql.DB

func main() {

	database = connectdb()

	bot, err := tgbotapi.NewBotAPI("1409094275:AAFBP7Vm2D-soxzHIht9pYXACDLOiJDqLdM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	SendMessage(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		reply := "Команда не найдена..."

		if update.Message.From.UserName != "Ell" {

			switch update.Message.Command() {
			case "hello":
				reply = "world"
			case "fendi":
				reply = "FENDI GUCHI FLIP-FLOP"
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}

//GetPerson ...
func GetPerson(database *sql.DB) (per []*Person) {

	rows, err := database.Query("select * from person")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	per = make([]*Person, 0)
	for rows.Next() {
		pe := new(Person)
		err := rows.Scan(&pe.chatID)
		if err != nil {
			fmt.Println(err)
			return
		}
		per = append(per, pe)
	}

	return per
}

//GetToken ...
func GetToken(database *sql.DB) (token string) {
	err := database.QueryRow("select token from tg_main").Scan(&token)
	if err != nil {
		log.Println(err)
	}

	return token
}

//SendMessage ...
func SendMessage(bot *tgbotapi.BotAPI) {
	person := GetPerson(database)

	arrChatID := make([]int, 0)
	for _, pe := range person {
		arrChatID = append(arrChatID, pe.chatID)

		message := map[string]interface{}{
			"chat_id": pe.chatID,
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
	for _, arr := range arrChatID {
		_, err = bot.Send(tgbotapi.NewPhotoUpload(int64(arr), image))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
