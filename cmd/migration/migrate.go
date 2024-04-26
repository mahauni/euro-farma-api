package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	dbUtils "github.com/mahauni/euro-farma-api/internal/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load the .env: %v\n", err)
	}

	currDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get the currDir: %v\n", err)
		os.Exit(1)
	}

	dbConnStr := dbUtils.PostgresDatabaseConnectionString{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
	}
	connString := dbUtils.CreatePostgresConnectionString(dbConnStr)
	// connString := fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL"))

	m, err := migrate.New(
		fmt.Sprintf("file://%s/internal/database/migrations", currDir),
		connString,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
