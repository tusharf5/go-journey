package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 2, -3, 4, -25, -29, 31, 33, 44, -44, -44}
	arr = append(arr, 3)
	result := max_subarray_sum(arr)
	fmt.Println(result)
}

func max_subarray_sum(array []int) float64 {

	global_best_sum := math.Inf(-1)

	current_sum := 0

	for i := 0; i < len(array); i++ {

		temp_sum := current_sum + int(array[i])

		if temp_sum >= 0 {
			current_sum = temp_sum
		} else {
			current_sum = 0
		}

		if current_sum > int(global_best_sum) {
			global_best_sum = float64(current_sum)
		}
	}

	return global_best_sum
}
