package main

import (
	"flag"
	"server/bot"
	"server/helper"
	"server/http"
)

func main() {
	hostFlag := flag.String("host", "127.0.0.1", "a string")
	flag.Parse()

	println(*hostFlag)
	if hostFlag != nil {
		helper.SERVER_IP = *hostFlag
	}

	println("Start web server")
	go http.NewServer(8080)

	println("Start bot")
	bot.New("7076970913:AAH8jfzYiJmdPrFSPD0VwMDmbwVvKLuJHfM")
}
