package main

import (
	"fmt"
	"os"

	"github.com/irlevesque/linkage/pkg/linkage"
)

func main() {
	var listenIP string
	var listenPort string

	if v, ok := os.LookupEnv("LISTEN_IP"); ok {
		listenIP = v
	} else {
		listenIP = "localhost"
	}

	if v, ok := os.LookupEnv("LISTEN_PORT"); ok {
		listenPort = v
	} else {
		listenPort = "8080"
	}

	linkageApp := linkage.New(fmt.Sprintf("%s:%s", listenIP, listenPort))
	linkageApp.Serve()
}
