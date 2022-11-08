package updates

type Geolocation struct{
	Type		string			`json:"type,omitempty"`
	Coordinates	[]Coordinate	`json:"coordinates,omitempty"`
	Place		Place			`json:"place,omitempty"`
	Showmap		string			`json:"showmap,omitempty"`
}

type Coordinate struct{
	Latitude	float64	`json:"latitude,omitempty"`
	Longitude	float64	`json:"longitude,omitempty"`
}

type Place struct{
	ID			int		`json:"id,omitempty"`
	Title		string	`json:"title,omitempty"`
	Latitude	int		`json:"latitude,omitempty"`
	Longitude	int		`json:"longitude,omitempty"`
	CreatedAt	int		`json:"created,omitempty"`
	Icon		string	`json:"icon,omitempty"`
	Country		string	`json:"country,omitempty"`
	City		string	`json:"city,omitempty"`
}