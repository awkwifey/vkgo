package main

import (
	"os"
	"fmt"
	"github.com/tildae/vkgo/samples/public/events"

	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
	"github.com/tildae/vkgo/longpoll"
)

func main() {
	fmt.Println("Bot has been started!")

	longpoll.Create(API.Create().
		Config("version", 5.131).
		Config("token", os.Getenv("TOKEN")),
	).Config("wait", 80).
	EndlessListen(
		scenes.Create().
		Add("message", events.Message),
	)
}