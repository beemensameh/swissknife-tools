package color

import (
	"fmt"
	"strings"
)

type color int

const reset string = "\033[0m"

const (
	Black color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	HiBlack color = iota + 90
	HiRed
	HiGreen
	HiYellow
	HiBlue
	HiMagenta
	HiCyan
	HiWhite
)

func (c color) format() string {
	return fmt.Sprintf("\033[%dm", c)
}

func PrintlnColor(message string, c color) {
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	fmt.Print(SprintfColor(message, c))
}

func SprintfColor(message string, c color) string {
	return fmt.Sprintf("%s%s%s", c.format(), message, reset)
}
