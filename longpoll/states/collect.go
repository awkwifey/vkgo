package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

type state func(bot *API.Options, server *API.ResponseLongpollSession, params updates.Updates, events scenes.Scenes)

var States = map[int]state{
	0: standart,
	1: outdated,
	2: expired,
	3: lost,
}