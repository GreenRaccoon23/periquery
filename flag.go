package main

import (
	//"flag"
	"os"
)

var (
	translation string = "ESV"
	date        string = today
	doBrowse    bool
	doList      bool
	doColor     bool
	doVerbose   bool
)

func ChkHelp() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "-h", "h", "help", "--help", "-H", "H", "HELP", "--HELP", "-help", "--h", "--H":
		Help()
	}

}

/*func flags() {
	var doListBrowse bool
	flag.StringVar(&translation, "t", "ESV", "")
	flag.StringVar(&date, "d", today, "")
	flag.BoolVar(&doBrowse, "b", false, "")
	flag.BoolVar(&doList, "l", false, "")
	flag.BoolVar(&doColor, "c", false, "")
	flag.BoolVar(&doVerbose, "v", false, "")

	flag.BoolVar(&doListBrowse, "lb", false, "Combination of `-l` and `-b`")
	flag.Parse()

	if doListBrowse {
		doBrowse = true
		doList = true
	}
}*/

func flags() {
	flagsS := map[string]*string{
		"t": &translation,
		"d": &date,
	}
	flagsB := map[string]*bool{
		"b": &doBrowse,
		"l": &doList,
		"c": &doColor,
		"v": &doVerbose,
	}

	for i, f := range os.Args {
		if len(f) == 0 {
			continue
		}
		if IsByteLetter(f[0], "-") == false {
			continue
		}

		for _, r := range f[1:] {
			s := string(r)
			BoolParse(flagsB, s)
			StrParse(flagsS, i, s)
		}
	}
}

func BoolParse(m map[string]*bool, f string) {
	for s, b := range m {
		if s != f {
			continue
		}
		*b = true
	}
}

func StrParse(m map[string]*string, i int, f string) {
	for s, t := range m {
		if s != f {
			continue
		}
		*t = NextArg(i)
	}
}

func NextArg(i int) string {
	if len(os.Args) <= i {
		Help()
	}
	return os.Args[i+1]
}

/*
func isFlag(q string) bool {
	for i, f := range os.Args {
		if len(f) == 0 {
			continue
		}
		if IsByteLetter(f[0], "-") == false {
			continue
		}

		if qFlag(f, q) {
			return true
		}
	}
	return false
}

func qFlag(a, q string) bool {
	for n, r := range a[1:] {

	}
}
*/
