package subscriber

import (
	cherr "github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/nats-io/stan.go"
)

func NSConnect(nsUrl, nsCluster, nsClient string) stan.Conn {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))
	cherr.CheckErr(err)

	return connection
}
