package webhook

import (
	"fmt"

	"github.com/tildae/vkgo/API"
	"github.com/tildae/vkgo/tools/colors"
	"github.com/tildae/vkgo/scenes"
)

type WebHook struct{
	Bot		*API.Options
	Scenes	scenes.Scenes
}

func Create(bot *API.Options, scenes scenes.Scenes) *WebHook {
	CallClear()
	var (
		mode	= "Unknown"
		meta	string
		tokens	= bot.Tokens
	)

	// Multibots
	if tokens != nil {
		mode = colors.Red(fmt.Sprintf("Multibot x%d", len(*tokens))).Bold().GetText()
	} else if bot.Token != "" {
		// Singlebots
		group := bot.GetGroup(0).Parse().GetGroup()
		bot.Config("token", &map[int]string{group.ID: bot.Token})
		mode = colors.Blue("Singlebot").Bold().GetText()
	} else {
		bot.Tokens = &map[int]string{}
	}

	if bot.Version != 0 {
		meta += fmt.Sprintf(
			"🕒 Версия API: %s\n💎 Режим: WebHook (%s)",
			colors.Magenta(fmt.Sprintf("%f", bot.Version)).Bold().GetText(), mode,
		)
	} else {
		fmt.Println("Отсутствует Version!")
		
	}

	fmt.Printf("%s 🎾\n\n%s\n%s\n\n%s\n",
		colors.Green("Сообщество запущено!").Bold().GetText(),
		botsCommunity,
		colors.Red("██").GetText() + colors.Green("██").GetText() + colors.Blue("██").GetText(),
		meta,
	)

	return &WebHook{
		Bot:	bot,
		Scenes:	scenes,
	}
}

func(webhook *WebHook) Config(typ string, option interface{}) *WebHook {
	switch typ {
		case "scenes":
			webhook.Scenes = option.(scenes.Scenes)
	}

	return webhook
}