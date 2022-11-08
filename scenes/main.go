package scenes

import (
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
)

var Types = map[string]string{
	"message":			"message_new",
	"editMessage":		"message_edit",

	"replyMessage":		"message_reply",
	"enableMessage":	"message_allow",
	"disableMessage":	"message_deny",
	"callback":			"message_event",

	"typing":			"message_typing_state",
}

type Scene func(bot *API.Options, update updates.Update)
type Scenes map[string]Scene

func Create() Scenes {
	return Scenes{}
}

func(scene Scenes) Add(typ string, callback func(bot *API.Options, update updates.Update)) Scenes {
	customType := Types[typ]
	if customType == "" { customType = typ }
	scene[customType] = callback
	return scene
}