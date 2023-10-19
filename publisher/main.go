package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

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

func NSConnect(nsUrl, nsCluster, nsClient string) stan.Conn {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))
	cherr.CheckErr(err)

	return connection
}

func main() {
	dataPath := os.Getenv("PUBLISHER_DATA_PATH")
	orders := ParseFile(dataPath)

	nsUrl := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_CLIENT")

	fmt.Printf("URL: %s, Cluster_ID: %s, Client_ID: %s \n", nsUrl, nsCluster, nsClient)
	fmt.Println("Connecting to nats-streaming-server...")

	connection := NSConnect(nsUrl, nsCluster, nsClient)

	natsChannel := os.Getenv("NATS_CHANNEL")

	var order model.Order
	var delay time.Duration = 10 * time.Second

	for {
		order.Id = rand.Intn(100)

		message, err := json.Marshal(orders)
		cherr.CheckErr(err)

		err = connection.Publish(natsChannel, message)
		cherr.CheckErr(err)

		fmt.Printf("Sent message with order's Id = %d \n", order.Id)

		time.Sleep(delay)
	}
}
