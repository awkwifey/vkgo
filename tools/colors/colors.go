package colors

const (
	standart		= "\033[0m"
	bold			= "\033[1m"
	intensity		= "\033[2m"
	underline		= "\033[4m"
	flashing		= "\033[5m"

	resetIntensity		= "\033[22m"
	cancelUnderline		= "\033[24m"
	cancelSemiBright	= "\033[25m"
	resetFlashing		= "\033[27m"

	blackMark	= "\033[30m"
	redMark		= "\033[31m"
	greenMark	= "\033[32m"
	yellowMark	= "\033[33m"
	blueMark	= "\033[34m"
	magentaMark	= "\033[35m"
	cyanMark	= "\033[36m"
	greyMark	= "\033[37m"

	blackBackground		= "\033[40m"
	redBackground		= "\033[41m"
	greenBackground		= "\033[42m"
	yellowBackground	= "\033[43m"
	blueBackground		= "\033[44m"
	magentaBackground	= "\033[45m"
	cyanBackground		= "\033[46m"
	greyBackground		= "\033[47m"
)

func Black(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + blackMark,
	}
}

func Red(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + redMark,
	}
}

func Green(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + greenMark,
	}
}

func Yellow(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + yellowMark,
	}
}

func Blue(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + blueMark,
	}
}

func Magenta(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + magentaMark,
	}
}

func Cyan(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + cyanMark,
	}
}

func Grey(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + greyMark,
	}
}

func White(text string) *colour {
	return &colour{
		text: text,
		color: "\033[" + standart,
	}
}