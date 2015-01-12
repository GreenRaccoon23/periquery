package main

import (
	"bytes"
	"flag"
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	//"github.com/fatih/color"
	"log"
	//"os"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

const layout string = "2006-01-02"

var (
	buffer      bytes.Buffer
	translation string
	baseUrl     []string
	date        string
	browse      bool
	globalPos   int
)

func init() {
	flag.StringVar(&translation, "t", "ESV", "translation abbreviation")
	flag.StringVar(&date, "d", time.Now().Format(layout), "date to search")
	flag.BoolVar(&browse, "b", false, "launch pericopes in webbrowser")
	flag.Parse()
	preUrl := "https://www.biblegateway.com/passage/?version="
	midUrl := "&search="
	baseUrl = []string{preUrl, translation, midUrl}
}

func str(slice []string) string {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated := buffer.String()
	buffer.Reset()
	return concatenated
}

func slc(args ...string) []string {
	return args
}

func combine(args ...string) string {
	return str(append(args))
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
	BGreen.Printf("Date: ")
	Blue.Println(found)
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

func urlGen(items [][]string) string {
	encoded := urlEncode(items)
	urlSlice := append(baseUrl, encoded)
	for _, s := range urlSlice {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func urlBrowse(passages [][]string) {
	url := urlGen(passages)
	fmt.Println(url)
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		_ = exec.Command("open", url).Start()
	}
}

func main() {
	defer ColEnd()
	BGreen.Printf("Translation: ")
	BWhite.Println(translation)
	BBlack.Printf("==> ")
	White.Printf("Searching lectionary for date closest to ")
	BWhite.Printf("%s...\n", date)
	dateQuery()
	data := PericopeData[globalPos]
	passages := [][]string{
		data.Psalm,
		data.First,
		data.Epistle,
		data.Gospel,
	}
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
	if browse == true {
		urlBrowse(passages)
	}
}
