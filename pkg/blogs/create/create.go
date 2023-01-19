package create

import (
	"dirStructureLecture/pkg"
	"dirStructureLecture/pkg/storage"
	"errors"
	"github.com/go-playground/validator/v10"
)

type Create struct {
	blog       Blog
	repository storage.Repository[*Blog]
}

func (c Create) Validate() error {
	validate := validator.New()

	model := newBlogValidationModel(c.blog.Title, c.blog.Description, c.blog.Content, c.blog.UserID)

	if err := validate.Struct(&model); err != nil {
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

func (c Create) Logic() (Blog, error) {
	if err := c.repository.Create(&c.blog); err != nil {
		return Blog{}, err
	}

	return c.blog, nil
}

func (c Create) Handle() (Blog, error) {
	if err := c.Validate(); err != nil {
		return Blog{}, err
	}

	if err := c.Authenticate(); err != nil {
		return Blog{}, err
	}

	if err := c.Authorize(); err != nil {
		return Blog{}, err
	}

	model, err := c.Logic()

	if err != nil {
		return Blog{}, err
	}

	return model, nil
}

func NewBlogCreate(blog Blog, repository storage.Repository[*Blog]) pkg.Job[Blog] {
	return Create{blog: blog, repository: repository}
}
