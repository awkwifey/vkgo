package main

import (
	"os"

	"github.com/tildae/vkgo/samples/public/events"
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
	"github.com/tildae/vkgo/user"
)

func main() {
	bot := API.Create()
	bot.Config("token", os.Getenv("USER_TOKEN"))
	bot.Config("version", 5.131)

	updates := scenes.Create()
	updates.Add("message", events.Message)

	session := user.Create(bot)
	session.Config("wait", 80)
	session.EndlessListen(updates)
}