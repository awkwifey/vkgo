package templates

import (
	"fmt"
	"encoding/json"
)

type Template struct{
	Type		string		`json:"type"`
	Elements	[]Element	`json:"elements"`
}

func Create(typ string) *Template {
	return &Template{
		Type: typ,
		Elements: []Element{},
	}
}

func(template *Template) JSON() string {
	data, err := json.Marshal(template)
	if err != nil { fmt.Println(err) }

	return string(data)
}