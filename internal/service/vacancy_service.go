package service

import (
	"context"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
)

type VacancyService struct {
	repository *repository.VacancyRepository
}

func NewVacancyService(repository *repository.VacancyRepository) *VacancyService {
	return &VacancyService{repository: repository}
}

func (s *VacancyService) GetVacancies(ctx context.Context, filters model.VacancyFilters) ([]model.Vacancy, error) {
	return s.repository.GetVacancies(ctx, filters)
}

func (s *VacancyService) GetVacancyByID(ctx context.Context, id int) (*model.OneVacancy, error) {
	return s.repository.GetVacancyByID(ctx, id)
}
