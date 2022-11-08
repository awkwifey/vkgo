package API

import (
	"encoding/json"
)

type ResponseLongpollServer struct{
	Bot		*Options
	Body	[]byte
	Session	*ResponseLongpollSession	`json:"response"`
	Error	Error		`json:"error"`
}
type ResponseLongpollSession struct{
	Server	string		`json:"server"`
	Key		string		`json:"key"`
	TS		interface{}	`json:"ts"`
}

func(server *ResponseLongpollServer) Parse() *ResponseLongpollServer {
	json.Unmarshal(server.Body, &server)
	return server
}