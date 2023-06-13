package main

import (
	"time"
	. "words/jsonHandler"
)

func main() {
	f := GetUpdates()

	for {

		f()
		time.Sleep(time.Second * 5)
	}

}
