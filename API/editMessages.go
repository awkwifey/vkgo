package API

import (
	"fmt"
	"net/url"
)

func(config *Options) EditMessage(message Message) *ResponseMessages {
	var (
		body		= config.Execute("messages.edit", message)
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) EditText(chatID, messageID int, text string) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.edit", fmt.Sprintf("peer_id=%d&conversation_message_id=%d&message=%s&keep_forward_messages=1", chatID, messageID, url.QueryEscape(text)))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) EditPhoto(chatID, messageID, ownerID, photoID int) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.edit", fmt.Sprintf("peer_id=%d&conversation_message_id=%d&attachment=photo%d_%d", chatID, messageID, ownerID, photoID))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}