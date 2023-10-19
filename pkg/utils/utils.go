package utils

import (
	"encoding/json"
	"os"

	cherr "github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"

	"github.com/nats-io/stan.go"
)

func NSConnect(nsUrl, nsCluster, nsClient string) stan.Conn {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))
	cherr.CheckErr(err)

	return connection
}

func ParseOrders(dataPath string) model.Order {
	file, err := os.ReadFile(dataPath)
	cherr.CheckErr(err)

	var orders model.Order

	err = json.Unmarshal(file, &orders)
	cherr.CheckErr(err)

	return orders
}
