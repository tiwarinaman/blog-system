package errors

import (
	"blog-system/internal/providers/validation"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errorList = make(map[string]string)

func Init() {
	errorList = map[string]string{}
}

func SetFromError(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMessage(fieldError.Tag()))
		}
	}

}

func Add(key string, value string) {
	errorList[strings.ToLower(key)] = value
}

func GetErrorMessage(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorList
}
