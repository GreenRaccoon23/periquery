package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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
	ColOff   = func() { color.Unset() }
)

func Log(err error) {
	fmt.Println(err)
}

func LogErr(err error) {
	if err == nil {
		return
	}
	Log(err)
}

func ColSet(c *color.Color) {
	if doColor {
		c.Set()
	}
}

func ImgINRI() {
	defer ColOff()
	ColSet(Red)
	image := []string{
		"              .======.              ",
		"              | INRI |              ",
		"              |      |              ",
		"              |      |              ",
		"     .========'      '========.     ",
		"     |   _      xxxx      _   |     ",
		"     |  /_;-.__ / _\\  _.-;_\\  |     ",
		"     |     `-._`'`_/'`.-'     |     ",
		"     '========.`\\   /`========'     ",
		"              | |  / |              ",
		"              |/-.(  |              ",
		"              |\\_._\\ |              ",
		"              | \\ \\`;|              ",
		"              |  > |/|              ",
		"              | / // |              ",
		"              | |//  |              ",
		"              | \\(\\  |              ",
		"              |  ``  |              ",
		"              |      |              ",
		"              |      |              ",
		"              |      |              ",
		"              |      |              ",
		"  \\\\jgs _  _\\\\| \\//  |//_   _ \\// _ ",
		" ^ `^`^ ^`` `^ ^` ``^^`  `^^` `^ `^ ",
	}

	PrintSlc(image)
}

func PrintSlc(sl []string) {
	for _, st := range sl {
		PrintCLn(st)
	}
}

func PrintCLn(st string) {
	sl := strings.Split(st, "\n")
	for _, s := range sl {
		PrintC(s)
	}
}

func PrintC(t string) {
	w := 39 - len(t)/2
	s := strings.Repeat(" ", w)
	fmt.Printf("%v%v%v\n", s, t, s)
}

func Break(s string) {
	fmt.Println(strings.Repeat(s, 79))
}

func Line(c *color.Color) {
	defer ColOff()
	ColSet(c)
	Break("-")
}

func BLine(c *color.Color) {
	defer ColOff()
	ColSet(c)
	Break("=")
}

func Help() {
	defer os.Exit(0)
	w := strings.Repeat(" ", 15)
	i := strings.Repeat(" ", 20)
	fmt.Printf(
		"%v\n  %v%q%v%v%v\n  %v%v%v\n%v%v\n%v%v\n  %v%v%v\n  %v%q%v%v%v\n  %v%v%v\n  %v%v%v\n",
		"periquery <options>",
		"-d ", today, ":  ",
		strings.Repeat(" ", 10-len(today)), "Date to search in the lectionary",
		"-l:",
		w, "Choose from a list of dates in the lectionary, where the",
		i, "list begins with the date closest to today",
		i, "or a date specified with '-d'",
		"-b:",
		w, "Launch found pericopes in a webbrowser",
		"-t ", translation, ":",
		strings.Repeat(" ", 12-len(translation)), "Translation to use",
		"-c:",
		w, "Colorize output",
		"-v:",
		w, "Display extra output",
	)
}

func DListFrom(iStart int) {
	defer ColOff()
	iEnd := len(Lectionary) - 1 //65
	for i := iEnd; i >= iStart; i-- {
		d := Lectionary[i].Date
		n := i - iStart + 1

		if n < 10 {
			fmt.Printf(" ")
		}

		ColSet(BGreen)
		fmt.Printf("  %d:  ", n)
		ColSet(Blue)
		fmt.Println(d)
	}
}

func DChoose(iStart int) (date string) {
	defer ColOff()
	DListFrom(iStart)

	Line(Blue)
	ColSet(BGreen)
	fmt.Printf("Select date:\n    -> ")

	choice := Input(BGreen)
	BLine(BGreen)
	chkChoice(choice)

	i := Int(choice) + iStart - 1
	date = Lectionary[i].Date
	return
}

func chkChoice(s string) {
	if IsInt(s) {
		return
	}
	Help()
}

func DDisplay(date string) {
	defer ColOff()
	ColSet(BGreen)
	fmt.Printf("Lectionary Date: ")
	ColSet(Blue)
	fmt.Println(date)
}

func Input(c *color.Color) (answer string) {
	defer ColOff()
	ColSet(c)
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(response)
	return
}

func (p Pericope) Display() {
	defer ColOff()
	d := p.Date
	r := p.Readings
	s := r.Str()

	ColSet(BGreen)
	fmt.Printf("Pericopes for ")
	ColSet(Blue)
	fmt.Printf("%v", d)
	ColSet(BGreen)
	fmt.Println(":")
	ColSet(BWhite)
	fmt.Println(s)
}

func (r Readings) Display() {
	defer ColOff()
	s := r.Str()

	ColSet(BGreen)
	fmt.Printf("Pericopes: ")
	ColSet(Blue)
	fmt.Println(s)
}
