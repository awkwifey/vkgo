package updates

type Updates struct{
	TS		interface{}		`json:"ts"`
	Updates	[]Update		`json:"updates"`
	Failed	int				`json:"failed"`
}

type Update struct{
	Type	string	`json:"type"`
	Object	Object	`json:"object"`
	GroupID	int		`json:"group_id"`
	Version	string	`json:"v"`
	EventID	string	`json:"event_id"`
}