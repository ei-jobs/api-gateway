package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS ei_jobs")
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	_, err = db.Exec("USE ei_jobs")
	if err != nil {
		return fmt.Errorf("failed to switch database: %w", err)
	}

	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			first_name VARCHAR(50) NULL,
			last_name VARCHAR(50) NULL,
			company_name VARCHAR(50) NULL,
			avatar_url VARCHAR(50) NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			phone VARCHAR(100) NOT NULL UNIQUE,
			role_id INT,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS roles (
		    id INT AUTO_INCREMENT PRIMARY KEY,
		    name VARCHAR(50) NOT NULL UNIQUE
		)`,
	}

	for _, query := range tableQueries {
		_, err = db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	log.Println("Database migration completed successfully")
	return nil
}
