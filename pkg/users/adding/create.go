package adding

import (
	"dirStructureLecture/pkg/helpers"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/types"
)

type Create struct {
	user       User
	repository storage.Repository[*User]
}

func (c Create) Validate() error {
	model := newUserValidationModel(c.user.Name, c.user.LastName, c.user.Email)

	return helpers.Validate(&model)
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

func NewUserCreate(user User, repository storage.Repository[*User]) types.Job[User] {
	return Create{user: user, repository: repository}
}
