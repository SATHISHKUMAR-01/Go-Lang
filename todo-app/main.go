package main

import (
	"bufio",
	"strings",
	"fmt",
	"os",
	"strconv"
)

type Task struct{
	task_name string
	status bool
}

func showMenu( ) {
	fmt.Println("Enter 1 to show all tasks")
	fmt.Println("Enter 2 to add a task")
	fmt.Println("Enter 3 to mark task as completed")
	fmt.Println("Enter 4 to save task to file")
}

func main(){
	tasks := []Task{}

	for{
		showMenu()
		option := getUserInput("Enter your choice")
	}
}

