package main

import (
	"fmt"
	"learn/algorithm/str"
	"learn/ds/heap"
	"learn/ds/hmap"
	"learn/ds/trie"
)

func main() {
	fmt.Println(str.FirstNonRepeat("sadjbaskjbdkasdnkldqjwiejqoijhfohsdnfjksndqpowjqoiewhfowsnqwpdo"))
	fmt.Println(str.FirstNonRepeat("abcdeffghijkklmnnopqwrsjhairo"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddwjjjkkkllsssppooiieewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("nnnnccccbbbddddjjjxkkkllsssppooiixeewuuuffhhskkq"))
	fmt.Println(str.FirstNonRepeat("abmmmcccnndjjkskkkslllsblldkjdfjjjfjjja"))
	fmt.Println(str.FirstNonRepeat("ll"))

	minheap := heap.GenMinHeap()

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

	trie := trie.NewTrie()

	trie.AddWord("Aeroplane")
	trie.AddWord("Alphabet")
	trie.AddWord("Tushar")
	trie.AddWord("Garima")
	trie.AddWord("Garry")
	trie.AddWord("Problem")
	trie.AddWord("Problemo")

	trie.Print()

	hasmap := hmap.New()

	hasmap.Set("kely", "value")
	hasmap.Set("1", 2)
	hasmap.Set("asd3", "my value")

	fmt.Println()
	fmt.Println(hasmap.Get("kely"))
	fmt.Println(hasmap.Get("1"))
	fmt.Println(hasmap.Get("asd3"))

}
