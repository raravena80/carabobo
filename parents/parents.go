package main

import (
	"fmt"
	"sort"
	"strings"
)

func parents() {
	var (
		parents map[int]int
		one     []int
		zero    []int
	)
	arr := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6},
		{5, 7}, {4, 5}, {4, 8}, {8, 9}}
	parents = make(map[int]int)
	for _, e := range arr {
		if _, ok := parents[e[0]]; !ok {
			parents[e[0]] = 0
		}
		if _, ok := parents[e[1]]; !ok {
			parents[e[1]] = 1
		} else {
			parents[e[1]] = parents[e[1]] + 1
		}
	}
	for i, j := range parents {
		if j == 0 {
			zero = append(zero, i)
		}
		if j == 1 {
			one = append(one, i)
		}
	}
	sort.Ints(zero)
	sort.Ints(one)
	fmt.Println("Zero parents:", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(zero)), ", "), "[]"))
	fmt.Println("One parent:", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(one)), ", "), "[]"))

}

func main() {
	parents()
}
