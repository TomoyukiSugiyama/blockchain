package main

import (
	"blockchain/server"
	"flag"
	"fmt"
)

func main() {
	serverType := flag.String("type", "master", "master or client")
	flag.Parse()

	if *serverType == "master" {
		server.StartServer()
	} else if *serverType == "client" {
		server.StartClientServer()
	} else {
		fmt.Println("Invalid server type")
	}
}
