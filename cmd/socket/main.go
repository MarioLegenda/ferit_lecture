package main

import (
	"dirStructureLecture/cmd/server"
	"dirStructureLecture/cmd/socket/users"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db := runDb()

	srv := echo.New()

	srv.GET("/user", users.CreateUserHandler(db))

	server.StartServer(srv, db)
}
