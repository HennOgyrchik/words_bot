package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Print("Hello")

	resp, err := http.Get("https://api.telegram.org/bot" + readToken() + "/getMe")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

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
