package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

func expired(bot *API.Options, session *API.ResponseLongpollSession, params updates.Updates, events scenes.Scenes) {
	newServer := bot.GetServerGroupLongpoll(bot.ID).Parse()
	session.Key = newServer.Session.Key
}