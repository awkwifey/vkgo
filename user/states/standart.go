package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

var typesEvent = map[int]func(bot *API.Options, message []interface{}, update *updates.Update) string {
	4: message,
}

func standart(bot *API.Options, session *API.ResponseLongpollSession, messages Messages, events scenes.Scenes) {
	for _, message := range messages.Updates {
		update := updates.Update{}

		evtype := typesEvent[int(message[0].(float64))]
		if evtype == nil { continue; }
		typ := evtype(bot, message, &update)
		event := events[typ]
		if event != nil {
			go event(bot, update)
		}
	}
	session.TS = messages.TS
}

func message(bot *API.Options, message []interface{}, update *updates.Update) string {
	if message[1] != 0 {
		update.Object.Message.MessageID = int(message[1].(float64))
	}

	update.Object.Message = bot.GetMessageByID(update.Object.Message.MessageID).Parse().GetMessage()

	return "message_new"
}