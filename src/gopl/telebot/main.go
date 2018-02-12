package main

import (
	"net/http"
	"net/url"
)

func main() {
	var debug bool

	robot := newRobot("507339805:AAHh3XXpxLD7vAf_tGC3ictm0822E9Qj_0I",
		"ZenJingBot", "", Proxy())
	robot.bot.Debug = debug
	robot.run()
}

// // office demo
//func main() {
//
//	bot, err := tgbotapi.NewBotAPIWithClient("507339805:AAHh3XXpxLD7vAf_tGC3ictm0822E9Qj_0I", Proxy())
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates, err := bot.GetUpdatesChan(u)
//
//	for update := range updates {
//		if update.Message == nil {
//			continue
//		}
//		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//		msg.ReplyToMessageID = update.Message.MessageID   // 引用消息id
//		bot.Send(msg)
//	}
//}

func Proxy() *http.Client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	transport := &http.Transport{Proxy: proxy}

	return &http.Client{Transport: transport}
}
