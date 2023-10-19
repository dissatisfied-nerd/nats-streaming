package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	cherr "github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"
)

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

	var order model.Orders
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
