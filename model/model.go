package model

import (
	"fmt"
	"os"
	"path"
	"qBittorrentBot/config"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	// Register some standard stuff
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// InitDB init db object
func InitDB() {
	connectDB()
	configDB()
	updateTable()
	log.Info("db init success")
}

func connectDB() {
	var err error
	db, err = gorm.Open("sqlite3", config.SQLitePath)
	if err != nil {
		log.Error(err.Error())
	}
}

// Disconnect disconnects from the database.
func Disconnect() {
	db.Close()
}

func configDB() {
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(config.DBLogMode)
	db.SetLogger(logger())
}

func updateTable() {
	createOrUpdateTable(&QBittorrent{})
}

// createOrUpdateTable create table or Migrate table
func createOrUpdateTable(model interface{}) {
	if !db.HasTable(model) {
		db.CreateTable(model)
	} else {
		db.AutoMigrate(model)
	}
}

func logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
