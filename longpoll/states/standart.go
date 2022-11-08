package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

func standart(bot *API.Options, session *API.ResponseLongpollSession, params updates.Updates, events scenes.Scenes) {
	for _, update := range params.Updates {
		event := events[update.Type]
		if event != nil {
			event(bot, update)
		}
	}
	session.TS = params.TS
}
