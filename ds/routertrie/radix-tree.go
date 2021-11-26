package routertrie

import (
	"fmt"
	"regexp"
	"strings"
)

type RTNode struct {
	IsWord   bool               `json:isLeaf`
	Value    []*Part            `json:value`
	Children map[string]*RTNode `json:children`
}

type Match struct {
	Path   string            `json:path`
	Url    string            `json:url`
	Params map[string]string `json:params`
}

type RadixTree struct {
	Children map[string]*RTNode `json:children`
}

func MakeRadixTree() *RadixTree {
	return &RadixTree{}
}

type Part struct {
	name     string
	value    string
	isStatic bool
}

func makeStringToParts(s string) []*Part {
	arr := strings.Split(s, "/")
	parts := make([]*Part, 0, len(arr))
	check, _ := regexp.Compile("^{(.+?)}$")
	for _, v := range arr {
		if v == "" || v == "/" {
			continue
		}
		if m := check.MatchString(v); m {
			parts = append(parts, &Part{name: v[1 : len(v)-1], value: "*", isStatic: false})
			continue
		}
		parts = append(parts, &Part{name: v, value: v, isStatic: true})
	}
	return parts
}

func (t *RadixTree) Insert(s string) {

	fmt.Println("Path", s, len(s))

	parts := makeStringToParts(s)

	fmt.Println("Parts", len(parts))

	// if no root
	if t.Children == nil {
		t.Children = map[string]*RTNode{}
	}
	if t.Children[(parts[0]).value] == nil {
		node := &RTNode{
			IsWord: true,
			Value:  parts,
		}
		t.Children[(parts[0]).value] = node
	} else {
		t.Children[(parts[0]).value].Insert(parts)
	}
}

func (n *RTNode) split(pivot int) {

	s1, s2 := n.Value[:pivot], n.Value[pivot:]
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

	n.Children[string(s2[0].value)] = newChild

	n.Value = s1
	n.IsWord = false

}

func (n *RTNode) Insert(parts []*Part) {

	// time.Sleep(1 * time.Second)
	// the new word atleast shares same prefix
	if n.Children == nil {
		n.Children = map[string]*RTNode{}
	}

	prefixLen, prefix, _ := commonMatch(&n.Value, &parts)
	splitNode := len(n.Value) > prefixLen
	// time.Sleep(1 * time.Second)

	if splitNode {
		fmt.Println("Splitting")
		// time.Sleep(1 * time.Second)
		n.split(prefixLen)
	}

	// case - there is still some characters left in the new word after the common prefix
	if len(parts[prefixLen:]) > 0 {
		// if there is a child for that, send it to em
		if _, ok := n.Children[string(parts[prefixLen].value)]; ok {
			fmt.Println("Pushing it to next node")
			// time.Sleep(1 * time.Second)
			n.Children[string(parts[prefixLen].value)].Insert(parts[prefixLen:])
			return
		} else {

			// time.Sleep(1 * time.Second)
			n.Children[string(parts[prefixLen].value)] = &RTNode{
				IsWord: true,
				Value:  parts[prefixLen:],
			}
			return
		}

	}

	// there are no characters left after the common prefix
	if ok := isEqual(&prefix, &parts); ok {
		fmt.Println("Exact match")
		// time.Sleep(1 * time.Second)
		n.IsWord = true
		return
	}

	fmt.Println("NO CASE MATCHED")
	// time.Sleep(1 * time.Second)

}

func isEqual(s1, s2 *[]*Part) bool {

	if len(*s1) == len(*s2) {
		for k := range *s1 {
			if (*(*s1)[k]).value != (*(*s2)[k]).value {
				return false
			}
		}
		return true
	}

	return false
}

func commonMatch(trieNode, incoming *[]*Part) (clen int, parts []*Part, params map[string]string) {
	params = map[string]string{}
	prefix, max := -1, -1
	if len(*trieNode) > len(*incoming) {
		max = len(*incoming)
	} else if len(*incoming) > len(*trieNode) {
		max = len(*trieNode)
	} else {
		max = len(*trieNode)
	}

	for i := 0; i < max; i++ {
		if (*(*trieNode)[i]).value == (*(*incoming)[i]).value {
			prefix = i
		} else if (*(*trieNode)[i]).value == "*" {
			prefix = i
			params[(*(*trieNode)[i]).name] = (*(*incoming)[i]).value
		} else {
			break
		}
	}

	if prefix >= 0 {
		return (prefix + 1), (*trieNode)[:prefix+1], params
	} else {
		return -1, nil, params
	}
}

func (t *RadixTree) Delete() {

}

func (t *RadixTree) MatchPath(s string) (found bool, match Match) {

	parts := makeStringToParts(s)

	path := []*RTNode{}

	_, ok := t.Children[string(parts[0].value)]

	if !ok {
		fmt.Println("Eearly exit")
		return false, match
	}

	curr := t.Children[string(parts[0].value)]

	if curr == nil {
		curr = t.Children["*"]
		if curr == nil {
			return false, match
		}
	}

	val := parts
	match.Params = map[string]string{}

	for {
		matchlen, _, p := commonMatch(&curr.Value, &val)

		if matchlen <= 0 {
			return false, match
		}

		if matchlen == len(val) {
			path = append(path, curr)
			for k, v := range p {
				match.Params[k] = v
			}
			break
		}

		for k, v := range p {
			match.Params[k] = v
		}

		path = append(path, curr)

		_, ok = curr.Children[string(val[matchlen].value)]

		if !ok {
			curr, ok = curr.Children["*"]
			if !ok {
				return false, match
			}
		} else {
			curr = curr.Children[string(val[matchlen].value)]
		}
		val = val[matchlen:]
	}

	match.Url = s

	for _, v := range path {
		nodeVal := ""
		for _, v := range v.Value {
			if v.isStatic {
				nodeVal = nodeVal + v.value + "/"
			} else {
				nodeVal = nodeVal + "{" + v.name + "}" + "/"
			}
		}
		match.Path = match.Path + nodeVal
	}

	return true, match
}

func (t *RadixTree) Traverse() {
	if t.Children != nil {
		for k := range t.Children {

			t.Children[k].Traverse(1)
		}
	}
}

func (rtn *RTNode) Traverse(pad int) {
	fmt.Print(strings.Repeat("|-", pad))
	for _, v := range rtn.Value {
		fmt.Print(v.value)
	}
	if rtn.IsWord {
		fmt.Print("<<")
	}
	fmt.Println()
	if rtn.Children != nil {
		for k := range rtn.Children {

			rtn.Children[k].Traverse(pad + 1)
		}
	}
}
