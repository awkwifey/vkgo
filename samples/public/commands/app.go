package commands

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/addons/keyboard"
	"github.com/tildae/vkgo/updates"
)

var (
	appText		= "(?i)^(app|приложение)$"
	appPayload	= `{"button": "location"}`
)

func appScript(bot *API.Options, message updates.Message) {
	keyboard := keyboard.Create(false, true)
	keyboard.App("Example", 200822304, 6441755, locationPayload, "")

	text := bot.GetUser(message.UserID).Parse().Name() + ", pong!"
	bot.SendMessage(API.Message{
		ChatIDs:	message.ChatID,
		Text:		text,
		Keyboard:	keyboard.JSON(),
		Forward:	API.GetForward(message.ChatID, message.ID, true),
		RandomID:	API.GetRandom(),
	})
}