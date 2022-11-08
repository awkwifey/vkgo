package updates

type Action struct{
	Type	string	`json:"type,omitempty"`
	Member	int		`json:"member_id,omitempty"`
	Email	string	`json:"email,omitempty"`
	Photo	Photo	`json:"photo,omitempty"`
}

type Photo struct{
	X50		string	`json:"photo_50,omitempty"`
	X100	string	`json:"photo_100,omitempty"`
	X200	string	`json:"photo_200,omitempty"`
}