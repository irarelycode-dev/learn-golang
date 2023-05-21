package main

import "fmt"

func variadicFuncs(nums ...int) (int, int) {
	sum := 0
	indices := 0
	for _, num := range nums {
		sum += num
		indices++
	}

	return sum, indices
}

func printVariadic() {
	var nums []int = []int{1, 2, 3, 4, 5}
	sum, indices := variadicFuncs(nums...)
	fmt.Println("sum and indices:", sum, indices)

}
