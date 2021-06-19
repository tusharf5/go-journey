package main

import "fmt"

type Stack struct {
	array []string
}

func (stack *Stack) push(el string) {
	stack.array = append(stack.array, el)
}

func (stack *Stack) pop() (bool, string) {

	if len(stack.array) == 0 {
		return true, ""
	}

	el := stack.array[len(stack.array)-1]
	stack.array = stack.array[0 : len(stack.array)-1]

	return false, el
}

func (stack *Stack) print() {
	for _, val := range stack.array {
		fmt.Print(val, ", ")
	}
	fmt.Println("")
}

func main() {
	// my_stack := Stack{}

	// my_stack.push("hello")
	// my_stack.push("sir")
	// my_stack.push("how")
	// my_stack.print()
	// my_stack.print()
	// my_stack.push("are")
	// my_stack.print()

	fmt.Println(reverse_string("hello"))
	fmt.Println(reverse_string("sir"))
	fmt.Println(reverse_string("how"))
	fmt.Println(reverse_string("aeroplane"))

}

func reverse_string(str string) string {

	str_stack := Stack{}
	rvr_string := ""

	for _, val := range str {
		str_stack.push(string(val))
	}

	is_empty, val := str_stack.pop()

	for !is_empty {
		rvr_string = rvr_string + val
		is_empty, val = str_stack.pop()
	}

	return rvr_string

}
