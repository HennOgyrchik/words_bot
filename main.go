package main

import (
	"fmt"
	. "words/jsonHandler"
)

func main() {

	resp := ParseResponse(GetMethodTgAPI("getUpdates"))
	fmt.Println(resp)

}
