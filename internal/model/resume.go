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

type ResumeSkill struct {
	ID       int    `json:"id"`        // Skill ID
	ResumeID int    `json:"resume_id"` // ID of the resume this skill is related to
	Skill    string `json:"skill"`     // The skill name
}
