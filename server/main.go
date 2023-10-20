package main

import (
	"os"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	NewDbclient(dbUser, dbPassword, dbName)
}
