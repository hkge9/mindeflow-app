package transport

import (
	"encoding/json"
	"mindeflow-app/backend/internal/inbox"
	"mindeflow-app/backend/internal/inbox/service"
	"mindeflow-app/backend/internal/utils"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateInboxItem(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	item, err := h.service.Create(r.Context(), inbox.CreateInput{
		Title: req.Title,
		Text:  req.Text,
	})
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, item)
}
