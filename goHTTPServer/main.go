package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Response struct {
	Message string `json:"message"`
	Task    Task   `json:"task"`
}

type TaskResponse struct {
	Message string `json:"message"`
}

type UpdateTask struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

var tasks = []Task{
	{
		ID:        1,
		Title:     "Learn Go",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Learn MongoDB",
		Completed: false,
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!! from Go Server")
}

// GET
func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// POST
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	tasks = append(tasks, task)
	response := Response{
		Message: "success",
		Task:    task,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// PATCH
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var req UpdateTask
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	found := false
	for index := range tasks {
		if tasks[index].ID == taskID {
			found = true
			if req.Title != nil {
				tasks[index].Title = *req.Title
			}
			if req.Completed != nil {
				tasks[index].Completed = *req.Completed
			}
		}
	}

	if !found {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	response := TaskResponse{
		Message: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

// DELETE
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "invalid taskID format", http.StatusBadRequest)
		return
	}

	found := false
	for index := range tasks {
		if tasks[index].ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	response := TaskResponse{
		Message: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// PUT
func putTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	app := chi.NewRouter()
	const PORT int = 3000

	app.Get("/", homeHandler)
	app.Get("/tasks", getTasksHandler)
	app.Post("/tasks", createTaskHandler)
	app.Delete("/tasks/{id}", deleteTaskHandler)

	address := ":" + strconv.Itoa(PORT)

	fmt.Printf("Server is running on PORT: %d", PORT)
	if err := http.ListenAndServe(address, app); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
