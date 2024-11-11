package handler

import (
	"net/http"
	"strconv"

	"github.com/aidosgal/ei-jobs-core/internal/model"
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

func (h *ResumeHandler) CreateResume(w http.ResponseWriter, r *http.Request) {
	var resume model.Resume
	if err := utils.ParseJSON(r, &resume); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	createdResume, err := h.service.CreateResume(&resume)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdResume)
}

func (h *ResumeHandler) UpdateResume(w http.ResponseWriter, r *http.Request) {
	resumeIDStr := chi.URLParam(r, "resumeID")
	resumeID, err := strconv.Atoi(resumeIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var resume model.Resume
	if err := utils.ParseJSON(r, &resume); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resume.ID = resumeID

	updatedResume, err := h.service.UpdateResume(&resume)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedResume)
}

func (h *ResumeHandler) DeleteResume(w http.ResponseWriter, r *http.Request) {
	resumeIDStr := chi.URLParam(r, "resumeID")
	resumeID, err := strconv.Atoi(resumeIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.DeleteResume(resumeID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Resume deleted successfully"})
}
