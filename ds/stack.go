package main

import "fmt"

type Stack struct {
	array []string
}

func (stack *Stack) push(el string) {
	stack.array = append(stack.array, el)
}

func (stack *Stack) pop() string {
	el := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return el
}

func (stack *Stack) print() {
	for _, val := range stack.array {
		fmt.Print(val, ", ")
	}
	fmt.Println("")
}

func main() {
	my_stack := Stack{}

	my_stack.push("hello")
	my_stack.push("sir")
	my_stack.push("how")
	my_stack.print()
	fmt.Println(my_stack.pop())
	my_stack.print()
	my_stack.push("are")
	my_stack.print()

}
