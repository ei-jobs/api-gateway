package handler

import (
	"net/http"

	"github.com/aidosgal/ei-jobs-core/internal/service"
)

type AssistanceHandler struct {
    service *service.AssistanceService
}

func NewAssistanceHandler(service *service.AssistanceService) *AssistanceHandler {
    return &AssistanceHandler{service: service}
}

func (h *AssistanceHandler) GetAssistancesByUserId (w http.ResponseWriter, r *http.Request) {
    
}
