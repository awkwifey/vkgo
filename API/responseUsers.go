package API

import (
	"encoding/json"
)

type ResponseUsers struct{
	Bot		*Options
	Body	[]byte
	Users	[]ResponseUser	`json:"response"`
	Error	Error	`json:"error"`
}
type ResponseUser struct{
	ID	int	`json:"id"`
	Name	string	`json:"first_name"`
	Surname	string	`json:"last_name"`
	Gender	int	`json:"sex"`
}

func(users *ResponseUsers) Parse() *ResponseUsers {
	json.Unmarshal(users.Body, &users)
	return users
}

func(users *ResponseUsers) Name() string {
	for _, user := range users.Users {
		return user.Name
	}
	return ""
}