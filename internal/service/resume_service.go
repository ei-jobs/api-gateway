package service

import (
	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
)

type ResumeService struct {
	repository *repository.ResumeRepository
}

func NewResumeService(repository *repository.ResumeRepository) *ResumeService {
	return &ResumeService{repository: repository}
}

// GetResumesByUserID fetches all resumes for a specific user by their user ID.
func (s *ResumeService) GetResumesByUserID(userID int) ([]*model.Resume, error) {
	// Fetch resumes from the repository
	resumes, err := s.repository.GetResumesByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Include additional logic for skills and experience
	for _, resume := range resumes {
		// Fetch skills for each resume
		skills, err := s.repository.GetSkillsByResumeID(resume.ID)
		if err != nil {
			return nil, err
		}
		resume.Skills = skills

		// Calculate total experience in years and months
		totalExperience, err := s.repository.CalculateTotalExperience(resume.ID)
		if err != nil {
			return nil, err
		}
		resume.TotalExperience = totalExperience
	}

	return resumes, nil
}
