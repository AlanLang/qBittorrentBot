package bot

import (
	"qBittorrentBot/model"
	"qBittorrentBot/qbt"

	log "github.com/sirupsen/logrus"
)

var qbClient *qbt.Client

// InitQbClient 初始化
func InitQbClient(qb model.QBittorrent) error {
	if qbClient == nil {
		qbClient = qbt.NewClient(qb.URL)
	} else {
		qbClient.URL = qb.URL
	}
	if !qbClient.IsLogin() {
		err := qbClient.Login(qb.Username, qb.Password)
		if err != nil {
			log.Error("qb login failed", "error", err.Error())
		}
		log.Info("qb login success")

		return err
	}
	return nil
}
