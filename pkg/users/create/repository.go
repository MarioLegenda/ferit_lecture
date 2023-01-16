package create

import "dirStructureLecture/pkg/storage"

type Repository[T any] interface {
	Create(user T) error
}

type userRepository[T any] struct {
	db *storage.PostgresDb
}

func (u userRepository[T]) Create(model T) error {
	if res := u.db.DB().Create(model); res.Error != nil {
		return res.Error
	}

	return nil
}

func NewRepository[T any](db *storage.PostgresDb) Repository[T] {
	return userRepository[T]{db: db}
}
