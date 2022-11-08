package main

import (
	"os"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/scenes"
	"github.com/tildae/vkgo/samples/public/events"
	"github.com/tildae/vkgo/webhook"
)

func main() {
	bot := API.Create()
	bot.Config("version", 5.131)
	// Singlebot
	// bot.Config("token", os.Getenv("TOKEN"))
	// Multibot
	bot.Config("token", &[]string{os.Getenv("TOKEN"), os.Getenv("TOKEN2")})

	updates := scenes.Create()
	updates.Add("message", events.Message)

	webhook := webhook.Create(bot, updates)
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/vkontakte/:confirmation", webhook.Fiber)
	log.Fatal(app.Listen(":3000"))
}