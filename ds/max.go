package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 2, -3, 4, -25, -29, 31, 33, 44, -44, -44, -44, -44, -49, 48, 47, -49, 50, -51, 100, -101, 102, 106, -106, 109}
	result, pair := max_sum(arr)
	fmt.Println(result, pair, arr[pair[0]:pair[1]+1])
}

func max_sum(array []int) (float64, []int) {

	pair := []int{-1, -1}

	global_sum := math.Inf(-1)

	local_sum := 0

	is_in_a_conitguous_path := false

	for i := 0; i < len(array); i++ {
		local_sum = int(math.Max(float64(local_sum+array[i]), float64(array[i])))
		if local_sum > int(global_sum) {
			fmt.Println(array[i], local_sum)
			if !is_in_a_conitguous_path {
				pair[0] = i
			}
			global_sum = float64(local_sum)
			pair[1] = i
			is_in_a_conitguous_path = true
		} else {
			is_in_a_conitguous_path = false
		}

	}

	return global_sum, pair
}
