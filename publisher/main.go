package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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

func main() {
	dataPath := os.Getenv("PUBLISHER_DATA_PATH")
	order := ParseOrders(dataPath)

	nsUrl := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_CLIENT")

	fmt.Printf("URL: %s, Cluster_ID: %s, Client_ID: %s \n", nsUrl, nsCluster, nsClient)
	fmt.Println("Connecting to nats-streaming-server...")

	connection := NSConnect(nsUrl, nsCluster, nsClient)

	natsChannel := os.Getenv("NATS_CHANNEL")

	var delay time.Duration = 10 * time.Second

	for {
		message, err := json.Marshal(order)
		cherr.CheckErr(err)

		err = connection.Publish(natsChannel, message)
		cherr.CheckErr(err)

		fmt.Printf("Sent message with order's Id = %s", order.Order_uid)

		time.Sleep(delay)
	}
}
