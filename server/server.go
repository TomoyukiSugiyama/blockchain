package server

import (
	"blockchain/internal/blockchain"
	"context"
	"log"
	"net"

	pb "blockchain/proto"

	"google.golang.org/grpc"
)

const address = "127.0.0.1:8080"

type server struct {
	pb.UnimplementedBlockchainServer
}

func (s *server) SayHello(_ context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetOp())
	return &pb.Reply{Message: "Hello " + in.GetOp()}, nil
}

func StartServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	blockchain.NewBlockchain()

	s := grpc.NewServer()
	pb.RegisterBlockchainServer(s, &server{})
	log.Printf("Starting server on %s", address)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
