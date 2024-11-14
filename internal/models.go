package internal

import "github.com/go-playground/validator/v10"

type ApiResponse struct {
	Error   *string     `json:"error"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

type SignupParams struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"Email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me" validate:"required"`
}

var Validate = validator.New()
