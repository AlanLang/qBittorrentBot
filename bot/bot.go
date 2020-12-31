package bot

import (
	"qBittorrentBot/bot/fsm"
	"qBittorrentBot/config"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	// UserState 用户状态，用于标示当前用户操作所在状态
	UserState map[int64]fsm.UserStatus = make(map[int64]fsm.UserStatus)
	// B telebot
	B *tb.Bot
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
		Token:  config.TelegraphToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Error(err)
		return
	}
	setCommands()
	setHandle()
	B.Start()
}

func setCommands() {
	// 设置bot命令提示信息
	commands := []tb.Command{
		tb.Command{Text: "start", Description: "开始使用"},
		tb.Command{Text: "list", Description: "查看正在下载的任务"},
		tb.Command{Text: "all", Description: "查看所有任务"},
		tb.Command{Text: "help", Description: "使用帮助"},
		tb.Command{Text: "config", Description: "配置qBittorrent服务器"},
	}

	if err := B.SetCommands(commands); err != nil {
		log.Error("set bot commands failed", "error", err.Error())
	}
}

func setHandle() {
	B.Handle(&tb.InlineButton{Unique: "qb_update_btn"}, updateQbCtr)
	B.Handle("/start", startCmdCtr)
	B.Handle("/list", listCmdCtr)
	B.Handle("/all", allCmdCtr)
	B.Handle("/help", helpCmdCtr)
	B.Handle("/config", configCmdCtr)
	B.Handle(tb.OnText, textCtr)
}
