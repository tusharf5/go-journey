package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	fmt.Println(alex)
}
