package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	ID          string `gorm:"primarykey"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Content     string `validate:"required"`

	UserID string `validate:"required"`
	User   User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Email    string `validate:"required,email"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	return
}

func (u *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	return
}
