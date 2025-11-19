package main

import (
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

	linkageApp := linkage.New(listenIP, listenPort)
	linkageApp.Serve()
}
