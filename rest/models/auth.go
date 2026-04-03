package models

import (
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (l LoginRequest) ToUser() User {
	return User{
		Email:    l.Email,
		Password: l.Password,
	}
}

func (s SignupRequest) ToUser() User {
	return User{
		ID:       uuid.New(),
		Email:    s.Email,
		Password: s.Password,
	}
}
