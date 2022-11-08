package user

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tildae/vkgo/scenes"
	"github.com/tildae/vkgo/user/states"
)

func(server *Longpoll) EndlessListen(events scenes.Scenes) error {
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

		link	= fmt.Sprintf("https://%s?act=a_check&key=%s&ts=%v&wait=%d&mode=3&version=2", session.Server, session.Key, session.TS, server.Wait)
	)
	fmt.Println(link)

	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(link)

	if err := agent.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := agent.Bytes()
	if errs != nil {
		return errs[0]
	}

	var updates states.Messages
	json.Unmarshal(body, &updates)

	state := states.States[updates.Failed]
	if state != nil {
		state(bot, session, updates, events)
	}
	return nil
}