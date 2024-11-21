package model

type Vacancy struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	City         string `json:"city"`
	Country      string `json:"country"`
	CompanyName  string `json:"company_name"`
	SalaryFrom   *int   `json:"salary_from"`
	SalaryTo     *int   `json:"salary_to"`
	SalaryPeriod string `json:"salary_period"`
	CreatedAt    string `json:"created_at"`
}

type OneVacancy struct {
	ID               int                     `json:"id"`
	Title            string                  `json:"title"`
	City             string                  `json:"city"`
	Country          string                  `json:"country"`
	CompanyName      string                  `json:"company_name"`
	SalaryFrom       *int                    `json:"salary_from"`
	SalaryTo         *int                    `json:"salary_to"`
	SalaryPeriod     string                  `json:"salary_period"`
	WorkFormat       string                  `json:"work_format"`
	WorkSchedule     string                  `json:"work_schedule"`
	Conditions       []*VacanyCondition      `json:"coniditions"`
	Requirements     []*VacanyRequirement    `json:"requirements"`
	Responsibilities []*VacanyResponsibility `json:"responsiblities"`
	CreatedAt        string                  `json:"created_at"`
}

type VacancyRequest struct {
	Title            string                         `json:"title"`
	City             string                         `json:"city"`
	Country          string                         `json:"country"`
	UserId           int                            `json:"user_id"`
	SpecializationId int                            `json:"specialization_id"`
	SalaryFrom       *int                           `json:"salary_from"`
	SalaryTo         *int                           `json:"salary_to"`
	SalaryPeriod     string                         `json:"salary_period"`
	WorkFormat       string                         `json:"work_format"`
	WorkSchedule     string                         `json:"work_schedule"`
	Conditions       []*VacanyConditionRequest      `json:"conditions"`
	Requirements     []*VacanyRequirementRequest    `json:"requirements"`
	Responsibilities []*VacanyResponsibilityRequest `json:"responsiblities"`
}

type VacanyCondition struct {
	ID        int    `json:"id"`
	VacancyId int    `json:"vacancy_id"`
	Icon      string `json:"icon"`
	Condition string `json:"condition"`
}

type VacanyResponsibility struct {
	ID             int    `json:"id"`
	VacancyId      int    `json:"vacancy_id"`
	Responsibility string `json:"responsibility"`
}

type VacanyRequirement struct {
	ID          int    `json:"id"`
	VacancyId   int    `json:"vacancy_id"`
	Requirement string `json:"requirement"`
}

type VacanyConditionRequest struct {
	Icon      string `json:"icon"`
	Condition string `json:"condition"`
}

type VacanyResponsibilityRequest struct {
	Responsibility string `json:"responsibility"`
}

type VacanyRequirementRequest struct {
	Requirement string `json:"requirement"`
}

type VacanyRequest struct {
	Requirement string `json:"requirement"`
}

type VacancyFilters struct {
	SpecializationID int
	Title            string
	City             string
	Country          string
	Salary           *int
}
