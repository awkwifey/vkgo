package API

import (
	"fmt"
)

func(config *Options) GetMessageByID(messageID int) *ResponseMessagesByID {
	var (
		body		= config.HeaderExecute("messages.getById", fmt.Sprintf("message_ids=%d", messageID))
		messages	= ResponseMessagesByID{}
	)
	messages.Bot	= config
	messages.Body	= body

	return &messages
}