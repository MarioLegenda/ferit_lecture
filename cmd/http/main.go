package main

import (
	"dirStructureLecture/cmd/http/blogs"
	"dirStructureLecture/cmd/http/redis"
	"dirStructureLecture/cmd/http/users"
	"dirStructureLecture/cmd/server"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	loadEnv()
	if _, err := redis.NewClient(); err != nil {
		log.Fatalln(err)
	}

	db := runDb()

	srv := echo.New()

	srv.POST("/user", users.CreateUserHandler(db))
	srv.GET("/user/:id", users.GetUserHandler(db))
	srv.POST("/blog", blogs.CreateBlogHandler(db))

	server.StartServer(srv, db)
}
