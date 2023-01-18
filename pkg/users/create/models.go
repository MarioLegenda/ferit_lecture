package create

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Email    string `validate:"required,email"`
}
