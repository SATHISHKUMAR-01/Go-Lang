package main

import (
	"bufio"
	"fmt"
	"os"
	"strings",
	"strconv"
)

type Task struct {
	task_name string
	status    bool
}

func showMenu() {
	fmt.Println("\nMenu : ")
	fmt.Println("Enter 1 to show all tasks")
	fmt.Println("Enter 2 to add a task")
	fmt.Println("Enter 3 to mark task as completed")
	fmt.Println("Enter 4 to save task to file")
}

func addTask(tasks *[]Task) {
	name := getUserInput("\nEnter the task name : ")
	*tasks = append(*tasks, Task{task_name: name})
	fmt.Println("\nTask added!")
}

func showTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("\nNo tasks available!\n")
	}

	for i, task := range tasks {
		is_completed := "Pending"
		if task.status {
			is_completed = "Done"
		}
		fmt.Printf("%d. %s ---> %s\n", i+1, task.task_name, is_completed)
	}
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func markTaskCompleted(tasks *[]Task) {
	showTask(*tasks)

	taskIdStr := getUserInput("\nEnter the task Id : ")

	taskId, err := parseTaskIndex(taskIdStr)

	if err != nil || taskId < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}

}

func parseTaskIndex(input string) (int, error) {
	return strconv.Atoi(input)
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("\nEnter your choice : ")

		switch option {
		case "1":
			fmt.Println("\n----> Your task list <----\n")
			showTask(tasks)
		case "2":
			fmt.Println("\n----> Add new tasks <----\n")
			addTask(&tasks)
		case "3":
			fmt.Println("\n----> Mark tasks as completed <----\n")
			markTaskCompleted(&tasks)
		case "4":
			fmt.Println("\n----> Save Tasks to file <----\n")
			// saveTasksToFile(tasks)
		}
	}
}
