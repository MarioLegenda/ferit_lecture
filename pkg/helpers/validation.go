package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ErrorMessages = map[string]map[string]string

func Validate[T any](model T) error {
	v := validator.New()
	messages := make(map[string]map[string]string)

	if err := v.Struct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			messages[field] = validationMessage(field, err.Tag())
		}
	}

	if len(messages) != 0 {
		return &ValidationError{messages: messages}
	}

	return nil
}

type ValidationError struct {
	messages ErrorMessages
}

func (e ValidationError) Error() string {
	return ""
}

func (e ValidationError) Messages() ErrorMessages {
	return e.messages
}

func validationMessage(field string, tag string) map[string]string {
	switch tag {
	case "required":
		t := make(map[string]string)
		t[tag] = fmt.Sprintf("Field %s is required", field)
		return t
	case "uuid":
		t := make(map[string]string)
		t[tag] = fmt.Sprintf("Field %s must be a valid uuid", field)
		return t
	case "email":
		t := make(map[string]string)
		t[tag] = fmt.Sprintf("Field %s must be a valid email", field)
		return t
	}

	return nil
}
