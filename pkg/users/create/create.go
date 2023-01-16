package create

import (
	"dirStructureLecture/pkg"
)

type Create[T User] struct {
	user       T
	repository Repository[T]
}

func (c Create[T]) Validate() error {
	return nil
}

func (c Create[T]) Authenticate() error {
	return nil
}

func (c Create[T]) Authorize() error {
	return nil
}

func (c Create[T]) Handle() (T, error) {
	if err := c.Validate(); err != nil {
		return T(User{}), err
	}

	if err := c.Authenticate(); err != nil {
		return T(User{}), err
	}

	if err := c.Authorize(); err != nil {
		return T(User{}), err
	}

	if err := c.repository.Create(c.user); err != nil {
		return T(User{}), err
	}

	return c.user, nil
}

func NewUserCreate[T User](user T, repository Repository[T]) pkg.Job[T] {
	return Create[T]{user: T(user), repository: repository}
}
