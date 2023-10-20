package main

import (
	"os"

	dbc "github.com/dissatisfied-nerd/nats-streaming/pkg/dbcontroller"
)

func main() {
	dbURL := os.Getenv("POSTGRES_URL")
	dbc.NewDbclient(dbURL)
}
