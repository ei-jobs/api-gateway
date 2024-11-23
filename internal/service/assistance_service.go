package service

import (
	"context"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
)

type AssistanceService struct {
	repository *repository.AssistanceRepository
}

func NewAssistanceService(repository *repository.AssistanceRepository) *AssistanceService {
	return &AssistanceService{repository: repository}
}

func (s *AssistanceService) CreateAssistance(ctx context.Context, assitance *model.AssistanceRequest) (*model.AssistanceRequest, error) {
	return s.repository.StoreAssistance(ctx, assitance)
}
