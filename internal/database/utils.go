package database

import "fmt"

type DatabaseConnectionString struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDb       string
}

func CreateConnectionString(dbConnString DatabaseConnectionString) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbConnString.PostgresUser,
		dbConnString.PostgresPassword,
		dbConnString.PostgresHost,
		dbConnString.PostgresPort,
		dbConnString.PostgresDb,
	)
}
