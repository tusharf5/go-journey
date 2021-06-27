package main

import (
	"fmt"
	"learn/algorithm/str"
	"learn/ds"
)

func main() {
	fmt.Println(str.FirstNonRepeat("sadjbaskjbdkasdnkldqjwiejqoijhfohsdnfjksndqpowjqoiewhfowsnqwpdo"))
	fmt.Println(str.FirstNonRepeat("abcdeffghijkklmnnopqwrsjhairo"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddwjjjkkkllsssppooiieewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddjjjxkkkllsssppooiixeewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("abmmmcccnndjjkskkkslllsblldkjdfjjjfjjja"))
	fmt.Println(str.FirstNonRepeat("ll"))

	minheap := ds.GenMinHeap()

	minheap.MinHeapInsert(0)
	minheap.MinHeapInsert(20)
	fmt.Println(minheap)

	minheap.MinHeapInsert(2)
	fmt.Println(minheap)

	minheap.MinHeapInsert(-2)
	fmt.Println(minheap)

	minheap.MinHeapInsert(9)
	minheap.MinHeapInsert(-10)

	fmt.Println(minheap)

}
