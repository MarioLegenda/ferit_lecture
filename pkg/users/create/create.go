package create

import (
	"dirStructureLecture/pkg"
	"errors"
	validator "github.com/go-playground/validator/v10"
)

type Create struct {
	user       User
	repository Repository[*User]
}

func (c Create) Validate() error {
	validate := validator.New()

	if err := validate.Struct(&c.user); err != nil {
		return errors.New("There are some validation errors")
	}

	return nil
}

func (c Create) Authenticate() error {
	return nil
}

func (c Create) Authorize() error {
	return nil
}

func (c Create) Logic() (User, error) {
	if err := c.repository.Create(&c.user); err != nil {
		return User{}, err
	}

	return c.user, nil
}

func (c Create) Handle() (User, error) {
	if err := c.Validate(); err != nil {
		return User{}, err
	}

	if err := c.Authenticate(); err != nil {
		return User{}, err
	}

	if err := c.Authorize(); err != nil {
		return User{}, err
	}

	model, err := c.Logic()

	if err != nil {
		return User{}, err
	}

	return model, nil
}

func NewUserCreate(user User, repository Repository[*User]) pkg.Job[User] {
	return Create{user: user, repository: repository}
}
