package server

import (
	"blockchain/internal/account"
	"blockchain/internal/blockchain"
	"blockchain/internal/transaction"
	pb "blockchain/proto"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const clientId = "client@0123"

func StartClientServer(rootAddress, appAddress, nodeAddress string) {
	server := server{
		bc:          blockchain.NewBlockchain(),
		accs:        map[string]*account.Account{},
		appAddress:  appAddress,
		nodeAddress: nodeAddress,
		nodes:       map[string]string{},
		tp:          transaction.NewTransactionPool(),
	}

	conn, err := grpc.NewClient(rootAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNodeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Register node
	server.registerNode(ctx, c)

	// Start client node
	nodeListener, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		panic(err)
	}
	defer nodeListener.Close()
	clientServer := grpc.NewServer()
	pb.RegisterNodeServer(clientServer, &server)
	log.Printf("Starting client node on %s", nodeAddress)
	go clientServer.Serve(nodeListener)

	// Sync
	server.sync(ctx, c)

	// Start application server
	appListener, err := net.Listen("tcp", appAddress)
	if err != nil {
		panic(err)
	}
	defer appListener.Close()
	toClient := grpc.NewServer()
	pb.RegisterBlockchainServer(toClient, &server)
	log.Printf("Starting application server on %s", appAddress)
	go toClient.Serve(appListener)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server exiting")
}

func (s *server) registerNode(ctx context.Context, c pb.NodeClient) {
	r, err := c.ResisterNode(ctx, &pb.ClientInfo{Id: clientId, Address: s.nodeAddress})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Node id: %s", r.GetId())
	log.Printf("Node address: %s", r.GetAddress())
	s.nodes[r.GetId()] = r.GetAddress()
}

func (s *server) sync(ctx context.Context, c pb.NodeClient) {
	sync, err := c.Sync(ctx, &pb.SyncInfo{Id: clientId, Type: "full"})
	if err != nil {
		log.Fatalf("could not sync: %v", err)
	}
	log.Printf("Sync: %s", sync.GetMessage())
}
