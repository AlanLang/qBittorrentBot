package main

import (
	"os"
	"os/signal"
	"qBittorrentBot/bot"
	"qBittorrentBot/config"
	"qBittorrentBot/model"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("qBittorrentBot starting...")
	config.Init()
	model.InitDB()
	go handleSignal()
	bot.Start()
}

func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	<-c

	model.Disconnect()
	os.Exit(0)
}
