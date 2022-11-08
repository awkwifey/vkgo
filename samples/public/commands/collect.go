package commands

import (
	"regexp"

	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/updates"
)

type command func(bot *API.Options, update updates.Message)
var commands = map[string]command {
	pingText:		pingScript,
	clearText:		clearScript,
	carouselText:	carouselScript,
	locationText:	locationScript,
	appText:		appScript,
}

func GetCommand(text string, bot *API.Options, message updates.Message) {
	for regular, command := range commands {
		locate, _ := regexp.MatchString(regular, text)
		if locate == true {
			command(bot, message)
		}
	}
}