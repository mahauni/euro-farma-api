package database

import "fmt"

type PostgresDatabaseConnectionString struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDb       string
}

func CreatePostgresConnectionString(dbConnString PostgresDatabaseConnectionString) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbConnString.PostgresUser,
		dbConnString.PostgresPassword,
		dbConnString.PostgresHost,
		dbConnString.PostgresPort,
		dbConnString.PostgresDb,
	)
}
