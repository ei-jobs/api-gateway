package handler

import (
	"net/http"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var login *model.UserLogin
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
