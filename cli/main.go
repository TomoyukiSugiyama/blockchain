package main

import (
	"blockchain/client"
	"blockchain/server"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "server", "server or app")
	serverType := flag.String("serverType", "root", "root or client")
	appAddress := flag.String("appAddress", "127.0.0.1:8080", "node address")
	targetRootAddress := flag.String("targetRootAddress", "127.0.0.1:9090", "node address")
	nodeAddress := flag.String("nodeAddress", "127.0.0.1:9090", "node address")
	targetNodeAddress := flag.String("targetNodeAddress", "127.0.0.1:8080", "target address")
	flag.Parse()

	if *mode == "server" {
		serverMode(*serverType, *targetRootAddress, *appAddress, *nodeAddress)
	} else if *mode == "app" {
		applicationMode(*targetNodeAddress)
	} else {
		fmt.Println("Invalid mode")
	}
}

func serverMode(serverType string, targetRootAddress string, appAddress string, nodeAddress string) {
	if serverType == "root" {
		server.StartRootServer(appAddress, nodeAddress)
	} else if serverType == "client" {
		server.StartClientServer(targetRootAddress, appAddress, nodeAddress)
	} else {
		fmt.Println("Invalid server type")
	}
}

func applicationMode(targetNodeAddress string) {
	client.Run(targetNodeAddress)
}
