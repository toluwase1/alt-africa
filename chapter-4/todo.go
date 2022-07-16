package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//welcome page function
//getting inputs from the console
//createTod
//deleteTodoById
//listAllTodos created

var todolist []string

func main() {

	for {
		welcomePage()
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("an error occured, please enter the right input")
		}

		switch input {
		case 1:
			fmt.Println("block 1")
			title, description := getUserInputToCreateTodo()
			createTodo(title, description)
		case 2:
			fmt.Println("block 2")
			id := getIdFromUser()
			deleteTodoById(id)
		case 3:
			fmt.Println("block 3")
			id := getIdFromUser()
			updateATodo(id)
		case 4:
			fmt.Println("block 4")
			listAllItemsInTodo()
		default:
			fmt.Println("please enter the correct input")
		}

	}
}

func welcomePage() {
	fmt.Println("Welcome to our Todo App")
	fmt.Println("---------------------------")
	fmt.Println("Type 1 to create a new todo")
	fmt.Println("Type 2 to delete a  todo")
	fmt.Println("Type 3 to update a todo")
	fmt.Println("Type 4 to list all todos")
}

func createTodo(title string, description string) []string {
	todo := title + " " + description
	todolist = append(todolist, todo)
	fmt.Println(todolist)
	return todolist
}

func getUserInputToCreateTodo() (string, string) {
	var title string
	var description string

	fmt.Println("Please enter the title of your task")
	title = readSpacedSentences()
	fmt.Println("Please enter the description of your task")
	description = readSpacedSentences()

	return title, description

}

func listAllItemsInTodo() {
	for i := 0; i < len(todolist); i++ {
		value := fmt.Sprintf("list of all todos: \n  %v: | %v", i, todolist[i])
		fmt.Println(value)
	}
}

func getIdFromUser() int {
	var id int
	fmt.Println("Please enter the id of your task: ")
	fmt.Scan(&id)

	return id
}
func updateATodo(id int) {
	title, description := getUserInputToCreateTodo()
	todo := title + ": " + description
	for i, _ := range todolist {
		if id == i {
			todolist[i] = todo
		}
	}
	fmt.Println(todolist)
}

func deleteTodoById(id int) error {

	if id < 0 || id > len(todolist) {
		fmt.Println("Please enter the right id of your todo")
		errors.New("Please enter the right id of your todo")
	}

	for i := 0; i < len(todolist); i++ {
		if id == i {
			todolist = append(todolist[:i], todolist[i+1:]...)
		}
	}
	fmt.Println(todolist)
	return nil
}

func readSpacedSentences() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	if scanner.Scan() {
		text = scanner.Text()
	}
	return text
}
