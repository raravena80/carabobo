package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		left []int
		right []int
	)
	// initializing lists
	topo := []int{3, 2, 2, 5, 2, 1, 7}
	n := len(topo)

	left = make([]int, n)
	right = make([]int, n)
	left[0] = topo[0]
        for t := 1; t < n; t++ {
		left[t] = int(math.Max(float64(left[t-1]), float64(topo[t])))
	}

	right[n - 1] = topo[n - 1]
	for t := n - 2; t >= 0; t-- {
		right[t] = int(math.Max(float64(right[t + 1]), float64(topo[t])))
	}
	addition := 0
	for t := 0; t < n; t++ {
		addition = addition + (int(math.Min(float64(left[t]), float64(right[t]))) - topo[t])
	}
	fmt.Println(addition)
}
