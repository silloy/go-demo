package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Robot struct {
	bot      *tgbotapi.BotAPI
	updates  <-chan tgbotapi.Update //update msg
	shutUp   bool                   //shut up the robot
	name     string                 //name from telegram
	nickName string                 //user defined name
}

// Action used in interaction mode
type Action struct {
	ActionName string
	ActionStep int
}

const Markdown = tgbotapi.ModeMarkdown

var responseMessage string

func newRobot(token, nickName, webHook string, client *http.Client) *Robot {
	var rb = new(Robot)
	var err error
	if client != nil {
		rb.bot, err = tgbotapi.NewBotAPIWithClient(token, client)
	} else {
		rb.bot, err = tgbotapi.NewBotAPI(token)
	}
	if err != nil {
		log.Fatal(err)
	}
	rb.name = rb.bot.Self.UserName //name from telegram
	rb.nickName = nickName         //name from yourself
	log.Printf("%s: Authorized on account %s", rb.nickName, rb.name)

	time.Sleep(time.Millisecond * 500)
	if webHook != "" {
		_, err = rb.bot.SetWebhook(tgbotapi.NewWebhook(webHook + rb.bot.Token))
		if err != nil {
			log.Fatal(err)
		}
		rb.updates = rb.bot.ListenForWebhook("/" + rb.bot.Token)
	} else {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		rb.updates, err = rb.bot.GetUpdatesChan(u)
	}
	return rb
}

func (rb *Robot) run() {
	//chatId := conn.ReadMasterId()
	//rawMsg := fmt.Sprintf("%s is coming back!", rb.nickName)
	//rb.Reply(chatId, rawMsg)
	//go loginZMZ()
	//reload tasks from redis
	//go restoreTasks(rb)  FIXME
	for update := range rb.updates {
		go handlerUpdate(rb, update)
	}
}

//all telegram updates are handled here
func handlerUpdate(rb *Robot, update tgbotapi.Update) {
	j, _ := json.Marshal(update)
	log.Println(string(j))

	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			log.Println(err)
		}
	}()


	if update.Message == nil && update.CallbackQuery == nil {
		return
	}

	if update.Message != nil {
		text := update.Message.Text
		log.Println("text: " + text)
		text = strings.Replace(text, "@"+rb.name, "", 1)
		var rawMsg string
		//received := strings.Split(text, " ")
		if update.Message.IsCommand() {
			rawMsg = inCommand(rb, update.Message.Command(), update)
		} else {
			rawMsg = rb.Talk(update)
		}

		if rawMsg == "" {
			return
		}
		if err := rb.Reply(update, rawMsg); err != nil {
			panic(err)
		}
	}

	if update.CallbackQuery != nil {
		var editMsg string
		data := update.CallbackQuery.Data
		log.Println("data: " + data)
		editMsg = inCallbackData(rb, data, update)
		if editMsg == "" {
			return
		}
		if err := rb.ReplyEditMsg(update, editMsg); err != nil {
			panic(err)
		}
	}
}


func inCommand(rb *Robot, endPoint string, update tgbotapi.Update) (rawMsg string) {
	switch endPoint {
	case "start":
		rawMsg = rb.Start(update)
	case "help":
		rawMsg = "type /sayhi or /status."
	case "sayhi":
		rawMsg = "Hi :)"
	case "status":
		rawMsg = "I'm ok."
	default:
		rawMsg = "I don't know that command"
	}
	return
}

func inCallbackData(rb *Robot, endPoint string, update tgbotapi.Update) (rawMsg string) {
	switch endPoint {
	case "start":
		rawMsg = rb.Start(update)
	case "AAA":
		rawMsg = "type /sayhi or /status."
	case "HELLO":
		rawMsg = "Hi :)"
	case "status":
		rawMsg = "I'm ok."
	default:
		rawMsg = "I don't know that command"
	}
	return
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)


func stringCallbackData(cmd string, matchID uint64) string {
	data := InlineCallbackData{Command: cmd, MatchID: matchID}
	b, _ := json.Marshal(data)
	return string(b)
}

type InlineCallbackData struct {
	Command string `json:"c"`
	MatchID uint64 `json:"m"`
}

// Talk with AI
func (rb *Robot) Talk(update tgbotapi.Update) string {
	info := update.Message.Text
	if strings.Contains(info, rb.name) {
		if strings.Contains(info, "闭嘴") || strings.Contains(info, "别说话") {
			rb.shutUp = true
		} else if rb.shutUp && strings.Contains(info, "说话") {
			rb.shutUp = false
			return fmt.Sprintf("%s终于可以说话啦", rb.nickName)
		}
		info = strings.Replace(info, fmt.Sprintf("@%s", rb.name), "", -1)
	}

	if rb.shutUp {
		return ""
	}
	log.Printf(info)

	log.Println(rb.nickName, rb.name)
	if rb.nickName != "ABC" {
		if chinese(info) {
			return tlAI(info)
		}
		return mitAI(info)
	} else {
		//jarvis use another AI
		return qinAI(info)
	}
}

func (rb *Robot) Start(update tgbotapi.Update) string {
	user := update.Message.Chat.UserName
	//go conn.CreateUserChatId(user, update.Message.Chat.ID)
	return fmt.Sprintf("welcome: %s.\nType '/help' see what can I do.", user)
}

// Reply is encapsulated robot message send action
func (rb *Robot) ReplyEditMsg(update tgbotapi.Update, rawMsg string) (err error) {
	inlineKeyboard := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Welcome", "HELLO"),
		tgbotapi.NewInlineKeyboardButtonData("Emoclew", "AAA"),
	}

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)


	e := tgbotapi.NewEditMessageText(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		update.CallbackQuery.Message.Text,
	)
	e.BaseEdit.ReplyMarkup = &inlineKeyboardMarkup
	e.ParseMode = Markdown

	_, err = rb.bot.Send(e)
	if err != nil {
		log.Println(err)
	}
	return

}

// Reply is encapsulated robot message send action
func (rb *Robot) Reply(update tgbotapi.Update, rawMsg string) (err error) {
	ChatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(ChatID, rawMsg)

	switch update.Message.Text {
	case "open":
		msg.ReplyMarkup = numericKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	inlineButton := tgbotapi.NewInlineKeyboardButtonData("测试一下", "AAA")
	inlineKeyboard := []tgbotapi.InlineKeyboardButton{
		inlineButton,
	}

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)
	msg.ReplyMarkup = inlineKeyboardMarkup

	msg.ParseMode = "markdown"
	log.Printf(rawMsg)
	_, err = rb.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return
}
