package jsonHandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RespErr struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int32  `json:"error_code"`
	Description string `json:"description"`
}

type From struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type TextMessage struct {
	MessageId int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

type Result struct {
	UpdateId int         `json:"update_id"`
	Message  TextMessage `json:"message"`
}

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

// Разборка json ответа на структуры
func ParseResponse(resp *http.Response) Response {
	jsonResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	switch resp.StatusCode {
	case 200:
		{
			var messageArray Response
			err = json.Unmarshal(jsonResponse, &messageArray)
			log.Println(messageArray)
			if err != nil {
				log.Fatal(err)
			}
			return messageArray
		}
	case 404:
		log.Fatal("404") //подумать над обработкой
	}

	return Response{}
}

func GetMethodTgAPI(method string) *http.Response {

	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/%s", readToken(), method))
	if err != nil {
		log.Fatalln(err)
	}
	return resp

}

func readToken() string {
	file, err := os.Open("token.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		return string(data[:n])
	}
	return "oops"
}

func PostMethodTgAPI(method string, data any) *http.Response {
	q, err := json.Marshal(data)
	r := bytes.NewReader(q)

	resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/%s", readToken(), method), "application/json", r)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func UpdateOffset(updateId int) {
	test := ParseResponse(PostMethodTgAPI("getUpdates", struct {
		Offset int `json:"offset"`
	}{updateId}))

	fmt.Println(test)
}
