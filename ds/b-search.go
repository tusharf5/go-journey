package main

import "fmt"

var array = []int{1, 2, 3, 4, 25, 29, 31, 33, 44, 44, 44, 44, 44, 49, 48, 47, 49, 50, 51, 100, 101, 102, 106, 106, 109, 120, 555, 1000, 1000, 1000, 99999, 99999, 99999, 99999, 99999, 99999, 99999, 99999}

func main() {

	find := 99999
	index := b_search(array, find, 0, len(array)-1)
	fmt.Println("want to find ", find)
	fmt.Println("index is ", index)
	fmt.Println("value is ", array[index])
	fmt.Println("repeating ", count_back_and_forth(array, index))

}

func count_back_and_forth(arr []int, start int) int {

	count := 1

	if (start + 1) < (len(arr) - 1) {
		// forward
		for i := start + 1; arr[i] == arr[start]; i++ {
			count++
			if i+1 > (len(arr) - 1) {
				break
			}
		}
	}

	if (start - 1) > 0 {
		// forward
		for j := start - 1; arr[j] == arr[start]; j-- {
			count++
			if j-1 < 0 {
				break
			}
		}
	}

	return count
}

func b_search(arr []int, find int, start int, finish int) int {

	length := (finish - start) + 1

	if start > finish {
		return -1
	}

	if start == finish {
		if find == arr[start] {
			return start
		} else {
			return -1
		}
	}

	var middle int

	if length%2 == 0 {
		middle = start + (length / 2)
	} else {
		middle = start + ((length + 1) / 2)
	}

	// fmt.Println("middle is ", middle)

	if find == arr[middle] {
		return middle
	} else if find < arr[middle] {
		return b_search(arr, find, start, middle-1)
	} else {
		return b_search(arr, find, middle+1, finish)
	}

}
