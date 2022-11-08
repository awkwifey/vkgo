package events

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/samples/public/commands"
	"github.com/tildae/vkgo/updates"
)

func Message(bot *API.Options, update updates.Update) {
	message	:= update.Object.Message

	if message.PayloadString != "" {
		message.Payload = message.PayloadString.Object()
		if message.Payload["button"] != nil {
			message.Text = message.Payload["button"].(string)
		}
	}

	commands.GetCommand(message.Text, bot, message)
}