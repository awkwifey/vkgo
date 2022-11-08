package response

type Users struct{
	Users	[]User	`json:"response"`
	Error	Error	`json:"error"`
}
type User struct{
	ID	int	`json:"id"`
	Name	string	`json:"first_name"`
	Surname	string	`json:"last_name"`
	Gender	int	`json:"sex"`
}