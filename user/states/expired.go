package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
)

func expired(bot *API.Options, session *API.ResponseLongpollSession, updates Messages, events scenes.Scenes) {
	newSession := bot.GetServerGroupLongpoll(bot.ID)
	session.Key = newSession.Session.Key
}