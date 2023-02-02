package blogs

import (
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/pkg/blogs/adding"
	"dirStructureLecture/pkg/helpers"
	"dirStructureLecture/pkg/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBlogHandler(db storage.Storage) func(e echo.Context) error {
	return func(c echo.Context) error {
		var blog request.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		handler := adding.NewBlogCreate(adding.Blog{
			Title:       blog.Title,
			Content:     blog.Content,
			Description: blog.Description,
			UserID:      blog.UserID,
		}, storage.NewRepository[*adding.Blog](db))

		createBlog, err := handler.Handle()

		if err != nil && err != (*helpers.ValidationError)(nil) {
			return c.JSON(http.StatusBadRequest, err.(*helpers.ValidationError).Messages())
		}

		return c.JSON(http.StatusCreated, createBlog)
	}
}
