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
	userID := m.Chat.ID
	qb := model.FineQb(userID)
	if qb.User != 0 {
		qb := model.FineQb(userID)
		message := "您已配置qBittorrent服务器"
		message += "\n地址：" + qb.URL
		if qb.Username == "" {
			message += "\n用户名：[未定义]"
		} else {
			message += "\n用户名：" + qb.Username
		}
		if qb.Password == "" {
			message += "\n密码：[未定义]"
		} else {
			message += "\n密码：********"
		}
		updateQbActionBtns := [][]tb.InlineButton{}
		updateQbActionBtns = append(updateQbActionBtns, []tb.InlineButton{
			tb.InlineButton{
				Unique: "qb_update_btn",
				Text:   "修改地址",
				Data:   fsm.ChangeQbURLBtn,
			},
		})
		updateQbActionBtns = append(updateQbActionBtns, []tb.InlineButton{
			tb.InlineButton{
				Unique: "qb_update_btn",
				Text:   "修改用户名",
				Data:   fsm.ChangeQbUserBtn,
			},
		})
		updateQbActionBtns = append(updateQbActionBtns, []tb.InlineButton{
			tb.InlineButton{
				Unique: "qb_update_btn",
				Text:   "修改密码",
				Data:   fsm.ChangeQbPassBtn,
			},
		})
		_, _ = B.Send(m.Chat, message, &tb.ReplyMarkup{
			InlineKeyboard: updateQbActionBtns,
		})
	} else {
		B.Send(m.Sender, "开始配置qBittorrent服务器")
		addUserAction(m, "请设置qBittorrent服务器地址", fsm.AddQbURL)
	}
}

func updateQbCtr(c *tb.Callback) {
	_ = B.Delete(c.Message)
	switch c.Data {
	case fsm.ChangeQbURLBtn:
		{
			addUserAction(c.Message, "请输入qBittorrent服务器地址", fsm.ChangeQbURL)
		}
	case fsm.ChangeQbUserBtn:
		{
			addUserAction(c.Message, "请输入qBittorrent服务器用户名", fsm.ChangeQbUser)
		}
	case fsm.ChangeQbPassBtn:
		{
			addUserAction(c.Message, "请输入qBittorrent服务器密码", fsm.ChangeQbPass)
		}
	}
}

func textCtr(m *tb.Message) {
	switch UserState[m.Chat.ID] {
	case fsm.AddQbURL:
		{
			err := setQbURL(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent服务器地址", fsm.AddQbURL)
				return
			}
			B.Delete(m)
			addUserAction(m, "名称设置成功，请继续设置qBittorrent用户名", fsm.AddQbUser)
		}
	case fsm.AddQbUser:
		{
			err := setQbUser(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent用户名", fsm.AddQbUser)
				return
			}
			B.Delete(m)
			addUserAction(m, "设置成功，请继续设置qBittorrent密码", fsm.AddQbPass)
		}
	case fsm.AddQbPass:
		{
			err := setQbPass(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent密码", fsm.AddQbPass)
				return
			}
			B.Delete(m)
			B.Send(m.Sender, "qBittorrent服务器配置成功")
		}
	case fsm.ChangeQbURL:
		{
			err := setQbURL(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent服务器地址", fsm.ChangeQbURL)
				return
			}
			B.Delete(m)
			B.Send(m.Sender, "qBittorrent服务器地址修改成功")
		}
	case fsm.ChangeQbUser:
		{
			err := setQbUser(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent服务器用户名", fsm.ChangeQbUser)
				return
			}
			B.Delete(m)
			B.Send(m.Sender, "qBittorrent服务器用户名修改成功")
		}
	case fsm.ChangeQbPass:
		{
			err := setQbPass(m.Chat.ID, m.Text)
			if err != nil {
				log.Error(err)
				addUserAction(m, "设置失败，请重新输入qBittorrent密码", fsm.ChangeQbPass)
				return
			}
			B.Delete(m)
			B.Send(m.Sender, "qBittorrent服务器密码修改成功")
		}
	default:
		{
		}
	}
}

func setQbURL(userID int64, URL string) error {
	qb := model.FineQb(userID)
	var err error
	if qb.User != 0 {
		qb.URL = URL
		err = model.UpdateQb(qb)
	} else {
		err = model.CreateQb(model.QBittorrent{
			User: userID,
			URL:  URL,
		})
	}
	return err
}

func setQbUser(userID int64, user string) error {
	qb := model.FineQb(userID)
	var err error
	if qb.User != 0 {
		qb.Username = user
		err = model.UpdateQb(qb)
	} else {
		err = model.CreateQb(model.QBittorrent{
			User:     userID,
			Username: user,
		})
	}
	return err
}

func setQbPass(userID int64, password string) error {
	qb := model.FineQb(userID)
	var err error
	if qb.User != 0 {
		qb.Password = password
		err = model.UpdateQb(qb)
	} else {
		err = model.CreateQb(model.QBittorrent{
			User:     userID,
			Password: password,
		})
	}
	return err
}

func addUserAction(m *tb.Message, message string, action fsm.UserStatus) {
	_, err := B.Send(m.Chat, message, &tb.ReplyMarkup{ForceReply: true})
	if err == nil {
		UserState[m.Chat.ID] = action
	}
}
