package deleting

import (
	"dirStructureLecture/pkg/helpers"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/types"
)

type DeleteById struct {
	model      UserID
	repository storage.Repository[*User]
}

func (c DeleteById) Validate() error {
	model := newUserIdValidationModel(c.model.ID)

	return helpers.Validate(&model)
}

func (c DeleteById) Authenticate() error {
	return nil
}

func (c DeleteById) Authorize() error {
	return nil
}

func (c DeleteById) Logic() (User, error) {
	if err := c.repository.Delete(c.model.ID, &User{}); err != nil {
		return User{}, err
	}

	return User{}, nil
}

func (c DeleteById) Handle() (User, error) {
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

func NewUserDeleteById(user UserID, repository storage.Repository[*User]) types.Job[User] {
	return DeleteById{model: user, repository: repository}
}
