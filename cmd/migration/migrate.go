package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load the .env: %v\n", err)
		os.Exit(1)
	}

	currDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get the currDir: %v\n", err)
		os.Exit(1)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/internal/database/migrations", currDir),
		fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL")),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
