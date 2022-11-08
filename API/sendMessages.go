package API

import (
	"fmt"
	"net/url"
	"strconv"
)

func(config *Options) SendMessage(message Message) *ResponseMessages {
	var (
		body		= config.Execute("messages.send", message)
		messages	ResponseMessages
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) SendText(chatID int, text string) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.send", fmt.Sprintf("peer_ids=%d&message=%s&disable_mentions=1&random_id=0", chatID, url.QueryEscape(text)))
		messages	ResponseMessages
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) SendForwardedText(chatID, messageID int, text string) *ResponseMessages {
	var (
		chat		= strconv.Itoa(chatID)
		body		= config.HeaderExecute("messages.send", "peer_ids=" + chat + "&message=" + url.QueryEscape(text) + "&disable_mentions=1&random_id=0&forward=%7B%22conversation_message_ids%22%3A" + strconv.Itoa(messageID) + "%2C%22is_reply%22%3Atrue%2C%22peer_id%22%3A" + chat + "%7D")
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) SendPhoto(chatID, ownerID, photoID int) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.send", fmt.Sprintf("peer_id=%d&attachment=photo%d_%d&disable_mentions=1&random_id=0", chatID, ownerID, photoID))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) SendSticker(chatID, stickerID int) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.send", fmt.Sprintf("peer_id=%d&sticker_id=%d&random_id=0", chatID, stickerID))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}