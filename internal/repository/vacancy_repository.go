package repository

import (
	"context"
	"database/sql"

	"github.com/aidosgal/ei-jobs-core/internal/model"
)

type VacancyRepository struct {
	db *sql.DB
}

func NewVacancyRepository(db *sql.DB) *VacancyRepository {
	return &VacancyRepository{db: db}
}

func (r *VacancyRepository) GetVacancies(ctx context.Context, filters model.VacancyFilters) ([]model.Vacancy, error) {
	query := `
        SELECT
            v.id, v.title, v.city, v.country, u.company_name,
            v.salary_from, v.salary_to, v.salary_period, v.created_at
        FROM vacancies v
        LEFT JOIN users u ON v.user_id = u.id
        WHERE 1=1
    `
	args := []interface{}{}

	if filters.SpecializationID != 0 {
		query += " AND v.specialization_id = ?"
		args = append(args, filters.SpecializationID)
	}

	if filters.Title != "" {
		query += " AND v.title LIKE ?"
		args = append(args, "%"+filters.Title+"%")
	}

	if filters.City != "" {
		query += " AND v.city = ?"
		args = append(args, filters.City)
	}

	if filters.Country != "" {
		query += " AND v.country = ?"
		args = append(args, filters.Country)
	}

	if filters.Salary != nil {
		query += " AND (v.salary_from <= ? OR v.salary_to >= ?)"
		args = append(args, *filters.Salary, *filters.Salary)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []model.Vacancy
	for rows.Next() {
		var v model.Vacancy
		err := rows.Scan(
			&v.ID, &v.Title, &v.City, &v.Country, &v.CompanyName,
			&v.SalaryFrom, &v.SalaryTo, &v.SalaryPeriod, &v.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}
	return vacancies, nil
}

func (r *VacancyRepository) GetVacancyByID(ctx context.Context, id int) (*model.OneVacancy, error) {
	// First get the main vacancy information
	vacancy := &model.OneVacancy{}
	err := r.db.QueryRowContext(ctx, `
        SELECT
            v.id, v.title, v.city, v.country, u.company_name,
            v.salary_from, v.salary_to, v.salary_period,
            v.work_format, v.work_schedule, v.created_at
        FROM vacancies v
        LEFT JOIN users u ON v.user_id = u.id
        WHERE v.id = ?
    `, id).Scan(
		&vacancy.ID, &vacancy.Title, &vacancy.City, &vacancy.Country,
		&vacancy.CompanyName, &vacancy.SalaryFrom, &vacancy.SalaryTo,
		&vacancy.SalaryPeriod, &vacancy.WorkFormat, &vacancy.WorkSchedule,
		&vacancy.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Get conditions
	conditions, err := r.getVacancyConditions(ctx, id)
	if err != nil {
		return nil, err
	}
	vacancy.Conditions = conditions

	// Get requirements
	requirements, err := r.getVacancyRequirements(ctx, id)
	if err != nil {
		return nil, err
	}
	vacancy.Requirements = requirements

	// Get responsibilities
	responsibilities, err := r.getVacancyResponsibilities(ctx, id)
	if err != nil {
		return nil, err
	}
	vacancy.Responsibilities = responsibilities

	return vacancy, nil
}

func (r *VacancyRepository) getVacancyConditions(ctx context.Context, vacancyID int) ([]*model.VacanyCondition, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, vacancy_id, icon, condition_text
        FROM vacancy_conditions
        WHERE vacancy_id = ?
    `, vacancyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conditions []*model.VacanyCondition
	for rows.Next() {
		condition := &model.VacanyCondition{}
		err := rows.Scan(&condition.ID, &condition.VacancyId, &condition.Icon, &condition.Condition)
		if err != nil {
			return nil, err
		}
		conditions = append(conditions, condition)
	}
	return conditions, nil
}

func (r *VacancyRepository) getVacancyRequirements(ctx context.Context, vacancyID int) ([]*model.VacanyRequirement, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, vacancy_id, requirement
        FROM vacancy_requirements
        WHERE vacancy_id = ?
    `, vacancyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requirements []*model.VacanyRequirement
	for rows.Next() {
		requirement := &model.VacanyRequirement{}
		err := rows.Scan(&requirement.ID, &requirement.VacancyId, &requirement.Requirement)
		if err != nil {
			return nil, err
		}
		requirements = append(requirements, requirement)
	}
	return requirements, nil
}

func (r *VacancyRepository) getVacancyResponsibilities(ctx context.Context, vacancyID int) ([]*model.VacanyResponsibility, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, vacancy_id, responsibility
        FROM vacancy_responsibilities
        WHERE vacancy_id = ?
    `, vacancyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responsibilities []*model.VacanyResponsibility
	for rows.Next() {
		responsibility := &model.VacanyResponsibility{}
		err := rows.Scan(&responsibility.ID, &responsibility.VacancyId, &responsibility.Responsibility)
		if err != nil {
			return nil, err
		}
		responsibilities = append(responsibilities, responsibility)
	}
	return responsibilities, nil
}
