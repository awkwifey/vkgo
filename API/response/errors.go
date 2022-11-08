package response

type Error struct{
	Code	int		`json:"error_code"`
	Message	string		`json:"error_msg"`
	Params	[]ErrorParams	`json:"request_params"`
}
type ErrorParams struct{
	Key	string	`json:"key"`
	Value	string	`json:"value"`
}