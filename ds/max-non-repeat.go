package main

import "fmt"

func main() {
	str := "aabbcccccbbddddeeeefnnghmpqswkkkkwookmkskmkqkqwwmkks"
	fmt.Println(max_non_repeat(str))
}

func max_non_repeat(word string) string {

	best_subarray := ""

	current_subarray := ""

	current_subarray_key_map := make(map[string]bool)

	// curr should contain what's at tht end

	for i := 0; i < len(word); i++ {

		if i == 0 {
			current_subarray = string(word[i])
		} else {

			if _, exist := current_subarray_key_map[string(word[i])]; exist {
				current_subarray = string(word[i])
				current_subarray_key_map = make(map[string]bool)
			} else {
				current_subarray = current_subarray + string(word[i])
			}

		}

		current_subarray_key_map[string(word[i])] = true

		if len(current_subarray) > len(best_subarray) {
			best_subarray = current_subarray
		}

	}

	return best_subarray
}
