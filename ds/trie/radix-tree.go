package trie

import (
	"fmt"
	"strconv"
	"strings"
)

type RTNode struct {
	IsWord   bool               `json:isLeaf`
	Value    string             `json:value`
	Children map[string]*RTNode `json:children`
}

type RadixTree struct {
	Children map[string]*RTNode `json:children`
}

func MakeRadixTree() *RadixTree {
	return &RadixTree{}
}

func (t *RadixTree) Insert(s string) {
	// if no root
	if t.Children == nil {
		t.Children = map[string]*RTNode{}
	}
	if t.Children[string(s[0])] == nil {
		node := &RTNode{
			IsWord: true,
			Value:  s,
		}
		t.Children[string(s[0])] = node
	} else {
		t.Children[string(s[0])].Insert(s)
	}
}

func (n *RTNode) split(pivot int) {

	fmt.Println("Splitting " + n.Value)
	s1, s2 := n.Value[:pivot], n.Value[pivot:]
	fmt.Println("Word 1 is " + s1)
	fmt.Println("Word 2 is " + s2)
	newChildren := map[string]*RTNode{}

	for k, v := range n.Children {
		newChildren[k] = v
	}

	n.Children = map[string]*RTNode{}
	newChild := &RTNode{
		IsWord:   true,
		Value:    s2,
		Children: newChildren,
	}

	n.Children[string(s2[0])] = newChild

	n.Value = s1
	n.IsWord = false
	fmt.Println("Current node is " + n.Value)
	fmt.Println("Current node child is " + n.Children[string(s2[0])].Value)

}

func (n *RTNode) Insert(s string) {
	fmt.Println("Adding " + s + " to " + n.Value)
	// time.Sleep(1 * time.Second)
	// the new word atleast shares same prefix
	if n.Children == nil {
		n.Children = map[string]*RTNode{}
	}

	prefixLen, prefix := commonMatch(n.Value, s)
	splitNode := len(n.Value) > prefixLen
	fmt.Println("Prefix is " + prefix + " of length " + strconv.FormatInt(int64(prefixLen), 10))
	// time.Sleep(1 * time.Second)

	if splitNode {
		fmt.Println("Splitting")
		// time.Sleep(1 * time.Second)
		n.split(prefixLen)
	}

	// case - there is still some characters left in the new word after the common prefix
	if len(s[prefixLen:]) > 0 {
		// if there is a child for that, send it to em
		if _, ok := n.Children[string(s[prefixLen])]; ok {
			fmt.Println("Pushing it to next node")
			// time.Sleep(1 * time.Second)
			n.Children[string(s[prefixLen])].Insert(s[prefixLen:])
			return
		} else {
			fmt.Println("Adding " + s[prefixLen:] + " left out as a new child")
			// time.Sleep(1 * time.Second)
			n.Children[string(s[prefixLen])] = &RTNode{
				IsWord: true,
				Value:  s[prefixLen:],
			}
			return
		}

	}

	// there are no characters left after the common prefix
	if prefix == s {
		fmt.Println("Exact match")
		// time.Sleep(1 * time.Second)
		n.IsWord = true
		return
	}

	fmt.Println("NO CASE MATCHED")
	// time.Sleep(1 * time.Second)

}

func commonMatch(s1, s2 string) (int, string) {
	prefix, max := -1, -1
	if len(s1) > len(s2) {
		max = len(s2)
	} else if len(s2) > len(s1) {
		max = len(s1)
	} else {
		max = len(s1)
	}

	for i := 0; i < max; i++ {
		if s1[i] == s2[i] {
			prefix = i
		} else {
			break
		}
	}

	if prefix >= 0 {
		return (prefix + 1), s1[:prefix+1]
	} else {
		return -1, ""
	}
}

func (t *RadixTree) Delete() {

}

func (t *RadixTree) Match(s string) []string {
	path := []*RTNode{}
	_, ok := t.Children[string(s[0])]
	if !ok {
		return []string{}
	}
	curr := t.Children[string(s[0])]
	val := s
	for {
		matchlen, _ := commonMatch(curr.Value, val)
		if matchlen <= 0 {
			break
		}
		if matchlen == len(val) {
			path = append(path, curr)
			break
		}
		// fmt.Println("Curr is "+curr.Value, matchlen, string(val[matchlen]))
		path = append(path, curr)
		curr, ok = curr.Children[string(val[matchlen])]
		if !ok {
			break
		}
		val = val[matchlen:]
	}

	result := []string{}

	for _, v := range path {
		result = append(result, v.Value)
	}

	if strings.Join(result, "") != s {
		return []string{}
	}

	return result
}

func (t *RadixTree) Traverse() {
	if t.Children != nil {
		for k := range t.Children {
			// fmt.Println("At Root" + " Going to " + k)
			t.Children[k].Traverse(1)
		}
	}
}

func (rtn *RTNode) Traverse(pad int) {
	fmt.Print(strings.Repeat("|-", pad) + rtn.Value)
	if rtn.IsWord {
		fmt.Print("**")
	}
	fmt.Println()
	if rtn.Children != nil {
		for k := range rtn.Children {
			// fmt.Println("At " + rtn.Value + " Going to " + k)
			rtn.Children[k].Traverse(pad + 1)
		}
	}
	// fmt.Println("Finished at " + rtn.Value)
}
