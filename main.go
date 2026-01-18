package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type Task struct {
	id          uint
	name        string
	description string
}

var tasks []Task

func main() {
	http.HandleFunc("/", helloUser)
	http.ListenAndServe(":8080", nil)

	fmt.Println("What should i call you? (Enter Username):")
	var userName string
	fmt.Scan(&userName)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter Task details: ")
		fmt.Println("Enter Task id: ")
		var task_id uint
		fmt.Scan(&task_id)
		fmt.Println("Enter Task name: ")
		scanner.Scan()
		task_name := scanner.Text()
		fmt.Println("Enter Task description: ")
		scanner.Scan()
		task_desc := scanner.Text()
		task := Task{
			id:          task_id,
			name:        task_name,
			description: task_desc,
		}
		tasks = append(tasks, task)
		//fmt.Println("---------------Task details:---------------")
		//fmt.Printf("Id: %v\n", task_id)
		//fmt.Printf("Name: %v\n", task_name)
		//fmt.Printf("Description: %v\n", task_desc)
		//fmt.Println("-------------------------------------------")
		//fmt.Printf("Tasks: %v\n", tasks)
		fmt.Println("Enter 0 if you have added all the tasks...")
		var exit uint
		fmt.Scan(&exit)
		if exit == 0 {
			break
		}
	}
	fmt.Println("Here are your tasks")
	for _, task := range tasks {
		fmt.Printf("%d. Task: %s Description: %s\n", task.id, task.name, task.description)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Welcome to Go TaskPlanner!"
	fmt.Fprintf(writer, greeting)
}
