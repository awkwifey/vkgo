package commands

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
)

var (
	clearText = "(?i)^(_|delete)$"
)

func clearScript(bot *API.Options, message updates.Message) {
	var IDs []int
	for _, message := range message.Forwards {
		IDs = append(IDs, message.ID)
	}
	IDs = append(IDs, message.ID)

	bot.DeleteMessages(message.ChatID, IDs, 1).SendText("ok")
}