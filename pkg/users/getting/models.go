package getting

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `validate:"required"`
	LastName string `validate:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserId struct {
	ID string `json:"id"`
}
