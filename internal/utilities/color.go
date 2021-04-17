package utilities

import (
	"fmt"
)

var (
	colors = map[string]int{
		"white":  0,
		"red":    1,
		"green":  2,
		"orange": 3,
		"blue":   4,
		"pink":   5,
		"cyan":   6,
		"gray":   7,
	}

	colorTypes = map[string]int{
		"bold":       1,
		"normal":     5,
		"underline":  4,
		"background": 7,
	}
)

// Colorize string
func ColorString(text, color, font string) string {
	if font == "" {
		font = "normal"
	}

	if color == "" {
		color = ""
	}

	return fmt.Sprintf(
		"\033[3%d;%dm%s\033[0m",
		colors[color],
		colorTypes[font],
		text,
	)
}
