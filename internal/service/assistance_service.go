package service

import "github.com/aidosgal/ei-jobs-core/internal/repository"

type AssistanceService struct {
    repository *repository.AssistanceRepository
}

func NewAssistanceService (repository *repository.AssistanceRepository) *AssistanceService {
    return &AssistanceService{repository: repository}
}
