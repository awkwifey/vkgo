package updates

type Callback struct {
	Text	string	`json:"message,omitempty"`
	ChatID	int		`json:"peer_id,omitempty"`
	UserID	int		`json:"user_id,omitempty"`
	EventID	string	`json:"event_id,omitempty"`
	Payload	map[string]interface{}	`json:"payload,omitempty"`
	ID		int		`json:"conversation_message_id,omitempty"`
}