package user

import (
	"github.com/tildae/vkgo/API"
)

type Longpoll struct{
	Bot		*API.Options
	Session	*API.ResponseLongpollSession
	Wait	int
}

func Create(bot *API.Options) *Longpoll {
	server := bot.GetServerUserLongpoll().Parse()

	return &Longpoll{
		Bot:		bot,
		Session:	server.Session,
	}
}

func(longpoll *Longpoll) Config(typ string, option interface{}) *Longpoll {
	switch typ {
		case "wait":
			longpoll.Wait = option.(int)
	}

	return longpoll
}