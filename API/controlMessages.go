package API

import (
	"fmt"
)

func(config *Options) PinMessage(chatID, messageID int) {
	config.HeaderExecute("messages.pin", fmt.Sprintf("peer_id=%d&conversation_message_id=%d", chatID, messageID))
}