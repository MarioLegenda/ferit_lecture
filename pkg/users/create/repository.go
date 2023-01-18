package create

import (
	"dirStructureLecture/pkg/storage"
)

type Repository[T any] interface {
	Create(model T) error
	Migrate() error
}

type userRepository[T any] struct {
	db storage.Storage
}

func (u userRepository[T]) Create(model T) error {
	if res := u.db.DB().Create(model); res.Error != nil {
		return res.Error
	}

	return nil
}

func (u userRepository[T]) Migrate() error {
	return u.db.DB().AutoMigrate(User{})
}

func NewRepository[T any](db storage.Storage) Repository[T] {
	return userRepository[T]{db: db}
}
