package API

import (
	"fmt"
	"time"
	"encoding/json"
)

type Message struct{
	ChatIDs			int			`strly:"peer_ids"`
	ChatID			int			`strly:"peer_id"`
	UserID			int			`strly:"from_id"`
	RandomID		int			`strly:"random_id"`
	ID				int			`strly:"conversation_message_id"`

	Text			string		`strly:"message"`
	Attachment		string		`strly:"attachment"`
	Template		string		`strly:"template"`
	Forward			string		`strly:"forward"`
	Keyboard		string		`strly:"keyboard"`

	KeepForward		int			`strly:"keep_forward_messages"`
	KeepSnippets	int			`strly:"keep_snippets"`
	DisableLinks	int			`strly:"dont_parse_links"`
	DisableMentions	int			`strly:"disable_mentions"`
}

type Forward struct{
	ChatID	int		`json:"peer_id"`
	ID		int		`json:"conversation_message_ids"`
	Reply	bool	`json:"is_reply"`
}

func GetRandom() int {
	return int(time.Now().Unix())
}

func GetForward(chatID, ID int, reply bool) string {
	text, err := json.Marshal(Forward{chatID, ID, reply})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(text)
}