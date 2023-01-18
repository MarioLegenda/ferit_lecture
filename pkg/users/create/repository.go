package create

import (
	"dirStructureLecture/pkg/storage"
	"fmt"
)

type Repository[T any] interface {
	Create(user T) error
}

type userRepository[T any] struct {
	db storage.Storage
}

func (u userRepository[T]) Create(model T) error {
	fmt.Println(model)
	if res := u.db.DB().Create(model); res.Error != nil {
		return res.Error
	}

	return nil
}

func NewRepository[T any](db storage.Storage) Repository[T] {
	return userRepository[T]{db: db}
}
