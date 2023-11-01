package request

import (
	"backend-golang/businesses/profiles"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserProfile struct {
	Name       string `json:"name" validate:"NotEmpty"`
	Nip        string `json:"nip" validate:"required,NotEmpty"`
	Division   string `json:"division" validate:"required,NotEmpty"`
	Image_Path string `json:"image_path"`
}

func (req *UserProfile) ToDomain() *profiles.Domain {
	return &profiles.Domain{
		Name:       req.Name,
		Nip:        req.Nip,
		Division:   req.Division,
		Image_Path: req.Image_Path,
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

func (req UserProfile) Validate() error {
	return validateRequest(req)
}
