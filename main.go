package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// Struktur untuk tugas
type Todo struct {
    ID     int    `json:"id"`
    Task   string `json:"task"`
    Status string `json:"status"`
}

var todos []Todo

func main() {
    http.HandleFunc("/todos", getTodos)
    http.HandleFunc("/add-todo", addTodo)

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler untuk mengambil daftar todos
func getTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

// Handler untuk menambahkan todo baru
func addTodo(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var newTodo Todo
    err := json.NewDecoder(r.Body).Decode(&newTodo)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    newTodo.ID = len(todos) + 1
    todos = append(todos, newTodo)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTodo)
}
