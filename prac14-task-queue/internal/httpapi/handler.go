package httpapi

import (
	"encoding/json"
	"net/http"

	"example.com/prac14-task-queue/internal/amqp"
	"example.com/prac14-task-queue/internal/jobs"
	"github.com/google/uuid"
)

type Handler struct {
	publisher *amqp.Publisher
}

func NewHandler(publisher *amqp.Publisher) *Handler {
	return &Handler{publisher: publisher}
}

type JobRequest struct {
	TaskID string `json:"task_id"`
}

func (h *Handler) SubmitJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if req.TaskID == "" {
		http.Error(w, "task_id is required", http.StatusBadRequest)
		return
	}

	job := jobs.TaskJob{
		Job:       "process_task",
		TaskID:    req.TaskID,
		Attempt:   1,
		MessageID: uuid.New().String(),
	}

	if err := h.publisher.PublishJob(amqp.MainQueue, job); err != nil {
		http.Error(w, "failed to publish job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "accepted",
		"task_id": req.TaskID,
	})
}
