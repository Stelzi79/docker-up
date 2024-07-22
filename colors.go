package main

// col is a map of col
type col struct {
	Black        string
	Red          string
	Green        string
	Yellow       string
	Blue         string
	Magenta      string
	Cyan         string
	White        string
	LightGray    string
	DarkGray     string
	DarkBlue     string
	DarkGreen    string
	DarkCyan     string
	DarkRed      string
	DarkMagenta  string
	LightRed     string
	LightGreen   string
	LightYellow  string
	LightBlue    string
	LightMagenta string
	LightCyan    string
	LightWhite   string
	LightBlack   string
}

type formt struct {
	Reset         string
	Bold          string
	Underline     string
	Blink         string
	Reverse       string
	Concealed     string
	ClearScreen   string
	ClearLine     string
	SaveCursor    string
	RestoreCursor string
	EraseLine     string
	EraseDown     string
	EraseUp       string
	EraseScreen   string
	EraseStart    string
	EraseEnd      string
	Tab           string
}

// Colors represents a collection of ANSI escape codes for different colors.
var Colors = col{
	Black:        "\033[30m",
	Red:          "\033[31m",
	Green:        "\033[0;32m",
	Yellow:       "\033[33m",
	Blue:         "\033[34m",
	Magenta:      "\033[35m",
	Cyan:         "\033[36m",
	White:        "\033[37m",
	LightGray:    "\033[0;37m",
	DarkGray:     "\033[1;30m",
	DarkBlue:     "\033[0;34m",
	DarkGreen:    "\033[0;32m",
	DarkCyan:     "\033[0;36m",
	DarkRed:      "\033[0;31m",
	DarkMagenta:  "\033[0;35m",
	LightRed:     "\033[1;31m",
	LightGreen:   "\033[1;32m",
	LightYellow:  "\033[1;33m",
	LightBlue:    "\033[1;34m",
	LightMagenta: "\033[1;35m",
	LightCyan:    "\033[1;36m",
	LightWhite:   "\033[1;37m",
	LightBlack:   "\033[0;30m",
}

// Format represents a collection of ANSI escape codes for different text formats.
var Format = formt{
	Reset:         "\033[0m",
	Bold:          "\033[1m",
	Underline:     "\033[4m",
	Blink:         "\033[5m",
	Reverse:       "\033[7m",
	Concealed:     "\033[8m",
	ClearScreen:   "\033[2J",
	ClearLine:     "\033[2K",
	SaveCursor:    "\033[s",
	RestoreCursor: "\033[u",
	EraseLine:     "\033[K",
	EraseDown:     "\033[J",
	EraseUp:       "\033[1J",
	EraseScreen:   "\033[2J",
	EraseStart:    "\033[1K",
	EraseEnd:      "\033[K",
	Tab:           "  ",
}
