package model

type Resume struct {
	ID               int            `json:"id"`
	UserID           int            `json:"user_id"`
	DateOfBirth      string         `json:"date_of_birth"`
	Gender           string         `json:"gender"`
	SpecializationID int            `json:"specialization_id"`
	Description      string         `json:"description"`
	SalaryFrom       int            `json:"salary_from"`
	SalaryTo         int            `json:"salary_to"`
	SalaryPeriod     string         `json:"salary_period"`
	CreatedAt        string         `json:"created_at"`
	Skills           []*ResumeSkill `json:"skills"`
	TotalExperience  string         `json:"total_experience"`
}

type OneResume struct {
	ID               int                   `json:"id"`
	UserID           int                   `json:"user_id"`
	DateOfBirth      string                `json:"date_of_birth"`
	Gender           string                `json:"gender"`
	SpecializationID int                   `json:"specialization_id"`
	Description      string                `json:"description"`
	SalaryFrom       int                   `json:"salary_from"`
	SalaryTo         int                   `json:"salary_to"`
	SalaryPeriod     string                `json:"salary_period"`
	CreatedAt        string                `json:"created_at"`
	Skills           []*ResumeSkill        `json:"skills"`
	Organizations    []*ResumeOrganization `json:"organizations"`
	TotalExperience  string                `json:"total_experience"`
}

type ResumeSkill struct {
	ID       int    `json:"id"`
	ResumeID int    `json:"resume_id"`
	Skill    string `json:"skill"`
}

type ResumeOrganization struct {
	ID               int     `json:"id"`
	ResumeID         int     `json:"resume_id"`
	OrganizationName string  `json:"organization_name"`
	SpecializationID int     `json:"specialization_id"`
	Description      string  `json:"description"`
	StartMonth       string  `json:"start_month"`
	StartYear        string  `json:"start_year"`
	EndMonth         *string `json:"end_month"`
	EndYear          *string `json:"end_year"`
}
