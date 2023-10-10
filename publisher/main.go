package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"

	model "publisher/dbmodel"

	"github.com/nats-io/stan.go"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseFile(dataPath string) model.Orders {
	file, err := os.ReadFile(dataPath)
	CheckErr(err)

	var orders model.Orders

	err = json.Unmarshal(file, &orders)
	CheckErr(err)

	return orders
}

func NatsConnect(natsUrl, natsCluster, natsClient string) stan.Conn {
	connection, err := stan.Connect(natsCluster, natsClient, stan.NatsURL(natsUrl))
	CheckErr(err)

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
		CheckErr(err)

		err = connection.Publish(natsChannel, message)
		CheckErr(err)
	}
}
