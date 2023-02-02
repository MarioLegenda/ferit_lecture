package main

import (
	blogsCreate "dirStructureLecture/pkg/blogs/adding"
	"dirStructureLecture/pkg/storage"
	userCreate "dirStructureLecture/pkg/users/adding"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := loadEnv(); err != nil {
		fmt.Println(err)
		return
	}

	db, err := db()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := migrate(db); err != nil {
		fmt.Println(err)
		return
	}
}

func migrate(db storage.Storage) error {
	if err := db.DB().AutoMigrate(userCreate.User{}); err != nil {
		return err
	}

	if err := db.DB().AutoMigrate(blogsCreate.Blog{}); err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	return nil
}

func db() (storage.Storage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Zagreb",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := storage.NewStorage(dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
