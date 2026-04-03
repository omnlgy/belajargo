package models

import (
	"fmt"

	"example.com/rest/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

func (u User) CreateUser() error {
	query := `INSERT INTO users (id, email, password) VALUES (?, ?, ?)`

	_, err := db.DB.Exec(query, u.ID, u.Email, u.Password)

	if err != nil {
		return fmt.Errorf("Failed to create User %w", err)
	}

	return nil
}

func (u User) GetUser() (*User, error) {
	query := `SELECT id, email, password, createdAt, updatedAt FROM users WHERE email = ?`

	var user User
	err := db.DB.QueryRow(query, u.Email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("Failed to get User %w", err)
	}

	return &user, nil
}
