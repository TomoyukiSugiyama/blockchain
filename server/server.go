package server

import (
	"blockchain/internal/blockchain"
	"bufio"
	"fmt"
	"net"
)

const address = "127.0.0.1:8080"

func StartServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	blockchain.NewBlockchain()
	fmt.Println("Starting server on", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// handle error
	}
	fmt.Println(status)
	fmt.Fprintf(conn, "HTTP/1.0 200 OK\r\n\r\n")

}
