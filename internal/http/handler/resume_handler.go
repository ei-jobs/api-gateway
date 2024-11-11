package handler

import (
	"net/http"
	"strconv"

	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type ResumeHandler struct {
	service *service.ResumeService
}

func NewResumeHandler(service *service.ResumeService) *ResumeHandler {
	return &ResumeHandler{service: service}
}

func (h *ResumeHandler) GetResumesByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resumes, err := h.service.GetResumesByUserID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resumes)
}
