package main

import (
	"fmt"
	"strings"
)

var (
	str, str1, str2 string
)

func main() {

	str1 = "waterbottle"
	str2 = "bottlewater"

	str = str1 + str1
	if strings.Contains(str, str2) {
		fmt.Printf("%v contains %v\n", str2, str1)
	} else {

		fmt.Printf("%v doesn't contain %v\n", str2, str1)
	}
}
