package main

import (
	"time"
	. "words/jsonHandler"
)

func main() {
	for {
		resp := ParseResponse(GetMethodTgAPI("getUpdates"))
		//fmt.Println(resp)

		x := len(resp.Result) - 1
		if x >= 0 {
			UpdateOffset(resp.Result[x].UpdateId + 1)
		}

		time.Sleep(time.Second * 5)
	}

}
