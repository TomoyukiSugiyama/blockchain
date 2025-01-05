package server

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/transaction"
	pb "blockchain/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func StartRootServer(appAddress, nodeAddress string) {
	bc := blockchain.NewBlockchain()
	acc := InitAccount()
	bc.CreateGenesisBlock(acc)
	server := server{
		bc:          bc,
		accs:        acc,
		appAddress:  appAddress,
		nodeAddress: nodeAddress,
		nodes:       map[string]string{},
		tp:          transaction.NewTransactionPool(),
	}

	nodeListener, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		panic(err)
	}
	defer nodeListener.Close()
	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server)
	log.Printf("Starting node server on %s", nodeAddress)
	go s.Serve(nodeListener)

	appListener, err := net.Listen("tcp", appAddress)
	if err != nil {
		panic(err)
	}
	defer appListener.Close()
	c := grpc.NewServer()
	pb.RegisterBlockchainServer(c, &server)
	log.Printf("Starting application server on %s", appAddress)
	go c.Serve(appListener)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server exiting")
}
