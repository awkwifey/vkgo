package API

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/tildae/vkgo/tools/structurly"
)

// 💎 Options - configuration structure of the bot used to receive and send messages

type Options struct{
	Token		string
	Tokens		*map[int]string
	ID			int
	Version		float64
}

func Create() *Options {
	return &Options{}
}

func(bot *Options) Config(event string, option interface{}) *Options {
	typ := reflect.TypeOf(option).String()
	switch event {
		case "token":
			if typ == "*map[int]string" {
				bot.Tokens	= option.(*map[int]string)
			} else if typ == "*[]string" {
				var (
					tokens	= map[int]string{}
					pointer	= option.(*[]string)
				)

				for _, token := range *pointer {
					bot.Token = token
					group := bot.GetGroup(0).Parse().GetGroup()
					tokens[group.ID] = token
				}
				bot.Token	= ""
				bot.Tokens	= &tokens
			} else if typ == "string" {
				bot.Token	= option.(string)
			}
		case "version":
			bot.Version	= option.(float64)
	}

	return bot
}

// 🚀 Execute - method of sending messages to the APIVK for their processing

func(config *Options) Execute(method string, params interface{}) []byte {
	var (
		query	= structurly.ToURL(params)
		agent	= fiber.AcquireAgent()
		request	= agent.Request()
		link	= fmt.Sprintf("https://api.vk.com/method/%s?access_token=%s&v=%f&%v", method, config.Token, config.Version, query)
	)

	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(link)

	if err := agent.Parse(); err != nil {
		fmt.Println(err)
		return nil
	}

	_, body, _ := agent.Bytes()
	return body
}

// 💫 HeaderExecute - a method that works much faster than default due to the lack of structure processing

func(config *Options) HeaderExecute(method string, params string) []byte {
	var (
		agent	= fiber.AcquireAgent()
		request	= agent.Request()
		link	= fmt.Sprintf("https://api.vk.com/method/%s?access_token=%s&v=%f&%v", method, config.Token, config.Version, params)
	)

	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(link)

	err := agent.Parse()
	if err != nil {
		fmt.Println(err)
	}

	_, body, errs := agent.Bytes()
	if errs != nil {
		fmt.Println(errs)
	}
	return body
}