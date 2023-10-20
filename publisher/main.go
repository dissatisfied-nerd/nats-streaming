package publisher

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	cherr "github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	utl "github.com/dissatisfied-nerd/nats-streaming/pkg/utils"
)

func main() {
	dataPath := os.Getenv("PUBLISHER_DATA_PATH")
	order := utl.ParseOrders(dataPath)

	nsUrl := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_CLIENT")

	fmt.Printf("URL: %s, Cluster_ID: %s, Client_ID: %s \n", nsUrl, nsCluster, nsClient)
	fmt.Println("Connecting to nats-streaming-server...")

	connection := utl.NSConnect(nsUrl, nsCluster, nsClient)

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
