package storage

type Repository[T any] interface {
	Create(model T) error
}

type repository[T any] struct {
	db Storage
}

func (u repository[T]) Create(model T) error {
	if res := u.db.DB().Create(model); res.Error != nil {
		return res.Error
	}

	return nil
}

func NewRepository[T any](db Storage) Repository[T] {
	return repository[T]{db: db}
}
