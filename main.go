package main

import (
	lectureHttp "dirStructureLecture/cmd/http"
	"dirStructureLecture/cmd/http/blogs"
	"dirStructureLecture/cmd/http/users"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db := runDb()

	srv := echo.New()

	srv.POST("/user", users.CreateUserHandler(db))
	srv.GET("/user/:id", users.GetUserHandler(db))
	srv.POST("/blog", blogs.CreateBlogHandler(db))

	lectureHttp.StartServer(srv, db)
}
