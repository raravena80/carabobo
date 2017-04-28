package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"regexp"
)

func main() {
	var (
		capture bool = false
		release bool = false
		events []string
	)

	file, e := ioutil.ReadFile("./samplelog")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	sampleLog := strings.Split(string(file), "\n")
	for _, l := range sampleLog {
		match, _ := regexp.MatchString("Non event", l)
		if match {
			continue
		}
		event, _ := regexp.MatchString("event ", l)
		if !capture {
			capture, _ = regexp.MatchString("Begin", l)
		}
		release, _ = regexp.MatchString("End", l)
		if release {
			capture = false
			release = false
		}
		if capture && event {
			events = append(events, l)
		}
	}
	for _, e := range events {
		fmt.Println(e)
	}
}
