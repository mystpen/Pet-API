package dto

type RegistrationRequest struct {
	UserName string  `json:"username" binding:"required" validate:"min=8,containsany=!@#?*"`
	Email    string  `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required" validate:"min=8"`
}