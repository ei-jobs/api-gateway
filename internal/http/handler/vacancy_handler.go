package handler

import (
	"net/http"
	"strconv"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type VacancyHandler struct {
	service *service.VacancyService
}

func NewVacancyHandler(service *service.VacancyService) *VacancyHandler {
	return &VacancyHandler{service: service}
}

func (h *VacancyHandler) GetAllVacancies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filters := model.VacancyFilters{
		SpecializationID: 0,
		Title:            query.Get("title"),
		City:             query.Get("city"),
		Country:          query.Get("country"),
	}

	if specID := query.Get("specialization_id"); specID != "" {
		if id, err := strconv.Atoi(specID); err == nil {
			filters.SpecializationID = id
		}
	}

	if salaryStr := query.Get("salary"); salaryStr != "" {
		if salary, err := strconv.Atoi(salaryStr); err == nil {
			filters.Salary = &salary
		}
	}

	vacancies, err := h.service.GetVacancies(r.Context(), filters)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vacancies)
}

func (h *VacancyHandler) GetVacancy(w http.ResponseWriter, r *http.Request) {
	vacancyID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(vacancyID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	vacancy, err := h.service.GetVacancyByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vacancy)
}

func (h *VacancyHandler) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	vacancy := &model.VacancyRequest{}

	err := utils.ParseJSON(r, vacancy)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.service.CreateVacancy(r.Context(), vacancy)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
	return
}

func (h *VacancyHandler) UpdateVacancy(w http.ResponseWriter, r *http.Request) {
    vacancyIdStr := chi.URLParam(r, "id") 
    vacancyId, err := strconv.Atoi(vacancyIdStr)
    if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    vacancy := &model.VacancyRequest{}
    
    err = utils.ParseJSON(r, vacancy)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

    response, err := h.service.UpdateVacancy(r.Context(), vacancy, vacancyId)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
		return
    }

    utils.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *VacancyHandler) DeleteVacancy(w http.ResponseWriter, r *http.Request) {
    vacancyIdStr := chi.URLParam(r, "id") 
    vacancyId, err := strconv.Atoi(vacancyIdStr)
    if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
    
    err = h.service.DeleteVacancy(r.Context(), vacancyId)
    if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

	utils.WriteJSON(w, http.StatusOK, "Вакансия успешно удалена")
    return
}
