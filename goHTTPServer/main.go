package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{
		ID:        "1",
		Title:     "Learn Go",
		Completed: false,
	},
	{
		ID:        "2",
		Title:     "Learn MongoDB",
		Completed: false,
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!! from Go Server")
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	app := chi.NewRouter()
	const PORT int = 3000

	app.Get("/", homeHandler)
	app.Get("/tasks", getTasksHandler)

	address := ":" + strconv.Itoa(PORT)

	fmt.Printf("Server is running on PORT: %d", PORT)
	if err := http.ListenAndServe(address, app); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
