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

func makePathToArray(s string) []*Part {
	arr := strings.Split(s, "/")
	parts := make([]*Part, 0, len(arr))
	check, _ := regexp.Compile("^{.+?}$")
	for _, v := range arr {
		if v == "" || v == "/" {
			continue
		}
		if m := check.MatchString(v); m {
			fmt.Println("Adding *")
			parts = append(parts, &Part{name: v, value: "*", isStatic: false})
			continue
		}
		fmt.Println("Adding " + v)
		parts = append(parts, &Part{name: v, value: v, isStatic: true})
	}
	return parts
}

func (t *RadixTree) Insert(s string) {

	fmt.Println("Path", s, len(s))

	parts := makePathToArray(s)

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

	// fmt.Println("Splitting " + n.Value)
	s1, s2 := n.Value[:pivot], n.Value[pivot:]
	// fmt.Println("Word 1 is " + s1)
	// fmt.Println("Word 2 is " + s2)
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
	// fmt.Println("Current node is " + n.Value)
	// fmt.Println("Current node child is " + n.Children[string(s2[0])].Value)

}

func (n *RTNode) Insert(parts []*Part) {
	// fmt.Println("Adding " + s + " to " + n.Value)

	// time.Sleep(1 * time.Second)
	// the new word atleast shares same prefix
	if n.Children == nil {
		n.Children = map[string]*RTNode{}
	}

	prefixLen, prefix := commonMatch(&n.Value, &parts)
	splitNode := len(n.Value) > prefixLen
	// fmt.Println("Prefix is " + prefix + " of length " + strconv.FormatInt(int64(prefixLen), 10))
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
			// fmt.Println("Adding " + s[prefixLen:] + " left out as a new child")
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
		for k, _ := range *s1 {
			if (*(*s1)[k]).value != (*(*s2)[k]).value {
				return false
			}
		}
		return true
	}

	return false
}

func commonMatch(s1, s2 *[]*Part) (int, []*Part) {
	prefix, max := -1, -1
	if len(*s1) > len(*s2) {
		max = len(*s2)
	} else if len(*s2) > len(*s1) {
		max = len(*s1)
	} else {
		max = len(*s1)
	}

	for i := 0; i < max; i++ {
		if (*(*s1)[i]).value == (*(*s2)[i]).value {
			prefix = i
		} else {
			break
		}
	}

	if prefix >= 0 {
		return (prefix + 1), (*s1)[:prefix+1]
	} else {
		return -1, nil
	}
}

func (t *RadixTree) Delete() {

}

func (t *RadixTree) Match(s string) []string {

	parts := makePathToArray(s)

	path := []*RTNode{}
	_, ok := t.Children[string(parts[0].value)]
	if !ok {
		fmt.Println("Eearly exit")
		return []string{}
	}
	curr := t.Children[string(parts[0].value)]
	val := parts
	for {
		matchlen, _ := commonMatch(&curr.Value, &val)
		if matchlen <= 0 {

			break
		}
		if matchlen == len(val) {

			path = append(path, curr)
			break
		}
		// fmt.Println("Curr is "+curr.Value, matchlen, string(val[matchlen]))
		path = append(path, curr)
		curr, ok = curr.Children[string(val[matchlen].value)]
		if !ok {
			break
		}
		val = val[matchlen:]
	}

	result := []string{}

	for _, v := range path {
		nodeVal := ""
		for _, v := range v.Value {
			nodeVal = nodeVal + v.value
		}
		fmt.Println("Node", nodeVal)
		result = append(result, nodeVal)
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
			// fmt.Println("At " + rtn.Value + " Going to " + k)
			rtn.Children[k].Traverse(pad + 1)
		}
	}
	// fmt.Println("Finished at " + rtn.Value)
}
