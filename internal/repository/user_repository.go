package repository

import (
	"database/sql"

	"github.com/aidosgal/ei-jobs-core/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}

	r.db.QueryRow(`
		SELECT
			id,
			first_name,
			last_name,
			company_name,
			email,
			phone,
			avatar_url,
			role_id,
			password
		FROM
			users
		WHERE phone = ?
	`, phone).Scan(&user.Id, &user.FirstName, &user.LastName, &user.CompanyName, &user.Email, &user.Phone, &user.AvatarUrl, &user.RoleId, &user.Password)

	return user, nil
}

func (r *UserRepository) CreateUser(register *model.UserRegisterRequest) (*model.User, error) {
	user := &model.User{}
	query := `
        INSERT INTO users (first_name, last_name, role_id, email, phone, password)
        VALUES (?, ?, 1, ?, ?, ?)
    `

	result, err := r.db.Exec(query, register.FirstName, register.LastName, register.Email, register.Phone, register.Password)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	role_id := 1
	user.Id = int(lastID)
	user.FirstName = &register.FirstName
	user.LastName = &register.LastName
	user.Email = &register.Email
	user.Phone = &register.Phone
	user.Password = register.Password
	user.RoleId = &role_id

	return user, nil
}
