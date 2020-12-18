package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

//Person ...
type Person struct {
	chatID int
}

//Token ...
type Token struct {
	token int
}

var database *sql.DB

func main() {

	db, err := sql.Open("postgres", "postres: @/tg_bot")

	if err != nil {
		log.Println(err)
	}

	database = db
	defer db.Close()

	person := GetPerson()
	for _, pe := range person {
		fmt.Println(pe.chatID)
	}

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

//GetPerson ...
func GetPerson() (per []*Person) {

	rows, err := database.Query("select * from rg_bot.person")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	// person = []Person{}

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

	// for _, pe := range per {
	// 	fmt.Println(pe.chatID)
	// }

	return per
}

// //GetToken ...
// func GetToken() {

// 	rows, err := database.Query("select * from rg_bot.tg_main")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer rows.Close()
// 	token := []Token{}

// 	for rows.Next() {
// 		p := Token{}
// 		err := rows.Scan(&p.token)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		token = append(token, p)
// 	}

// 	return token
// }
