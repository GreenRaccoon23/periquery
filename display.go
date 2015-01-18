package main

import (
	"github.com/fatih/color"
	"strings"
)

var (
	Red      = color.New(color.FgRed)
	Blue     = color.New(color.FgBlue)
	Green    = color.New(color.FgGreen)
	Magenta  = color.New(color.FgMagenta)
	White    = color.New(color.FgWhite)
	Black    = color.New(color.FgBlack)
	BRed     = color.New(color.FgRed, color.Bold)
	BBlue    = color.New(color.FgBlue, color.Bold)
	BGreen   = color.New(color.FgGreen, color.Bold)
	BMagenta = color.New(color.FgMagenta, color.Bold)
	BWhite   = color.New(color.Bold, color.FgWhite)
	BBlack   = color.New(color.Bold, color.FgBlack)
)

func Break(c *color.Color, s string) {
	c.Println(strings.Repeat(s, 79))
}

func Line(c *color.Color) {
	Break(c, "-")
}

func BLine(c *color.Color) {
	Break(c, "=")
}

func ColEnd() {
	color.Unset()
}

func DisplayPericopes(passages [][]string) {
	BGreen.Printf("Pericopes: ")
	i := len(passages) - 1
	for p := 0; p < len(passages); p++ {
		for s := 0; s < len(passages[p]); s++ {
			if p < i {
				Blue.Printf("%s, ", passages[p][s])
			} else {
				Blue.Printf("%s\n", passages[p][s])
			}
		}
	}
}
