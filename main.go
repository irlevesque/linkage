package main

import (
	"fmt"
	"os"

	"github.com/irlevesque/linkage/pkg/linkage"
)

func main() {
	listenIP := os.Getenv("LISTEN_IP")
	listenPort := os.Getenv("LISTEN_PORT")

	linkage.LoadConfig()

	// defaults to localhost:8080 if not specified
	linkageApp := linkage.New(fmt.Sprintf("%s:%s", listenIP, listenPort))
	linkageApp.Serve()
}
