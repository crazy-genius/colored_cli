package color

type Color string

const (
	White  Color = "\033[37m"
	Red          = "\033[31m"
	Black        = "\033[30m"
	Blue         = "\033[34m"
	Green        = "\033[32m"
	Cyan         = "\033[36m"
	Purple       = "\033[35m"
	Brown        = "\033[33m"
	Yellow       = "\033[1;33m"

	DarkGray    = "\033[1;30m"
	LightBlue   = "\033[1;34m"
	LightGreen  = "\033[1;32m"
	LightCyan   = "\033[1;36m"
	LightRed    = "\033[1;31m"
	LightPurple = "\033[1;35m"
	LightGray   = "\033[37m"

	NoEntered = ""
)
