package main

import (
	"fmt"
)

func getSumArray(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}
	return sum
}

func main() {
	var (
		arr1 []int
		arr2 []int
		sum1 int
		sum2 int
	)
	arr := []int{3, 4, 1, 2, 3, 5, 2}
	n := len(arr) - 1
	for i, e := range arr {
		if i+1 > n {
			arr1 = append(arr1, e)
			arr2 = []int{}
			sum1 = getSumArray(arr1)
			sum2 = 0
		} else {
			arr1 = append(arr1, e)
			arr2 = arr[i+1 : n+1]
			sum1 = getSumArray(arr1)
			sum2 = getSumArray(arr2)
		}
		if sum1 == sum2 {
			fmt.Printf("Array1: %v, Sum1: %v\n", arr1, sum1)
			fmt.Printf("Array2: %v, Sum2: %v\n", arr2, sum2)
		}
	}
}
