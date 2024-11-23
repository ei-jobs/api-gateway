package handler

import (
	"net/http"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
)

type AssistanceHandler struct {
	service *service.AssistanceService
}

func NewAssistanceHandler(service *service.AssistanceService) *AssistanceHandler {
	return &AssistanceHandler{service: service}
}

func (h *AssistanceHandler) GetAssistancesByUserId(w http.ResponseWriter, r *http.Request) {

}

func (h *AssistanceHandler) CreateAssistance(w http.ResponseWriter, r *http.Request) {
	assitance := &model.AssistanceRequest{}
	if err := utils.ParseJSON(r, assitance); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	result, err := h.service.CreateAssistance(r.Context(), assitance)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, result)
	return
}

func (h *AssistanceHandler) UpdateAssistance(w http.ResponseWriter, r *http.Request) {

}

func (h *AssistanceHandler) DeleteAssistance(w http.ResponseWriter, r *http.Request) {

}
