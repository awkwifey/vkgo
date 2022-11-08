package colors

type colour struct{
	text	string
	color	string
}

func(clr *colour) Bold() *colour {
	clr.color = clr.color + bold
	return clr
}

func(clr *colour) Underline() *colour {
	clr.color = clr.color + underline
	return clr
}

func(clr *colour) Intensity() *colour {
	clr.color = clr.color + intensity
	return clr
}

func(clr *colour) GetText() string {
	return clr.color + clr.text + standart
}