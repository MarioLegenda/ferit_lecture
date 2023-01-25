package getting

import (
	"dirStructureLecture/pkg"
	"dirStructureLecture/pkg/storage"
	"errors"
	"github.com/go-playground/validator/v10"
)

type GetById struct {
	user       UserId
	repository storage.Repository[*User]
}

func (c GetById) Validate() error {
	validate := validator.New()

	model := newUserIdValidationModel(c.user.ID)

	if err := validate.Struct(&model); err != nil {
		return errors.New("There are some validation errors")
	}

	return nil
}

func (c GetById) Authenticate() error {
	return nil
}

func (c GetById) Authorize() error {
	return nil
}

func (c GetById) Logic() (User, error) {
	var user User
	if err := c.repository.Get(c.user.ID, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (c GetById) Handle() (User, error) {
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

func NewGetById(user UserId, repository storage.Repository[*User]) pkg.Job[User] {
	return GetById{user: user, repository: repository}
}
