package API

import (
	"encoding/json"
	"github.com/tildae/vkgo/updates"
)

type ResponseMessagesByID struct{
	Bot			*Options
	Body		[]byte
	Messages	ResponseMessageByID	`json:"response"`
	Error		Error				`json:"error"`
}
type ResponseMessageByID struct{
	Count	int					`json:"count"`
	Items	[]updates.Message	`json:"items"`
}

func(messages *ResponseMessagesByID) Parse() *ResponseMessagesByID {
	json.Unmarshal(messages.Body, &messages)
	return messages
}

func(messages *ResponseMessagesByID) GetMessage() updates.Message {
	return messages.Messages.Items[0]
}