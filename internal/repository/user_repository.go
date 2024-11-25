package repository

import (
	"database/sql"
	"fmt"
	"time"

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

func (r *UserRepository) GetUserById(id int) (*model.UserResponse, error) {
	user := &model.UserResponse{}

	err := r.db.QueryRow(`
        SELECT
            id,
            first_name,
            last_name,
            company_name,
            email,
            phone,
            avatar_url,
            role_id
        FROM
            users
        WHERE id = ?
    `, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.CompanyName, &user.Email, &user.Phone, &user.AvatarUrl, &user.RoleId)

	if err != nil {
		return nil, err
	}

	servicesRows, err := r.db.Query(`
        SELECT
            us.id,
            us.name,
            us.description,
            us.price,
            us.deadline
        FROM
            user_services us
        WHERE us.user_id = ?
    `, id)
	if err != nil {
		return nil, err
	}
	defer servicesRows.Close()

	for servicesRows.Next() {
		service := &model.AssistanceResponse{}
		err := servicesRows.Scan(
			&service.Id,
			&service.Name,
			&service.Description,
			&service.Price,
			&service.Deadline,
		)
		if err != nil {
			return nil, err
		}

		user.Services = append(user.Services, service)
	}

	if err := servicesRows.Err(); err != nil {
		return nil, err
	}

	resumesRows, err := r.db.Query(`
        SELECT
            id,
            user_id,
            date_of_birth,
            gender,
            specialization_id,
            description,
            salary_from,
            salary_to,
            salary_period,
            created_at
        FROM
            resumes
        WHERE user_id = ?
    `, id)
	if err != nil {
		return nil, err
	}
	defer resumesRows.Close()

	for resumesRows.Next() {
		resume := &model.Resume{}
		err := resumesRows.Scan(
			&resume.ID,
			&resume.UserID,
			&resume.DateOfBirth,
			&resume.Gender,
			&resume.SpecializationID,
			&resume.Description,
			&resume.SalaryFrom,
			&resume.SalaryTo,
			&resume.SalaryPeriod,
			&resume.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orgRows, err := r.db.Query(`
            SELECT
                id,
                oraganization_name,
                specialization_id,
                description,
                start_month,
                start_year,
                end_month,
                end_year
            FROM
                resume_organizations
            WHERE resume_id = ?
        `, resume.ID)
		if err != nil {
			return nil, err
		}
		defer orgRows.Close()

		totalExperienceMonths := 0
		for orgRows.Next() {
			org := &model.ResumeOrganization{}
			var startMonth, startYear, endMonth, endYear string
			err := orgRows.Scan(
				&org.ID,
				&org.OrganizationName,
				&org.SpecializationID,
				&org.Description,
				&startMonth,
				&startYear,
				&endMonth,
				&endYear,
			)
			if err != nil {
				return nil, err
			}

			startDate, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", startYear, startMonth))
			if err != nil {
				return nil, err
			}

			var endDate time.Time
			if endYear == "" || endMonth == "" {
				endDate = time.Now()
			} else {
				endDate, err = time.Parse("2006-01", fmt.Sprintf("%s-%s", endYear, endMonth))
				if err != nil {
					return nil, err
				}
			}

			duration := endDate.Sub(startDate)
			totalExperienceMonths += int(duration.Hours() / (24 * 30))
		}

		years := totalExperienceMonths / 12
		months := totalExperienceMonths % 12
		resume.TotalExperience = fmt.Sprintf("%d years %d months", years, months)

		user.Resumes = append(user.Resumes, resume)
	}

	if err := resumesRows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUsersByRoleId(roleID int) ([]*model.Company, error) {
	var companies []*model.Company

	rows, err := r.db.Query(`
        SELECT
            u.id,
            u.first_name,
            u.last_name,
            u.description,
            u.company_name,
            u.email,
            u.phone,
            u.avatar_url,
            u.role_id
        FROM
            users u
        WHERE u.role_id = ?
    `, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		company := &model.Company{}
		err := rows.Scan(
			&company.Id,
			&company.FirstName,
			&company.LastName,
            &company.Description,
			&company.CompanyName,
			&company.Email,
			&company.Phone,
			&company.AvatarUrl,
			&company.RoleId,
		)
		if err != nil {
			return nil, err
		}

		var priceFrom int
		err = r.db.QueryRow(`
            SELECT COALESCE(MIN(price), 0)
            FROM user_services
            WHERE user_id = ?
        `, company.Id).Scan(&priceFrom)
		if err != nil {
			return nil, err
		}
		company.PriceFrom = priceFrom

		var review float64
		err = r.db.QueryRow(`
            SELECT COALESCE(AVG(mark), 0)
            FROM reviews
            WHERE company_id = ? AND user_id = ?
        `, company.Id, company.Id).Scan(&review)
		if err != nil {
			return nil, err
		}
		company.Review = review

		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *UserRepository) CreateUser(register *model.UserRegisterRequest) (*model.User, error) {
	user := &model.User{}
	query := `
        INSERT INTO users (first_name, last_name, company_name, role_id, email, phone, password)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `

	result, err := r.db.Exec(query, register.FirstName, register.LastName, register.CompanyName, register.RoleId, register.Email, register.Phone, register.Password)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	role_id := register.RoleId
	user.Id = int(lastID)
	user.FirstName = &register.FirstName
	user.LastName = &register.LastName
	user.CompanyName = &register.CompanyName
	user.Email = &register.Email
	user.Phone = &register.Phone
	user.Password = register.Password
	user.RoleId = &role_id

	return user, nil
}
