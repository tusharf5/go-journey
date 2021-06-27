package main

import (
	"fmt"
	"learn/algorithm/str"
	"learn/ds"
)

func main() {
	fmt.Println(str.FirstNonRepeat("sadjbaskjdkasdnkldqjwiejqoijhfohsdnfjksndqpowjqoiewhfowsnqwpdo"))
	fmt.Println(str.FirstNonRepeat("abcdeffghijkklmnnopqwrsjhiro"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddjjjkkkllsssppooiieewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddjjjxkkkllsssppooiieewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("abmmmcccnndjjkskkkslllslldkjdfjjjfjjja"))
	fmt.Println(str.FirstNonRepeat("l"))

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
