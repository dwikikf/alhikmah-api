package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i any) map[string]string {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	errorMessages := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {

		field := err.Field()

		tag := err.Tag()

		param := err.Param()

		switch tag {
		case "required":
			errorMessages[field] = fmt.Sprintf("Field %s tidak boleh kosong.", field)
		case "numeric":
			errorMessages[field] = fmt.Sprintf("%s harus berupa angka.", field)
		case "min":
			errorMessages[field] = fmt.Sprintf("%s minimal %s karakter.", field, param)
		case "max":
			errorMessages[field] = fmt.Sprintf("%s maksimal %s karakter.", field, param)
		case "len":
			errorMessages[field] = fmt.Sprintf("%s harus tepat %s karakter.", field, param)
		case "oneof":
			errorMessages[field] = fmt.Sprintf("%s harus salah satu dari: %s.", field, param)
		// case "gt":
		// 	errorMessages[field] = fmt.Sprintf("%s harus lebih besar dari %s.", field, param)
		// case "alphanum":
		// 	errorMessages[field] = fmt.Sprintf("%s hanya boleh berisi huruf dan angka.", field)
		// case "email":
		// 	errorMessages[field] = fmt.Sprintf("Format %s tidak valid.", field)
		default:
			errorMessages[field] = fmt.Sprintf("Keliru pada field %s dengan kriteria: %s.", field, tag)
		}
	}

	return errorMessages
}
