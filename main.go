package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type respErr struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int32  `json:"error_code"`
	Description string `json:"description"`
}

type from struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type textMessage struct {
	MessageId int    `json:"message_id"`
	From      from   `json:"from"`
	Chat      chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

type result struct {
	UpgradeId int         `json:"upgrade_id"`
	Message   textMessage `json:"message"`
}

type response struct {
	Ok     bool     `json:"ok"`
	Result []result `json:"result"`
}

func main() {
	resp, code := getMethodTgAPI("getUpdates")

	if code == 404 {
		var errMsg respErr
		err := json.Unmarshal(resp, &errMsg)
		if err != nil {
			log.Fatalln(err)
		}
		log.Fatal(errMsg)
	}

	if code == 200 {
		var messageArray response
		err := json.Unmarshal(resp, &messageArray)
		log.Println(messageArray)
		if err != nil {
			log.Fatalln(err)
		}

	}

}

func getMethodTgAPI(method string) ([]byte, int) {

	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/%s", readToken(), method))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, resp.StatusCode
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
