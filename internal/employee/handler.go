package employee

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := h.repository.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to get employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid employee id", http.StatusBadRequest)
		return
	}

	employee, err := h.repository.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(employee); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
