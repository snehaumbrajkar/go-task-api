package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task
var id = 1

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = id
	id++
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service running")
}

func main() {

	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/add", addTask)
	http.HandleFunc("/health", health)

	fmt.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)
}
