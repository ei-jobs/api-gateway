package handler

import (
	"net/http"
	"strconv"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
	return
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	login := &model.UserLogin{}
	err := utils.ParseJSON(r, login)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.Login(login)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	register := &model.UserRegisterRequest{}
	if err := utils.ParseJSON(r, register); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	user, err := h.service.Register(register)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
	return
}

func (h *UserHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.service.GetAllCompanies()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, companies)
	return
}
