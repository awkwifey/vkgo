package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

func outdated(bot *API.Options, session *API.ResponseLongpollSession, params updates.Updates, events scenes.Scenes) {
	session.TS = params.TS
}