package main

import (
	"fmt"
	"net/http"
)

var tasks = []string{"WakeUp", "Gym", "Breakfast"}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Welcome to Go TaskPlanner!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintf(writer, task)
	}
}
