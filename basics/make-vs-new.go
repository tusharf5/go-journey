package main

import (
	"fmt"
	"math/rand"
	"time"
)

// so new is for initialization to zero type
// it also returns a pointer
// good for structs

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-1234567890"

type Todo struct {
	id     string
	text   string
	status string
}

type TodoList struct {
	todos *[]Todo
}

type User struct {
	first_name string
	last_name  string
	todos      *TodoList
}

func main() {

	user := create_user("Tushar", "Sharms")

	user.todos = create_todo_list()

	todos := *(user.todos.todos)
	todos = append(todos, create_todo("Learn Go"))

	user.todos.todos = &todos
	fmt.Println(user)
	fmt.Println(user.todos.todos)

}

func create_user(first_name, last_name string) *User {

	user := new(User)
	user.first_name = first_name
	user.last_name = last_name

	return user
}

func create_todo(text string) Todo {
	id := create_id()
	todo := Todo{}

	todo.id = id
	todo.text = text
	todo.status = "TODO"

	return todo

}

func create_todo_list() *TodoList {

	todo_list := new(TodoList)
	todos := make([]Todo, 10)
	todo_list.todos = &todos
	return todo_list

}

func create_id() string {

	max := len(letters)
	id := ""

	rand_gen := rand.New(rand.NewSource(time.Now().Local().UnixNano()))

	for i := 0; i < 10; i++ {
		index := rand_gen.Intn(max)
		id = id + string(letters[index])
	}

	return id
}
