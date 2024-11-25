package model

type User struct {
	Id          int     `json:"id"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	CompanyName *string `json:"company_name"`
	AvatarUrl   *string `json:"avatar_url"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Password    string  `json:"password"`
	RoleId      *int    `json:"role_id"`
}

type UserResponse struct {
	Id          int                   `json:"id"`
	FirstName   *string               `json:"first_name"`
	LastName    *string               `json:"last_name"`
	CompanyName *string               `json:"company_name"`
	AvatarUrl   *string               `json:"avatar_url"`
	Email       *string               `json:"email"`
	Phone       *string               `json:"phone"`
	Resumes     []*Resume             `json:"resumes"`
	Services    []*AssistanceResponse `json:"services"`
	RoleId      *int                  `json:"role_id"`
}

type UserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
    Description string `json:"description"`
	RoleId      int    `json:"role_id"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

type Company struct {
	Id          int     `json:"id"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	CompanyName *string `json:"company_name"`
    Description *string `json:"description"`
	AvatarUrl   *string `json:"avatar_url"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	PriceFrom   int     `json:"price_from"`
	Review      float64 `json:"review"`
	Password    string  `json:"password"`
	RoleId      *int    `json:"role_id"`
}
