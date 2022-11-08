package API

import (
	"fmt"
)

func(config *Options) DeleteMessage(chatID, messageID, global int) *ResponseMessages {
	var (
		body		= config.HeaderExecute("messages.delete", fmt.Sprintf("peer_id=%d&cmids=%d&delete_for_all=%d", chatID, messageID, global))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}

func(config *Options) DeleteMessages(chatID int, messagesID []int, global int) *ResponseMessages {
	var ids string
	for _, id := range messagesID {
		ids += fmt.Sprintf("%d", id) + "%2C+"
	}

	var (
		body		= config.HeaderExecute("messages.delete", fmt.Sprintf("peer_id=%d&cmids=%s&delete_for_all=%d", chatID, ids, global))
		messages	= ResponseMessages{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}