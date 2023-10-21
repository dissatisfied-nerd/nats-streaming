package server

import (
	"os"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	NewDbclient(dbUser, dbPassword, dbName)

	nsURL := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_CLIENT")

	ns := NewNSConnection(nsURL, nsCluster, nsClient)

	ns.Listen()
}
