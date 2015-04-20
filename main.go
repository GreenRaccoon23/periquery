package main

import (
	"time"
)

const (
	DATE_LAYOUT = "2006-01-02"
	URL_PRE     = "https://www.biblegateway.com/passage/?version="
	URL_MID     = "&search="
)

var (
	today string = time.Now().Format(DATE_LAYOUT)
)

func init() {
	ChkHelp()
	flags()
	DFmt(&date)
}

func main() {
	if doVerbose {
		ImgINRI()
	}

	d, i := DFind(date)
	if doList {
		q := DChoose(i)
		d, i = DFind(q)
	}
	if doVerbose {
		DDisplay(d)
	}

	pericope := Lectionary[i]
	passages := pericope.Readings
	if doVerbose {
		passages.Display()
		Line(BWhite)
	}
	pericope.Display()

	if doBrowse {
		urlbase := Concat(URL_PRE, translation, URL_MID)
		passages.Browse(urlbase)
	}
}

func DFmt(d *string) {
	n := len(*d)
	switch n {
	case 10:
		if IsByteLetter((*d)[4], "-") && IsByteLetter((*d)[6], "-") {
			return
		}
	case 8:
		if IsByteLetter((*d)[2], "-") && IsByteLetter((*d)[4], "-") {
			*d = Concat("20", *d)
			return
		}
	case 5:
		if IsByteLetter((*d)[2], "-") {
			*d = Concat("2015-", *d)
			return
		}
	case 4:
		if IsByteLetter((*d)[1], "-") {
			*d = Concat("2015-0", *d)
			return
		}
	}
}

func DFind(date string) (string, int) {
	s := StoD(date)
	return DSearch(s)
}

func DSearch(date time.Time) (d string, i int) {
	var found bool
	for day := 0; day < 366; day++ {
		q := DtoS(date)
		d, i, found = DQuery(q)

		if found {
			return
		}
		date = date.AddDate(0, 0, 1)
	}
	return
}

func DQuery(q string) (d string, i int, found bool) {
	for i = 0; i < len(Lectionary); i++ {
		d = Lectionary[i].Date
		if d == q {
			found = true
			return
		}
	}
	found = false
	return
}
