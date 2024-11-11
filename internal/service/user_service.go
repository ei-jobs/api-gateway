package service

import (
	"fmt"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Login(login *model.UserLogin) (*model.User, error) {
	user, err := s.repository.GetUserByPhone(login.Phone)
	if err != nil {
		return nil, fmt.Errorf("Phone not found")
	}

	if utils.CheckPasswordHash(user.Password, login.Password) {
		return user, nil
	}

	return nil, fmt.Errorf("Invalid credentials")
}