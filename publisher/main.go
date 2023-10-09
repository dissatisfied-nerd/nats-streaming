package main

import (
	"fmt"
	"os"
)

func main() {
	natsUrl := os.Getenv("NATSURL")
	natsCluster := os.Getenv("NATSCLUSTER")
	natsClient := os.Getenv("NATSCLIENT")
	natsChannel := os.Getenv("NATSCHANNEL")

	fmt.Println(natsUrl, natsCluster, natsClient, natsChannel)
}
