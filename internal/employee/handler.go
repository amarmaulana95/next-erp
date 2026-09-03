package employee

import (
	"encoding/json"
	"net/http"
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
