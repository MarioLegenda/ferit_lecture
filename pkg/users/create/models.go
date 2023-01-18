package create

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID     string `gorm:"primaryKey"`
	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Email    string `validate:"required,email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}
