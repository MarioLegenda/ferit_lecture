package main

import (
	lectureHttp "dirStructureLecture/cmd/http"
	"dirStructureLecture/cmd/http/users"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db := runDb()

	srv := echo.New()

	srv.POST("/user", users.CreateUserHandler(db))

	lectureHttp.StartServer(srv, db)
}
