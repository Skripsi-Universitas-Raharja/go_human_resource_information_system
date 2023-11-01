package request

import (
	"backend-golang/businesses/users"

	"strings"

	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegistration struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required,NotEmpty"`
	Nip      string `json:"nip" validate:"required,NotEmpty"`
	Division string `json:"division" validate:"required,NotEmpty"`
	Role     string `json:"role"`
}

func (req *UserLogin) ToDomainLogin() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserRegistration) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Nip:      req.Nip,
		Division: req.Division,
		Role:     req.Role,
	}
}

func validateRequest(req interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", NotEmpty)

	err := validate.Struct(req)

	return err
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	return inputData != ""
}

func (req *UserLogin) Validate() error {
	return validateRequest(req)
}

func (req *UserRegistration) Validate() error {
	return validateRequest(req)
}

// func (req UserProfile) Validate() error {
// 	return validateRequest(req)
// }
