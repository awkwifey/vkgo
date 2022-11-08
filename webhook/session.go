package webhook

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/tildae/vkgo/updates"
)

func(webhook *WebHook) Fiber(ctx *fiber.Ctx) error {
	var (
		bot		= webhook.Bot
		tokens	= *bot.Tokens
		update	= updates.Update{}
	)
	json.Unmarshal(ctx.Body(), &update)
	bot.Config("token", tokens[update.GroupID])

	event := webhook.Scenes[update.Type]
	if event != nil { go event(bot, update) }

	if update.Type == "confirmation" {
		return ctx.SendString(ctx.Params("confirmation"))
	}
	return ctx.SendString("ok")
}