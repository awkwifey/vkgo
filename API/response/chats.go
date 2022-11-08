package response

type Chats struct{
	Chats	Chat	`json:"response"`
	Error	Error	`json:"error"`
}

type Chat struct{
	ChatID	int	`json:"chat_id"`
	PeersID	[]int	`json:"peer_ids"`
	Link	string	`json:"link"`
	ChatInfo
}
type ChatInfo struct{
	ID	int	`json:"id"`
	Type	string	`json:"type"`
	Title	string	`json:"title"`
	Admin	int	`json:"admin_id"`
	Users	[]int	`json:"users"`
}