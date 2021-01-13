package bot

import (
	"qBittorrentBot/model"

	"github.com/AlanLang/qbt"
)

var qbClient *qbt.Client

// InitQbClient 初始化
func InitQbClient(qb model.QBittorrent) error {
	if qbClient == nil {
		qbClient, _ = qbt.New(qb.URL, qb.Username, qb.Password)
	} else {
		qbClient.SetConfig(qbt.Config{
			URL: qb.URL,
		})
		// qbClient.URL = qb.URL
	}
	return nil
}

func getDownloadList(qb model.QBittorrent) (map[string]qbt.Torrent, error) {
	err := InitQbClient(qb)
	if err != nil {
		return nil, err
	}
	s, e := qbClient.Sync("0")
	return s.Torrents, e
}

func download(qb model.QBittorrent, url string) error {
	err := InitQbClient(qb)
	if err != nil {
		return err
	}
	return qbClient.Add(url)
}
