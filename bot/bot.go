package bot

import (
	"fmt"
	"github.com/catsworld/qq-bot-api"
	"log"
)

type Api interface {
	ReceivePrivateMsg(update qqbotapi.Update)

	ReceiveGroupMsg(update qqbotapi.Update)
	SendMsgToGroup(chatID int64, chatType string, msg string)
}

type BotAPI struct {
	Bot    *qqbotapi.BotAPI
	Update qqbotapi.Update
}

func NewBotAPI(debug bool) *BotAPI {
	accessToken := "MyCoolqHttpToken"
	listenHost := "http://localhost:5700"
	secret := "CQHTTP_SECRET"
	botapi, err := qqbotapi.NewBotAPI(accessToken, listenHost, secret)
	if err != nil {
		log.Fatal(err)
	}
	botapi.Debug = debug
	fmt.Println("qq机器人已启动")
	return &BotAPI{botapi, qqbotapi.Update{}}

}
func (b *BotAPI) MessageHandler() {

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = false
	updates := b.Bot.ListenForWebhook(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//	log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)
		//	PrintJson(update.Message.Message.CQString())
		b.bind(update)

		switch update.MessageType {
		case "private":
			b.ReceivePrivateMsg()
		case "group":
			b.ReceiveGroupMsg()
		}

	}
}
func (b *BotAPI) bind(update qqbotapi.Update) {
	b.Update = update
}
func (b *BotAPI) ReceivePrivateMsg() {
	ReceivePrivateMsg(b)

}

func (b *BotAPI) ReceiveGroupMsg() {
	ReceiveGroupMsg(b)

}

func (b *BotAPI) ReplyMsg(msg interface{}) {
	_, _ = b.Bot.SendMessage(b.Update.Message.Chat.ID, b.Update.Message.Chat.Type, msg)
}
func (b *BotAPI) SendMsgToGroup(chatID int64, msg interface{}) {
	_, _ = b.Bot.SendMessage(chatID, "group", msg)
}
func (b *BotAPI) SendMsgToGroups(msg interface{}, chatIDs ...int64) {
	for _, chatID := range chatIDs {
		_, _ = b.Bot.SendMessage(chatID, "group", msg)
	}

}
func (b *BotAPI) SendMsgToPrivate(chatID int64, msg interface{}) {
	_, _ = b.Bot.SendMessage(chatID, "private", msg)
}
