package commands

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/addons/keyboard"
	"github.com/tildae/vkgo/updates"
)

var (
	locationText	= "(?i)^(location|местоположение)$"
	locationPayload	= `{"button": "location"}`
)

func locationScript(bot *API.Options, message updates.Message) {
	keyboard := keyboard.Create(false, true)
	keyboard.Location(locationPayload)

	text := bot.GetUser(message.UserID).Parse().Name() + ", pong!"
	bot.SendMessage(API.Message{
		ChatIDs:	message.ChatID,
		Text:		text,
		Keyboard:	keyboard.JSON(),
		Forward:	API.GetForward(message.ChatID, message.ID, true),
		RandomID:	API.GetRandom(),
	}).Parse()
}