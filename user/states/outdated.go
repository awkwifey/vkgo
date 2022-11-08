package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
)

func outdated(bot *API.Options, session *API.ResponseLongpollSession, updates Messages, events scenes.Scenes) {
	session.TS = updates.TS
}