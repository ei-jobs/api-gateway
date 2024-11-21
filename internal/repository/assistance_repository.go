package repository

import "database/sql"

type AssistanceRepository struct {
    db *sql.DB
}

func NewAssistanceRepository (db *sql.DB) *AssistanceRepository {
    return &AssistanceRepository{db: db}
}
