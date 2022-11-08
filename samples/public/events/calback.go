package events

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
)

func Callback(bot *API.Options, update updates.Update) {
	message	:= update.Object

	if message.Payload["button"] == "ping" {
		bot.EditText(message.ChatID, message.ID, "OK!")
	}
}