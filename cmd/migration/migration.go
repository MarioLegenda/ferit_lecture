package main

import (
	blogsCreate "dirStructureLecture/pkg/blogs/create"
	"dirStructureLecture/pkg/storage"
	userCreate "dirStructureLecture/pkg/users/adding"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadEnv()
	db := db()

	if err := db.DB().AutoMigrate(userCreate.User{}); err != nil {
		log.Fatalln(err)
	}

	if err := db.DB().AutoMigrate(blogsCreate.Blog{}); err != nil {
		log.Fatalln(err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
}

func db() storage.Storage {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Zagreb",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := storage.NewStorage(dsn)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
