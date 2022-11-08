package keyboard

import (
	"encoding/json"
	"fmt"
)

func(keyboard *Keyboard) Add(button Button) {
	rows := len(keyboard.Buttons)
	button.Color = colors[button.Color]

	if rows == 0 {
		keyboard.Row()
		keyboard.Buttons[0] = append(keyboard.Buttons[0], button)
	} else {
		keyboard.Buttons[rows - 1] = append(keyboard.Buttons[rows - 1], button)
	}
}

func(keyboard *Keyboard) JSON() string {
	data, err := json.Marshal(keyboard)
	if err != nil { fmt.Println(err) }

	return string(data)
}

func(keyboard *Keyboard) Row() {
	keyboard.Buttons = append(keyboard.Buttons, make([]Button, 0))
}

func(keyboard *Keyboard) Text(label, payload, color string) {
	keyboard.Add(Button{
		Action: Action{
			Type:		"text",
			Label:		label,
			Payload:	payload,
		},
		Color: color,
	})
}

func(keyboard *Keyboard) Callback(label, payload, color string) *Keyboard {
	keyboard.Add(Button{
		Action: Action{
			Type:		"callback",
			Label:		label,
			Payload:	payload,
		},
		Color: color,
	})
	return keyboard
}

func(keyboard *Keyboard) Link(label, link, payload string) *Keyboard {
	keyboard.Add(Button{
		Action: Action{
			Type:		"open_link",
			Label:		label,
			Link:		link,
			Payload:	payload,
		},
	})
	return keyboard
}

func(keyboard *Keyboard) App(label string, ownerID, appID int, payload, hash string) *Keyboard {
	keyboard.Add(Button{
		Action: Action{
			Type:		"open_app",
			Label:		label,
			OwnerID:	ownerID,
			AppID:		appID,
			Hash:		hash,
			Payload:	payload,
		},
	})
	return keyboard
}

func(keyboard *Keyboard) Location(payload string) *Keyboard {
	keyboard.Add(Button{
		Action: Action{
			Type:		"location",
			Payload:	payload,
		},
	})
	return keyboard
}