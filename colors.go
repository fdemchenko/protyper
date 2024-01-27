package main

import (
	"errors"

	"github.com/fatih/color"
)

var ColorNames = map[string]color.Attribute{
	"red": color.FgRed, "green": color.FgGreen, "yellow": color.FgYellow,
	"blue": color.FgBlue, "magenta": color.FgMagenta, "cyan": color.FgCyan, "white": color.FgWhite,
}

type ColorArg color.Attribute

func (c *ColorArg) Set(arg string) error {
	*c = ColorArg(color.FgWhite) // default color is white
	if colorAttribute, exists := ColorNames[arg]; exists {
		*c = ColorArg(colorAttribute)
		return nil
	}
	return errors.New("unknown color")
}

func (c ColorArg) String() string {
	for colorName, colorAttribute := range ColorNames {
		if c == ColorArg(colorAttribute) {
			return colorName
		}
	}
	return "unknown"
}
