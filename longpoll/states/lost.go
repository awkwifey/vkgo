package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

func lost(bot *API.Options, session *API.ResponseLongpollSession, params updates.Updates, events scenes.Scenes) {
	getServer := bot.GetServerGroupLongpoll(bot.ID).Parse()
	session.Key	= getServer.Session.Key
	session.TS	= getServer.Session.TS
}