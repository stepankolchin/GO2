package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/prac13-rabbit/internal/amqp"
	"example.com/prac13-rabbit/internal/task"
)

type Handler struct {
	repo      *task.Repo
	publisher *amqp.Publisher
}

func NewHandler(repo *task.Repo, publisher *amqp.Publisher) *Handler {
	return &Handler{repo: repo, publisher: publisher}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	t.ID = fmt.Sprintf("t_%d", time.Now().UnixNano())
	_, err := h.repo.Create(t)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	requestID := r.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = "unknown"
	}

	if err := h.publisher.PublishTaskCreated(t.ID, requestID); err != nil {
		http.Error(w, "task created but event publish failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
