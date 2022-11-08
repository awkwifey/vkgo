package API

import (
	"time"
	"encoding/json"
)

type ResponseMessages struct{
	Bot			*Options
	Body		[]byte
	Messages	[]ResponseMessage	`json:"response"`
	Error		Error				`json:"error"`
}
type ResponseMessage struct{
	ChatID		int			`json:"peer_id"`
	UserID		int			`json:"user_id"`
	ID			int			`json:"conversation_message_id"`
}

func(messages *ResponseMessages) Parse() *ResponseMessages {
	json.Unmarshal(messages.Body, &messages)
	return messages
}

func(messages *ResponseMessages) Wait(second int) *ResponseMessages {
	time.Sleep(time.Duration(second) * time.Second)
	return messages
}

func(messages *ResponseMessages) GetMessage() *ResponseMessage {
	for _, message := range messages.Messages {
		return &message
	}
	return &ResponseMessage{}
}

func(messages *ResponseMessages) SendText(text string) *ResponseMessages {
	for _, message := range messages.Messages {
		messages.Bot.SendText(message.ChatID, text)
	}
	return messages
}

func(messages *ResponseMessages) EditText(text string) *ResponseMessages {
	for _, message := range messages.Messages {
		messages.Bot.EditText(message.ChatID, message.ID, text)
	}
	return messages
}

func(messages *ResponseMessages) DeleteMessage(global int) *ResponseMessages {
	for _, message := range messages.Messages {
		messages.Bot.DeleteMessage(message.ChatID, message.ID, global)
	}
	return messages
}