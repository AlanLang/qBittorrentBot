package bot

import (
	"time"

	"github.com/AlanLang/go-qbittorrent/qbt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	// B telebot
	B  *tb.Bot
	qb *qbt.Client
)

// Commands 命令集合
type Commands struct {
	start   string
	help    string
	version string
}

// Start 开始
func Start() {
	var err error
	B, err = tb.NewBot(tb.Settings{
		Token:  TelegraphToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Error(err)
		return
	}
	qbInit()
	setCommands()
	setHandle()
	B.Start()
}

func setCommands() {
	// 设置bot命令提示信息
	commands := []tb.Command{
		tb.Command{Text: "start", Description: "开始使用"},
		tb.Command{Text: "help", Description: "使用帮助"},
	}

	if err := B.SetCommands(commands); err != nil {
		log.Error("set bot commands failed", "error", err.Error())
	}
}

func setHandle() {
	B.Handle("/start", startCmdCtr)
	B.Handle("/help", helpCmdCtr)
}

func qbInit() {
	qb = qbt.NewClient(QBittorrentURL)
	islogin, err := qb.Login(QBittorrentName, QBittorrentPass)
	if islogin {
		log.Info("qb login success")
	}

	if err != nil {
		log.Error("qb login failed", "error", err.Error())
	}
}
