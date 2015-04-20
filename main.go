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
	flags()
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
