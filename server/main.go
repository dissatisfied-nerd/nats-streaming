package main

import (
	"os"

	"time"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	db := NewDbclient(dbUser, dbPassword, dbName)

	nsURL := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_SUBSCRIBER")
	nsChannel := os.Getenv("NATS_CHANNEL")

	ns := NewNSConnection(nsURL, nsCluster, nsClient)

	ns.Channel = nsChannel

	ns.Listen(db)

	time.Sleep(10 * time.Second)
}
