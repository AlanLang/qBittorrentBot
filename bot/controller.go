package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func startCmdCtr(m *tb.Message) {
	B.Send(m.Sender, "你好，欢迎使用QBittorrentBot")
}

func helpCmdCtr(m *tb.Message) {
	message := `
	命令：
	/start 开始使用
	/help 帮助
	`
	B.Send(m.Chat, message)
}
