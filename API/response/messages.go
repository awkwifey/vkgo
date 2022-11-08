package response

import "github.com/tildae/vkgo/updates"

type Messages struct{
	Messages	[]Message	`json:"response"`
	Error		Error		`json:"error"`
}
type Message struct{
	ChatID		int			`json:"peer_id"`
	UserID		int			`json:"user_id"`
	ID			int			`json:"conversation_message_id"`
}

type MessagesByID struct{
	Messages	MessageByID	`json:"response"`
	Error		Error		`json:"error"`
}
type MessageByID struct{
	Count	int		`json:"count"`
	Items	[]updates.Message	`json:"items"`
}