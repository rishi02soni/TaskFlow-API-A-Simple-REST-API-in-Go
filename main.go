package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task struct
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// In-memory DB
var tasks []Task
var currentID int = 1

// Create Task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	task.ID = currentID
	currentID++
	tasks = append(tasks, task)

	json.NewEncoder(w).Encode(task)
}

// Get All Tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// Get Single Task
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Update Task
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, task := range tasks {
		if task.ID == id {
			var updated Task
			json.NewDecoder(r.Body).Decode(&updated)

			tasks[i].Title = updated.Title
			tasks[i].Done = updated.Done

			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Delete Task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
