package updates

import (
	"encoding/json"
)

type payload string
func(payload payload) Object() map[string]interface{} {
	object := make(map[string]interface{})
	json.Unmarshal([]byte(payload), &object)
	return object
}

type Message struct{
	Date		int		`json:"date,omitempty"`

	ChatID		int		`json:"peer_id,omitempty"`
	UserID		int		`json:"from_id,omitempty"`
	RandomID	int		`json:"random_id,omitempty"`
	MessageID	int		`json:"id,omitempty"`
	ID			int		`json:"conversation_message_id,omitempty"`

	Text		string		`json:"text,omitempty"`
	Payload		map[string]interface{}	`json:"payloadx,omitempty"`
	PayloadString	payload	`json:"payload,omitempty"`
	Attachments	[]Attachment`json:"attachments,omitempty"`
	Forwards	[]Message	`json:"fwd_messages,omitempty"`
	Reply		Reply		`json:"reply_message,omitempty"`

	Referal			string		`json:"ref,omitempty"`
	ReferalSource	string		`json:"ref_source,omitempty"`
	Geolocation		Geolocation	`json:"geo,omitempty"`
	Action			Action		`json:"action,omitempty"`
	Tag				string		`json:"message_tag,omitempty"`

	AuthorID	int			`json:"admin_author_id,omitempty"`
	Cropped		bool		`json:"is_cropped,omitempty"`
	Members		int			`json:"members_count,omitempty"`
	UpdatedAt	int			`json:"update_time,omitempty"`
	PinnedAt	int			`json:"pinned_at,omitempty"`
	Listened	bool		`json:"was_listened,omitempty"`
	Out			int			`json:"out,omitempty,omitempty"`
	Important	bool		`json:"important,omitempty"`
	Hidden		bool		`json:"is_hidden,omitempty"`
}

type Reply struct{
	Date	int	`json:"date,omitempty"`
	PeerID	int	`json:"peer_id,omitempty"`
	UserID	int	`json:"from_id,omitempty"`
	Text	int	`json:"text,omitempty"`
	ID	int	`json:"conversation_message_id,omitempty,omitempty"`
}

type Info struct{
	Buttons		[]string	`json:"button_actions,omitempty"`
	Keyboard	bool		`json:"keyboard,omitempty"`
	Inline		bool		`json:"inline,omitempty"`
	Carousel	bool		`json:"carousel,omitempty"`
	Lang		int			`json:"lang_id,omitempty"`
}