package users

import (
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/pkg/helpers"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/users/adding"
	"dirStructureLecture/pkg/users/getting"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUserHandler(db storage.Storage) func(e echo.Context) error {
	return func(c echo.Context) error {
		var user request.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		handler := adding.NewUserCreate(adding.User{
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		}, storage.NewRepository[*adding.User](db))

		createdUser, err := handler.Handle()

		if err != nil && err != (*helpers.ValidationError)(nil) {
			return c.JSON(http.StatusBadRequest, err.(*helpers.ValidationError).Messages())
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func GetUserHandler(db storage.Storage) func(e echo.Context) error {
	return func(c echo.Context) error {
		handler := getting.NewGetById(getting.UserId{
			ID: c.Param("id"),
		}, storage.NewRepository[*getting.User](db))

		fetchedUser, err := handler.Handle()

		if err != nil && err != (*helpers.ValidationError)(nil) {
			return c.JSON(http.StatusBadRequest, err.(*helpers.ValidationError).Messages())
		}

		return c.JSON(http.StatusOK, fetchedUser)
	}
}
