package main

import (
	"qBittorrentBot/bot"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("qBittorrentBot starting...")
	bot.ConfigInit()
	bot.Start()
}
