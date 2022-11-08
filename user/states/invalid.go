package states

import (
	"fmt"

	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
)

func invalid(bot *API.Options, session *API.ResponseLongpollSession, updates Messages, events scenes.Scenes) {
	fmt.Printf("Invalid Version: \nMin%d\nMax: %d\n", updates.MinVersion, updates.MaxVersion)
	panic("")
}