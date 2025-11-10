package main

import (
	"fmt"
	"log"
	"os"

	"github.com/irlevesque/linkage/pkg/config"
	"github.com/irlevesque/linkage/pkg/handlers"
	"github.com/irlevesque/linkage/pkg/linkage"
)

func main() {
	listenIP := os.Getenv("LISTEN_IP")
	listenPort := os.Getenv("LISTEN_PORT")

	configManager, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	handlers.SetLinkConfigWriter(configManager)
	handlers.SetSearchEngineConfigWriter(configManager)

	// defaults to localhost:8080 if not specified
	linkageApp := linkage.New(fmt.Sprintf("%s:%s", listenIP, listenPort))
	linkageApp.Serve()
}
