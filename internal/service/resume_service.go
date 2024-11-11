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

func (s *ResumeService) GetResumesByUserID(userID int) ([]*model.Resume, error) {
	resumes, err := s.repository.GetResumesByUserID(userID)
	if err != nil {
		return nil, err
	}
	for _, resume := range resumes {
		skills, err := s.repository.GetSkillsByResumeID(resume.ID)
		if err != nil {
			return nil, err
		}
		resume.Skills = skills
		totalExperience, err := s.repository.CalculateTotalExperience(resume.ID)
		if err != nil {
			return nil, err
		}
		resume.TotalExperience = totalExperience
	}
	return resumes, nil
}

func (s *ResumeService) CreateResume(resume *model.Resume) (*model.Resume, error) {
	createdResume, err := s.repository.CreateResume(resume)
	if err != nil {
		return nil, err
	}

	skills, err := s.repository.GetSkillsByResumeID(createdResume.ID)
	if err != nil {
		return nil, err
	}
	createdResume.Skills = skills

	totalExperience, err := s.repository.CalculateTotalExperience(createdResume.ID)
	if err != nil {
		return nil, err
	}
	createdResume.TotalExperience = totalExperience

	return createdResume, nil
}

func (s *ResumeService) UpdateResume(resume *model.Resume) (*model.Resume, error) {
	updatedResume, err := s.repository.UpdateResume(resume)
	if err != nil {
		return nil, err
	}

	skills, err := s.repository.GetSkillsByResumeID(updatedResume.ID)
	if err != nil {
		return nil, err
	}
	updatedResume.Skills = skills

	totalExperience, err := s.repository.CalculateTotalExperience(updatedResume.ID)
	if err != nil {
		return nil, err
	}
	updatedResume.TotalExperience = totalExperience

	return updatedResume, nil
}

func (s *ResumeService) DeleteResume(resumeID int) error {
	return s.repository.DeleteResume(resumeID)
}
