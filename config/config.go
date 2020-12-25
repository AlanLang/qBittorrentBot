package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// TelegraphToken TelegraphToken
	TelegraphToken string
	// SQLitePath sqlite文件地址
	SQLitePath string
	// DBLogMode DBLogMode
	DBLogMode bool
)

// Init 初始化配置
func Init() {
	log.Info("set config")
	viper.SetConfigFile(filepath.Join(getWorkPath(), "/config/config.yml"))
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Error(err)
		os.Exit(3)
	}
	TelegraphToken = viper.GetString("telegraph_token")
	if viper.IsSet("sqlite_path") {
		SQLitePath = viper.GetString("sqlite_path")
	} else {
		SQLitePath = filepath.Join(getWorkPath(), "/config/data.db")
	}
	if viper.IsSet("db_log") {
		DBLogMode = viper.GetBool("db_log")
	}
	log.Info("config is loaded")
}

func getWorkPath() string {
	WorkPath, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}
	return WorkPath
}
