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
	var user *model.User

	r.db.QueryRow(`
		SELECT
			id,
			first_name,
			last_name,
			comapny_name,
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
