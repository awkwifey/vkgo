package API

type LongpollServer struct{
	Session	LongpollSession	`json:"response"`
	Error	Error		`json:"error"`
}
type LongpollSession struct{
	Server	string	`json:"server"`
	Key	string	`json:"key"`
	TS	interface{}	`json:"ts"`
}

type StreamingServer struct{
	Session	StreamingSession	`json:"response"`
	Error	Error			`json:"error"`
}
type StreamingSession struct{
	Endpoint	string	`json:"endpoint"`
	Key		string	`json:"key"`
}