package main

import (
	lectureHttp "dirStructureLecture/cmd/http"
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/users/create"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()
	db := runDb()

	srv := echo.New()

	srv.POST("/user", func(c echo.Context) error {
		var user request.User
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		handler := create.NewUserCreate(create.User{
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		}, create.NewRepository[create.User](db))

		createdUser, err := handler.Handle()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, createdUser)
	})

	lectureHttp.StartServer(srv, db)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
}

func runDb() *storage.PostgresDb {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Zagreb",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := storage.NewPostgresDb(dsn)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
