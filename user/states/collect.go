package states

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
)

type Messages struct{
	TS	interface{}	`json:"ts,omitempty"`
	Updates	[][]interface{}	`json:"updates,omitempty"`
	Failed	int	`json:"failed,omitempty"`
	MaxVersion	int	`json:"max_version"`
	MinVersion	int	`json:"min_version"`
}

type state func(bot *API.Options, session *API.ResponseLongpollSession, updates Messages, events scenes.Scenes)

var States = map[int]state{
	0: standart,
	1: outdated,
	2: expired,
	3: lost,
	4: invalid,
}