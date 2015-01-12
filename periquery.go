package main

import (
	"bytes"
	"flag"
	"fmt"
	//"golang.org/x/net/html"
	//"io"
	//"io/ioutil"
	//"net/http"
	//"github.com/PuerkitoBio/goquery"
	//"github.com/GreenRaccoon23/periquery"
	"github.com/fatih/color"
	"log"
	//"os"
	//"database/sql"
	//"strconv"
	"net/url"
	"os/exec"
	"reflect"
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
	//PericopeData = PericopeData
	globalPos int
	BGreen    = color.New(color.Bold, color.FgGreen).SprintFunc()
	BBlue     = color.New(color.Bold, color.FgBlue).SprintFunc()
)

func init() {
	//translation = os.Args[1]
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
		//fmt.Printf("%q %T, %q %T\n", d, d, query, query)
		if query == d {
			//fmt.Println(BGreen("Date found: ", d))
			found = d
			globalPos = i
			break
		}
	}
	return found
}

func dateQuery() {
	dateDate, err := time.Parse(layout, date)
	if err != nil {
		log.Panic(err)
	}
	var found string
	for d := 0; d < 365; d++ {
		query := dateDate.AddDate(0, 0, d).Format(layout)
		found := dateLoop(query)
		if found != "" {
			break
		}
	}
	date = found
	fmt.Printf("%s%s\n", BGreen("Date: "), BBlue(date))
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
	//var err error
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		_ = exec.Command("open", "http://localhost:4001/").Start()
	}
}

func main() {
	fmt.Println("type:", reflect.TypeOf(PericopeData[0]))
	fmt.Println(PericopeData[0].Date)
	fmt.Println(translation)
	fmt.Println(date)
	dateQuery()
	data := PericopeData[globalPos]
	//fmt.Println(data)
	passages := [][]string{
		data.Psalm,
		data.First,
		data.Epistle,
		data.Gospel,
	}
	//fmt.Println(len(passages))
	fmt.Printf("%s", BGreen("Pericopes: "))
	i := len(passages) - 1
	for p := 0; p < len(passages); p++ {
		for s := 0; s < len(passages[p]); s++ {
			if p < i {
				fmt.Printf("%s", BBlue(passages[p][s], ", "))
			} else {
				fmt.Printf("%s\n", BBlue(passages[p][s]))
			}
		}
	}
	//fmt.Println(passages)
	//fmt.Println(passages[0])
	if browse == true {
		urlBrowse(passages)
	}
	//for _, a := range bible {
	//i := a.index
	//x, _ := strconv.Atoi(i)
	//fmt.Println(a, x)
	//}
	//translation := "ESV"
	//preUrl := "https://www.biblegateway.com/passage/?version="
	//midUrl := "&search="
	//b := []byte(preUrl)
	//baseUrl := []string{preUrl, translation, midUrl}
	//curr_chap := []string{bible[0].book, bible[0].chapter_range}
	//url := concatUrl(bible[0].book, bible[0].chapter_range)
	//curr_chap := strconv.Itoa(bible[0].chapter_range)
	//a := concatUrl(baseUrl, bible[0].book, curr_chap)
	//bookLoop(bible[0])
	//bookLoop(bible)
	//fmt.Println(a)
	//urlSlice := append(baseUrl, midUrl)
	//buffer.WriteString(translation)
	//url := buffer.String()
	//baseUrl := copy(preUrl, midUrl)
	//resp, err := http.Get("http://example.com/")
	//fmt.Println(baseUrl)
}
