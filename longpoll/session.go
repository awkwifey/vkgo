package longpoll

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tildae/vkgo/longpoll/states"
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/scenes"
)

func(server *Longpoll) EndlessListen(events scenes.Scenes) {
	for {
		server.Listen(events)
	}
}

func(server *Longpoll) Listen(events scenes.Scenes) error {
	var (
		agent	= fiber.AcquireAgent()
		request	= agent.Request()
		bot		= server.Bot
		session	= server.Session

		link	= fmt.Sprintf("%s?act=a_check&key=%s&ts=%v&wait=%d", session.Server, session.Key, session.TS, server.Wait)
	)

	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(link)

	if err := agent.Parse(); err != nil {
		fmt.Println(err)
		return err
	}

	_, body, errs := agent.Bytes()
	if errs != nil {
		fmt.Println(errs)
		return errs[0]
	}

	var updates updates.Updates
	json.Unmarshal(body, &updates)

	state := states.States[updates.Failed]
	if state != nil {
		state(bot, session, updates, events)
	}
	return nil
}