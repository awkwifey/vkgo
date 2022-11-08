package API

import (
	"encoding/json"
)

type ResponseGroupsByID struct{
	Bot			*Options
	Body		[]byte
	Response	ResponseGroups	`json:"response"`
	Error		Error			`json:"error"`
}
type ResponseGroups struct{
	Groups	[]ResponseGroup	`json:"groups"`
}

type ResponseGroup struct{
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	ScreenName	string	`json:"screen_name"`
	Closed		bool	`json:"is_closed"`
	Type		string	`json:"type"`
}

func(groups *ResponseGroupsByID) Parse() *ResponseGroupsByID {
	json.Unmarshal(groups.Body, &groups)
	return groups
}

func(groups *ResponseGroupsByID) GetGroup() *ResponseGroup {
	for _, group := range groups.Response.Groups {
		return &group
	}
	return &ResponseGroup{}
}