package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *TodoService
}

func NewHandler(service *TodoService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.ListTodos()
	if err != nil {
		http.Error(w, "Failed to list todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("❌ Failed to decode request: %v", err) // ← log added
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	if todo.Title == "" {
		log.Printf("❌ Missing title in request") // ← log added
		http.Error(w, `{"error": "Title is required"}`, http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateTodo(&todo)
	if err != nil {
		log.Printf("❌ Failed to create todo: %v", err) // ← log added
		http.Error(w, `{"error": "Failed to create todo"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Todo created: %+v", created) // ← log added
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.Printf("❌ Failed to encode response: %v", err) // ← log added
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	todo, err := h.service.GetTodo(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var updated Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo, err := h.service.UpdateTodo(id, &updated)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
