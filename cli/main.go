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
	clientAddress := flag.String("clientAddress", "127.0.0.1:8080", "node address")
	rootAddress := flag.String("rootAddress", "127.0.0.1:9090", "node address")
	nodeAddress := flag.String("nodeAddress", "127.0.0.1:9090", "node address")
	targetNodeAddress := flag.String("address", "127.0.0.1:8080", "target address")
	flag.Parse()

	if *mode == "server" {
		serverMode(*serverType, *rootAddress, *clientAddress, *nodeAddress)
	} else if *mode == "client" {
		clientMode(*targetNodeAddress)
	} else {
		fmt.Println("Invalid mode")
	}
}

func serverMode(serverType string, rootAddress string, clientAddress string, nodeAddress string) {
	if serverType == "master" {
		server.StartServer(clientAddress, nodeAddress)
	} else if serverType == "client" {
		server.StartClientServer(rootAddress, clientAddress, nodeAddress)
	} else {
		fmt.Println("Invalid server type")
	}
}

func clientMode(targetNodeAddress string) {
	client.Run(targetNodeAddress)
}
