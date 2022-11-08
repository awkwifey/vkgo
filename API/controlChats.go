package API
/*
import (
	"encoding/json"
	"fmt"
	"net/url"
)

func(config *Options) CreateChat(name string, usersID []int) (response.Chat, error) {
	var ids string
	for _, id := range usersID {
		ids += fmt.Sprintf("%d", id) + "%2C+"
	}

	body := config.HeaderExecute("messages.createChat", fmt.Sprintf("user_ids=%s&title=%s", ids, url.QueryEscape(name)))

	var chats response.Chats
	json.Unmarshal(body, &chats)

	if chats.Error.Code != 0 {
		message := chats.Error.Message
		return response.Chat{}, fmt.Errorf(message)
	} else {
		return chats.Chats, nil
	}
}

func(config *Options) GetChat(chatID int) (response.Chat, error) {
	body := config.HeaderExecute("messages.getChat", fmt.Sprintf("chat_id=%d", chatID))

	var chats response.Chats
	json.Unmarshal(body, &chats)

	if chats.Error.Code != 0 {
		message := chats.Error.Message
		return response.Chat{}, fmt.Errorf(message)
	} else {
		return chats.Chats, nil
	}
}

func(config *Options) GetChatLink(chatID, reset int) (string, error) {
	body := config.HeaderExecute("messages.getInviteLink", fmt.Sprintf("peer_id=%d&reset=%d", 2e9 + chatID, reset))

	var chats response.Chats
	json.Unmarshal(body, &chats)

	if chats.Error.Code != 0 {
		message := chats.Error.Message
		return "", fmt.Errorf(message)
	} else {
		return chats.Chats.Link, nil
	}
}*/