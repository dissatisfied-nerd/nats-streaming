package main

import (
	"encoding/json"
	"math/rand"
	"os"

	cherr "github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"
	"github.com/nats-io/stan.go"
)

func ParseFile(dataPath string) model.Orders {
	file, err := os.ReadFile(dataPath)
	cherr.CheckErr(err)

	var orders model.Orders

	err = json.Unmarshal(file, &orders)
	cherr.CheckErr(err)

	return orders
}

func NatsConnect(natsUrl, natsCluster, natsClient string) stan.Conn {
	connection, err := stan.Connect(natsCluster, natsClient, stan.NatsURL(natsUrl))
	cherr.CheckErr(err)

	return connection
}

func main() {
	dataPath := os.Getenv("PUBLISHER_DATA_PATH")
	orders := ParseFile(dataPath)

	natsUrl := os.Getenv("NATS_URL")
	natsCluster := os.Getenv("NATS_CLUSTER")
	natsClient := os.Getenv("NATS_CLIENT")

	connection := NatsConnect(natsUrl, natsCluster, natsClient)

	natsChannel := os.Getenv("NATS_CHANNEL")

	for {
		orders.Id = rand.Intn(100)

		message, err := json.Marshal(orders)
		cherr.CheckErr(err)

		err = connection.Publish(natsChannel, message)
		cherr.CheckErr(err)
	}
}
