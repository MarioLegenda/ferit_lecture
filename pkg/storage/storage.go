package storage

import "gorm.io/gorm"

type Storage interface {
	DB() *gorm.DB
	Connect(dsn string) error
}
