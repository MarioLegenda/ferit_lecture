package users

import (
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/users/create"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUserHandler(db storage.Storage) func(e echo.Context) error {
	return func(c echo.Context) error {
		var user request.User
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		handler := create.NewUserCreate(create.User{
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		}, storage.NewRepository[*create.User](db))

		createdUser, err := handler.Handle()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, createdUser)
	}
}
