package main

/*
This is an altertive solution using for like while loops in
other languages
*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		events    []string
		sampleLog []string
		l         string
		match     bool
	)

	file, e := ioutil.ReadFile("./samplelog")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	sampleLog = strings.Split(string(file), "\n")
	for len(sampleLog) > 0 {
		l, sampleLog = sampleLog[0], sampleLog[1:]
		nonEvent, _ := regexp.MatchString("Non event", l)
		if nonEvent {
			continue
		}
		check, _ := regexp.MatchString("Begin ", l)
		if check {
			l, sampleLog = sampleLog[0], sampleLog[1:]
			match, _ = regexp.MatchString("End", l)
			for !match {
				events = append(events, l)
				l, sampleLog = sampleLog[0], sampleLog[1:]
				match, _ = regexp.MatchString("End", l)
				//fmt.Println(l)
			}

		}
	}
	for _, e := range events {
		fmt.Println(e)
	}
}
