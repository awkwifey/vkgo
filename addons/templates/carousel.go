package templates

import (
	"github.com/tildae/vkgo/addons/keyboard"
)

type Element struct{
	Title		string				`json:"title"`
	Description	string				`json:"description"`
	PhotoID		string				`json:"photo_id"`
	Action		Action				`json:"action"`
	Buttons		*keyboard.Buttons	`json:"buttons"`
}

type Action struct{
	Type	string	`json:"type"`
	Link	string	`json:"link,omitempty"`
}

func(template *Template) Add(element Element) {
	template.Elements = append(template.Elements, element)
}

func(template *Template) Action(typ, link string) Action {
	return Action{
		Type: typ,
		Link: link,
	}
}

func(template *Template) Element(title, desc string, action Action, photo string, buttons *keyboard.Buttons) {
	template.Add(Element{
		Title: title,
		Description: desc,
		Action: action,
		PhotoID: photo,
		Buttons: buttons,
	})
}