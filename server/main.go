package main

import (
	"server/bot"
	"server/http"
)

func main() {
	println("Start web server")
	go http.NewServer(8080)

	println("Start bot")
	bot.New("7076970913:AAH8jfzYiJmdPrFSPD0VwMDmbwVvKLuJHfM")
}
