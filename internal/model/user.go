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

type UserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RoleId    string `json:"role_id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
