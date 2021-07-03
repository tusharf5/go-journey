package trie

import "fmt"

type Node struct {
	isWord   bool
	value    string
	children *map[string]*Node
}

type Trie map[string]*Node

func NewTrie() *Trie {
	base := make(Trie)
	return &base
}

func (d *Trie) AddWord(word string) {

	node, ok := (*d)[string(word[0])]
	isLast := len(word) == 1

	if ok {
		if isLast {
			node.isWord = true
		} else {
			node.AddWord(word[1:])
		}
	} else {

		children := make(map[string]*Node)
		newNode := &Node{
			isLast,
			string(word[0]),
			&children,
		}
		(*d)[string(word[0])] = newNode

		if !isLast {
			newNode.AddWord(word[1:])
		}
	}

}

func (d *Node) AddWord(word string) {
	node, ok := (*(d.children))[string(word[0])]
	isLast := len(word) == 1

	if ok {
		if isLast {
			node.isWord = true
		} else {
			node.AddWord(word[1:])
		}
	} else {
		children := make(map[string]*Node)
		newNode := &Node{
			isLast,
			string(word[0]),
			&children,
		}
		(*(d.children))[string(word[0])] = newNode
		if !isLast {
			newNode.AddWord(word[1:])
		}
	}
}

func (d *Trie) Print() {

	for _, val := range *d {

		fmt.Print(string(val.value))

		val.PrintChildren()

	}

}

func (d *Node) PrintChildren() {

	for _, val := range *(d.children) {

		fmt.Print(string(val.value))

		val.PrintChildren()

	}

}
