package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"job-queue-system/services"
)

type JobHandler struct {
	service services.JobService
}

func NewJobHandler(s services.JobService) *JobHandler {
	return &JobHandler{s}
}

func (h *JobHandler) HandleJobSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Payload string `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	job, err := h.service.SubmitJob(input.Payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func (h *JobHandler) HandleJobStatus(w http.ResponseWriter, r *http.Request) {
	// ✅ Correct path trimming
	idStr := strings.TrimPrefix(r.URL.Path, "/job/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.service.GetJob(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
