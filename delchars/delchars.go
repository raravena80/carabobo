package main

import (
	"fmt"
	"os"
)

var (
	tail int
	n    int
	arr  []string
)

func main() {
	arr = []string{"a", "a", "a", "c", "d", "c", "d", "e", "h", "j"}
	n = len(arr)
	if arr == nil {
		os.Exit(0)
	}
	if n < 2 {
		fmt.Println(arr)
		os.Exit(0)
	}
	tail = 1
	for i := 1; i < n; i++ {
		for j := 0; j <= tail; j++ {
			if arr[i] == arr[j] {
				break
			}
			if j == tail {
				arr[tail] = arr[i]
				tail++
				break
			}

		}

	}
	newarr := arr[0:tail]
	fmt.Println(newarr)

}
