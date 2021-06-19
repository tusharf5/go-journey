package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Status string

var john User = User{
	Name: "Tushar",
	Age:  23,
}

// by defaukt akways start with a pointer

func modufyUser(user *User) {
	// the selector operator autimatiaclly deferences pointers no need to do *
	user.Name = "modified"
	*user = User{"whhh", 3}
}

func main() {
	jj := User{
		Name: "Jimm",
		Age:  24,
	}

	kk := &User{
		Name: "Jimm",
		Age:  24,
	}

	fmt.Println(jj)
	// modufyUser(jj) // won't take it since jj is not a pointer
	modufyUser(kk)
	fmt.Println(kk)
}
