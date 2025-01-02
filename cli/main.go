package main

import (
	"blockchain/client"
	"blockchain/server"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "server", "server or client")
	serverType := flag.String("type", "master", "master or client")
	targetNodeAddress := flag.String("address", "127.0.0.1:8080", "target address")
	flag.Parse()

	if *mode == "server" {
		serverMode(*serverType)
	} else if *mode == "client" {
		clientMode(*targetNodeAddress)
	} else {
		fmt.Println("Invalid mode")
	}
}

func serverMode(serverType string) {
	if serverType == "master" {
		server.StartServer()
	} else if serverType == "client" {
		server.StartClientServer()
	} else {
		fmt.Println("Invalid server type")
	}
}

func clientMode(targetNodeAddress string) {
	client.Run(targetNodeAddress)
}
