package commands

import (
	"fmt"

	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/addons/keyboard"
	"github.com/tildae/vkgo/addons/templates"
	"github.com/tildae/vkgo/updates"
)

var (
	carouselText = "(?i)^(carousel|каруселька)$"
)

func carouselScript(bot *API.Options, message updates.Message) {
	carousel := templates.Create("carousel")

	carousel.Element(
		"Title",
		"Description",
		carousel.Action("open_link", "https://vk.com/cdbot"),
		"-109837093_457242809",
		keyboard.App("Example", 200822304, 6441755, "{\"button\": \"application\"}", "").
		Link("Link", "https://vk.com/cdbot", "{\"button\": \"link\"}").
		Callback("Callback", "{\"button\": \"ping\"}", "white"),
	)

	carouselRes := bot.SendMessage(API.Message{
		ChatIDs:	message.ChatID,
		Text:		"пон",
		Template:	carousel.JSON(),
		Forward:	API.GetForward(message.ChatID, message.ID, true),
		RandomID:	API.GetRandom(),
	}).Parse()
	fmt.Println(carouselRes.Error)
}