package main

import (
	"fmt"
	"learn/algorithm/str"
	"learn/ds/bst"
	"learn/ds/heap"
	"learn/ds/hmap"
	"learn/ds/routertrie"
	"learn/ds/trie"
)

type MMM struct {
	a string
	b int
	c map[string]int
	d interface{}
	e func()
	f bool
	g rune
	h byte
	i []byte
	j []string
	k []bool
}

func main() {

	rt := routertrie.MakeRadixTree()
	rt.Insert("/")
	rt.Insert("/{module}/")
	rt.Insert("/users/{userId}/")
	rt.Insert("/users/{userId}/posts")
	rt.Insert("/posts/")
	rt.Insert("/posts/{postId}")
	rt.Insert("/posts/{postId}/comments")
	rt.Insert("/posts/{postId}/comments/{commentId}/users/{userId}")
	rt.Insert("/posts/{postId}/comments/{commentId}/users/{userId}/likes")
	rt.Insert("/posts/{postId}/comments/{commentId}/users/likes")
	rt.Insert("/posts/{postId}/comments/{commentId}")

	// rt.Traverse()
	fmt.Println(rt.MatchPath("/users/"))
	fmt.Println(rt.MatchPath("/users"))
	fmt.Println(rt.MatchPath("/users/12"))
	fmt.Println(rt.MatchPath("/posts/gg/comments"))
	fmt.Println(rt.MatchPath("/posts/33/comments/"))
	fmt.Println(rt.MatchPath("/posts/33/comment/"))
	fmt.Println(rt.MatchPath("/users/12/delete/"))
	fmt.Println(rt.MatchPath("/users/12/delet/"))
	fmt.Println(rt.MatchPath("/posts/88/comments/99/users/likes"))
	fmt.Println(rt.MatchPath("/"))
	fmt.Println(rt.MatchPath("/mmmmmmmmm"))
	fmt.Println(rt.MatchPath("/mmmmmmmmm/ffff"))
}

func test() {
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

	bstree := bst.New()
	bstree.Add(1)
	bstree.Add(8)
	bstree.Add(10)
	bstree.Add(2)
	bstree.Add(5)
	bstree.Add(9)

	fmt.Println(bstree)

	ss := new(MMM)
	fmt.Println(ss.a)
	fmt.Println(ss.b)
	fmt.Println(ss.c == nil)
	fmt.Println(ss.d == nil)
	fmt.Println(ss.e == nil)
	fmt.Println(ss.f)
	fmt.Println(ss.g)
	fmt.Println(ss.h)
	fmt.Println(ss.i == nil)
	fmt.Println(ss.j == nil)
	fmt.Println(ss.k == nil)
	fmt.Printf("%+v\n", ss)

	kk := MMM{}
	fmt.Printf("%+v\n", kk)

}
