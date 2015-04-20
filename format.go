package main

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

var (
	buffer bytes.Buffer
)

func Str(slice []string) (concatenated string) {
	for _, s := range slice {
		buffer.WriteString(s)
	}
	concatenated = buffer.String()
	buffer.Reset()
	return
}

func Slc(args ...string) []string {
	return args
}

func Concat(args ...string) string {
	return Str(args)
}

func ConcatBy(slc []string, separator string) (separated string) {
	for _, s := range slc {
		separated = Concat(separated, s, separator)
	}
	separated = strings.TrimSuffix(separated, separator)
	return
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	LogErr(err)
	return i
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	}
	return false
}

func StoD(s string) time.Time {
	d, err := time.Parse(DATE_LAYOUT, s)
	if err != nil {
		Log(err)
		Help()
	}
	return d
}

func DtoS(d time.Time) string {
	return d.Format(DATE_LAYOUT)
}

func IsByteLetter(b uint8, args ...string) bool {
	letter := string(b)
	for _, a := range args {
		if a == letter {
			return true
		}
	}
	return false
}
