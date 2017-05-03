package main

import (
	"fmt"
	"sort"
	"strings"
)

func notContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}

func main() {
	list := []string{"ear", "era", "mug", "gum", "are", "ear", "sit"}
	anagrams := make(map[string][]string)
	for _, i := range list {
		tmp := strings.Split(i, "")
		sort.Strings(tmp)
		k := strings.Join(tmp, "")
		if notContains(anagrams[k], i) {
			anagrams[k] = append(anagrams[k], i)
		}

	}

	for k, v := range anagrams {
		fmt.Println(k, v)
	}
}
