package updates

type Object struct{
	Message	Message	`json:"message,omitempty"`
	Info	Info	`json:"client_info,omitempty"`
	Callback
	Typing
}
type Typing struct{
	State	string	`json:"state,omitempty"`
	UserID	int		`json:"from_id,omitempty"`
	GroupID	int		`json:"to_id,omitempty"`
}