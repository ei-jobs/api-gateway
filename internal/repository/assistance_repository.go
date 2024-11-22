package repository

import (
	"context"
	"database/sql"

	"github.com/aidosgal/ei-jobs-core/internal/model"
)

type AssistanceRepository struct {
    db *sql.DB
}

func NewAssistanceRepository (db *sql.DB) *AssistanceRepository {
    return &AssistanceRepository{db: db}
}

func (r *AssistanceRepository) GetAssistancesByUserId(ctx context.Context, user_id int) ([]*model.AssistanceResponse, error) {
}

func (r *AssistanceRepository) StoreAssistance(ctx context.Context, assistance *model.AssistanceRequest) (*model.AssistanceRequest, error) {
    _, err := r.db.Exec(`
       INSERT INTO user_services (name, price, user_id, deadline, description)
       VALUES (?, ?, ?, ?, ?);
    `, assistance.Name, assistance.Price, assistance.UserId, assistance.Deadline, assistance.Description)
    if err != nil {
        return nil, err
    }

    return assistance, nil
}

func (r *AssistanceRepository) UpdateAssistance(ctx context.Context, assistance *model.AssistanceRequest, id int) (*model.AssistanceRequest, error) {
    _, err := r.db.Exec(`
        UPDATE user_services 
        SET 
            name = ?,
            price = ?,
            user_id = ?, 
            deadline = ?, 
            description = ?
        WHERE user_id = ?;
    `, assistance.Name, assistance.Price, assistance.UserId, assistance.Deadline, assistance.Description, id)
    if err != nil {
        return nil, err
    }

    return assistance, nil
}

func (r *AssistanceRepository) DeleteAssistance(ctx context.Context, id int) error {
    _, err := r.db.Exec(`
        DELETE FROM user_services 
        WHERE user_id = ?;
    `, id)
    if err != nil {
        return err
    }

    return nil
}
