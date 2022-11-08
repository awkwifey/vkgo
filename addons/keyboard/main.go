package keyboard

type Keyboard struct {
	OneTime	bool		`json:"one_time,omitempty"`
	Inline	bool		`json:"inline,omitempty"`
	Buttons	[]Buttons	`json:"buttons"`
}

type Buttons []Button
type Button struct {
	Action	Action	`json:"action,omitempty"`
	Color	string	`json:"color,omitempty"`
}
type Action struct {
	Type	string	`json:"type,omitempty"`
	Link	string	`json:"link,omitempty"`
	Label	string	`json:"label,omitempty"`

	OwnerID	int		`json:"owner_id,omitempty"`
	AppID	int		`json:"app_id,omitempty"`
	Hash 	string	`json:"hash,omitempty"`

	Payload	string	`json:"payload,omitempty"`
}

var colors = map[string]string{
	"white":	"default",
	"blue":		"primary",
	"green":	"positive",
	"red":		"negative",
}

func Create(oneTime, inline bool) *Keyboard {
	return &Keyboard{
		OneTime: oneTime,
		Inline: inline,
		Buttons: make([]Buttons, 0, 2),
	}
}

func CreateButtons(button Button) *Buttons {
	return &Buttons{button}
}