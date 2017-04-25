package main

import (
	"fmt"
)

func main() {
	var (
		output []int
		x      int
	)
	// initializing lists
	list1 := []int{3, 6, 7, 8, 9}
	list2 := []int{6, 8, 10, 30, 40}
	for len(list1) > 0 && len(list2) > 0 {
		if list1[0] < list2[0] {
			x, list1 = list1[0], list1[1:]
			output = append(output, x)
		} else {
			x, list2 = list2[0], list2[1:]
			output = append(output, x)
		}
	}
	output = append(output, list1...)
	output = append(output, list2...)
	fmt.Println(output)
}
