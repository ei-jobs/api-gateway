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

func (s *UserService) Register(register *model.UserRegisterRequest) (*model.User, error) {
	var err error
	register.Password, err = utils.HashUserPassword(register.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.CreateUser(register)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllCompanies() ([]*model.Company, error) {
	companies, err := s.repository.GetUsersByRoleId(2)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *UserService) GetUserById(id int) (*model.UserResponse, error) {
	return s.repository.GetUserById(id)
}
