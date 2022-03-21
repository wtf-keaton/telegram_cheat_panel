package main

import (
	"telegram_webpanel/internal/client_api"
	"telegram_webpanel/internal/dbApi"
	"telegram_webpanel/internal/telegram"
)

func main() {

	go telegram.HandleTelegram()
	go dbApi.HandleDB()

	mux := client_api.GetMultiplexer()

	mux.Run("localhost:1488")
}
