package bot

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// TelegraphToken TelegraphToken
	TelegraphToken string
	// QBittorrentURL 地址
	QBittorrentURL string
	// QBittorrentName 用户名
	QBittorrentName string
	// QBittorrentPass 密码
	QBittorrentPass string
)

// ConfigInit 初始化配置
func ConfigInit() {
	log.Info("set config")
	viper.SetConfigFile("./config/config.yml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Error(err)
	}
	TelegraphToken = viper.GetString("telegraph_token")
	QBittorrentName = viper.GetString("username")
	QBittorrentPass = viper.GetString("password")
	QBittorrentURL = viper.GetString("url")
	log.Info("config is loaded")
}
