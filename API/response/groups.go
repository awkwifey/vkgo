package response

type GroupsByID struct{
	Response	Groups	`json:"response"`
	Error		Error	`json:"error"`
}

type Groups struct{
	Groups	[]Group `json:"groups"`
}
type Group struct{
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	ScreenName	string	`json:"screen_name"`
	Closed		bool	`json:"is_closed"`
	Type		string	`json:"type"`
}