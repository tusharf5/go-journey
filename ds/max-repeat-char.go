package main

import "fmt"

func main() {

	str := "aabbcccccbbddddeeeefnnghmpqswkkkkwookmkskmkqkqwwmkks"
	fmt.Println(max_repeat_char(str))
}

func max_repeat_char(chars string) string {

	// kadane's also
	// if you have a list
	// it might have a sub-list somewhere in between the start and the end that yields the maximum
	// it might have a sub-list that ends at the last element that yields the maximum

	// when you add a new element to it

	// 3 things could happen

	// Case 1 - the new element has no effect on the maximum sub list as adding it does not make a new maximum
	// Case 2 - it creates a new maximum sublist ending at the last element which is this newly added
	// Case 3 - in some cases the new element is the maximum itself

	max_repeat := ""

	local_max_repeat := "-"

	// local will reset for every satisfyiable subarray
	// maximum will reset for only the max out of every satisfyiable subarray

	for i := 0; i < len(chars); i++ {

		if i != 0 && chars[i] == chars[i-1] && chars[i] == local_max_repeat[0] {
			// Case 2 check
			local_max_repeat = local_max_repeat + string(chars[i])
		} else {
			// Case 3 check
			local_max_repeat = string(chars[i])
		}

		if len(local_max_repeat) > len(max_repeat) {
			max_repeat = local_max_repeat
		} // else {
		// Case 1
		//	}

	}

	return max_repeat

}
