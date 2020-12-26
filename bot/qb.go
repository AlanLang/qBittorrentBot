package bot

import (
	"qBittorrentBot/model"

	"github.com/AlanLang/go-qbittorrent/qbt"
	log "github.com/sirupsen/logrus"
)

var qbClient *qbt.Client
var qbLinked = false

// InitQbClient 初始化
func InitQbClient(qb model.QBittorrent) error {
	if qbClient == nil {
		qbClient = qbt.NewClient(qb.URL)
	} else {
		qbClient.URL = qb.URL
	}
	islogin, err := qbClient.Login(qb.Username, qb.Password)
	if islogin {
		log.Info("qb login success")
		qbLinked = true
	}

	if err != nil {
		qbLinked = false
		log.Error("qb login failed", "error", err.Error())
	}
	return err
}
