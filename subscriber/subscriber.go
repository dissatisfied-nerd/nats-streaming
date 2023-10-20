package subscriber

import (
	"log"

	"github.com/nats-io/stan.go"
)

func NSConnect(nsUrl, nsCluster, nsClient string) stan.Conn {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
