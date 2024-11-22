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

func (r *UserRepository) GetUserById(id int) (*model.User, error) {
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
        WHERE id = ?
    `, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.CompanyName, &user.Email, &user.Phone, &user.AvatarUrl, &user.RoleId, &user.Password)

    return user, nil
}

func (r *UserRepository) GetUsersByRoleId(roleID int) ([]*model.User, error) {
    var users []*model.User

    rows, err := r.db.Query(`
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
        WHERE role_id = ?
    `, roleID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        user := &model.User{}
        if err := rows.Scan(
            &user.Id,
            &user.FirstName,
            &user.LastName,
            &user.CompanyName,
            &user.Email,
            &user.Phone,
            &user.AvatarUrl,
            &user.RoleId,
            &user.Password,
        ); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}


func (r *UserRepository) CreateUser(register *model.UserRegisterRequest) (*model.User, error) {
	user := &model.User{}
	query := `
        INSERT INTO users (first_name, last_name, role_id, email, phone, password)
        VALUES (?, ?, ?, ?, ?, ?)
    `

	result, err := r.db.Exec(query, register.FirstName, register.LastName, register.RoleId, register.Email, register.Phone, register.Password)
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
