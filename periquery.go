package main

import (
	"flag"
	//"github.com/PuerkitoBio/goquery"
	//"github.com/fatih/color"
	"bytes"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

const layout string = "2006-01-02"

var (
	buffer      bytes.Buffer
	translation string
	baseUrl     string
	date        string
	browse      bool
	globalPos   int
)

func init() {
	flag.StringVar(&translation, "t", "ESV", "Translation to use")
	flag.StringVar(&date, "d", time.Now().Format(layout), "Date to search in lectionary")
	flag.BoolVar(&browse, "b", false, "Launch found pericopes in a webbrowser")
	flag.Parse()
	preUrl := "https://www.biblegateway.com/passage/?version="
	midUrl := "&search="
	baseUrl = combine(preUrl, translation, midUrl)
}

func main() {
	defer ColEnd()
	Line(Black)
	dateQuery()
	data := PericopeData[globalPos]
	passages := [][]string{
		data.Psalm,
		data.First,
		data.Epistle,
		data.Gospel,
	}
	DisplayPericopes(passages)
	if browse == true {
		urlBrowse(passages)
	}
	Line(Black)
}

func dateLoop(query string) string {
	var found string
	for i := 0; i < len(PericopeData); i++ {
		d := PericopeData[i].Date
		if query == d {
			found = d
			globalPos = i
			break
		}
	}
	return found
}

func dateQuery() {
	defer ColEnd()
	dateDate, err := time.Parse(layout, date)
	if err != nil {
		log.Panic(err)
	}
	var found string
	for d := 0; d < 365; d++ {
		query := dateDate.AddDate(0, 0, d).Format(layout)
		found = dateLoop(query)
		if found != "" {
			break
		}
	}
	date = found
	BGreen.Printf("Lectionary Date: ")
	Blue.Println(found)
	return
}

func urlBrowse(passages [][]string) {
	url := urlGen(passages)
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		_ = exec.Command("open", url).Start()
	}
}

func urlGen(items [][]string) (url string) {
	encoded := urlEncode(items)
	url = combine(baseUrl, encoded)
	return
}

func urlEncode(items [][]string) string {
	var formatted []string
	for p := 0; p < len(items); p++ {
		for f := 0; f < len(items[p]); f++ {
			formatted = append(formatted, items[p][f], ", ")
		}
	}
	combined := str(formatted)
	encoded := url.QueryEscape(combined)
	return encoded
}
