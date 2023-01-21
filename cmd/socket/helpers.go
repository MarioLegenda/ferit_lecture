package main

import (
	"dirStructureLecture/pkg/storage"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
}

func runDb() storage.Storage {
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
