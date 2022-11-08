package commands

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/addons/keyboard"
	"github.com/tildae/vkgo/updates"
)

var (
	pingText	= "(?i)^(ping|пинг|bot|бот)$"
	pingPayload	= `{"button": "ping"}`
)

func pingScript(bot *API.Options, message updates.Message) {
	keyboard := keyboard.Create(false, true)
	keyboard.Text("Повторить", pingPayload, "white")

	text := bot.GetUser(message.UserID).Parse().Name() + ", pong!"
	bot.SendMessage(API.Message{
		ChatIDs:	message.ChatID,
		Text:		text,
		Keyboard:	keyboard.JSON(),
		Forward:	API.GetForward(message.ChatID, message.ID, true),
		RandomID:	API.GetRandom(),
	}).Parse().EditText(text).Wait(10).DeleteMessage(1).GetMessage()
}