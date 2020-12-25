package bot

import (
	"qBittorrentBot/bot/fsm"
	"qBittorrentBot/model"

	log "github.com/sirupsen/logrus"
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
	/config 配置qBittorrent服务器
	`
	B.Send(m.Chat, message)
}

func configCmdCtr(m *tb.Message) {
	qbs, err := model.GetAllQb()
	if err != nil {
		log.Error(err)
	}
	if len(qbs) == 0 {
		addService(m)
	} else {
		unsubFeedItemBtns := [][]tb.InlineButton{}
		for _, qb := range qbs {
			unsubFeedItemBtns = append(unsubFeedItemBtns, []tb.InlineButton{
				tb.InlineButton{
					Unique: "qb_item_btn",
					Text:   qb.Name,
					Data:   qb.Name,
				},
			})
		}
		_, _ = B.Send(m.Chat, "请选择你要设置的服务器", &tb.ReplyMarkup{
			InlineKeyboard: unsubFeedItemBtns,
		})
	}
}

func addService(m *tb.Message) {
	B.Send(m.Sender, "开始配置qBittorrent服务器")
	addUserAction(m, "请自定义一个qBittorrent服务器的名称", fsm.AddQbName)
}

func textCtr(m *tb.Message) {
	log.Info(m.Text)
	switch UserState[m.Chat.ID] {
	case fsm.AddQbName:
		{
			name := m.Text
			qb := model.FineQb(name)
			if qb.Name != "" {
				addUserAction(m, "名称已存在，请换个名称吧", fsm.AddQbName)
				return
			}
			err := model.CreateQb(model.QBittorrent{
				Name: m.Text,
			})
			if err != nil {
				log.Error(err)
			}
			addUserAction(m, "名称设置成功，请继续设置qBittorrent服务器地址", fsm.AddQbPath)
		}
	}
}

func addUserAction(m *tb.Message, message string, action fsm.UserStatus) {
	_, err := B.Send(m.Chat, message, &tb.ReplyMarkup{ForceReply: true})
	if err == nil {
		UserState[m.Chat.ID] = action
	}
}
