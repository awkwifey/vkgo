package keyboard

import (
	"encoding/json"
	"fmt"
)

func(buttons *Buttons) Add(button Button) *Buttons {
	*buttons = append(*buttons, button)
	return buttons
}

func(buttons *Buttons) JSON() string {
	data, err := json.Marshal(buttons)
	if err != nil { fmt.Println(err) }

	return string(data)
}

/*  ====================
	= BUTTONS GENERATE =
	-    TYPE: TEXT    -
	==================== */

// Text - is a function...

func Text(label, payload, color string) *Buttons {
	return CreateButtons(Button{
		Action: Action{
			Type:		"text",
			Label:		label,
			Payload:	payload,
		},
		Color:			colors[color],
	})
}

// Text - is a method...

func(buttons *Buttons) Text(label, payload, color string) *Buttons {
	return buttons.Add(Button{
		Action: Action{
			Type:		"text",
			Label:		label,
			Payload:	payload,
		},
		Color:			colors[color],
	})
}

/*  ====================
	= BUTTONS GENERATE =
	-  TYPE: CALLBACK  -
	==================== */

// Callback - is a function...

func Callback(label, payload, color string) *Buttons {
	return CreateButtons(Button{
		Action: Action{
			Type:		"callback",
			Label:		label,
			Payload:	payload,
		},
		Color:			colors[color],
	})
}

// Callback - is a method...

func(buttons *Buttons) Callback(label, payload, color string) *Buttons {
	return buttons.Add(Button{
		Action: Action{
			Type:		"callback",
			Label:		label,
			Payload:	payload,
		},
		Color:			colors[color],
	})
}


/*  ====================
	= BUTTONS GENERATE =
	-    TYPE: LINK    -
	==================== */

// Link - is a function...

func Link(label, link, payload string) *Buttons {
	return CreateButtons(Button{
		Action: Action{
			Type:		"open_link",
			Label:		label,
			Link:		link,
			Payload:	payload,
		},
	})
}

// Link - is a method...

func(buttons *Buttons) Link(label, link, payload string) *Buttons {
	return buttons.Add(Button{
		Action: Action{
			Type:		"open_link",
			Label:		label,
			Link:		link,
			Payload:	payload,
		},
	})
}

/*  ====================
	= BUTTONS GENERATE =
	-     TYPE: APP    -
	==================== */

// App - is a function...

func App(label string, ownerID, appID int, payload, hash string) *Buttons {
	return CreateButtons(Button{
		Action: Action{
			Type:		"open_app",
			Label:		label,
			OwnerID:	ownerID,
			AppID:		appID,
			Hash:		hash,
			Payload:	payload,
		},
	})
}

// App - is a method...

func(buttons *Buttons) App(label string, ownerID, appID int, payload, hash string) *Buttons {
	return buttons.Add(Button{
		Action: Action{
			Type:		"open_app",
			Label:		label,
			OwnerID:	ownerID,
			AppID:		appID,
			Hash:		hash,
			Payload:	payload,
		},
	})
}

/*  ====================
	= BUTTONS GENERATE =
	-  TYPE: LOCATION  -
	==================== */

// Location - is a function...

func Location(payload string) *Buttons {
	return CreateButtons(Button{
		Action: Action{
			Type:		"location",
			Payload:	payload,
		},
	})
}

// Location - is a method...

func(buttons *Buttons) Location(payload string) *Buttons {
	return buttons.Add(Button{
		Action: Action{
			Type:		"location",
			Payload:	payload,
		},
	})
}
