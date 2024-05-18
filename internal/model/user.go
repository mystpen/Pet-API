package model

import (
	"errors"

	uuid "github.com/google/uuid"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrRecordNotFound = errors.New("record not found")
	ErrNoMatch        = errors.New("no match data")
)

type User struct {
	ID            uuid.UUID `json:"id"`
	UserName      string    `json:"username" binding:"required" validate:"min=8,containsany=!@#?*"`
	Email         string    `json:"email" binding:"required,email"`
	Password      []byte    `json:"password" binding:"required" validate:"min=8"`
	PlainPassword string    `json:"-"`
}
