package client

import (
	"context"
	"log"
	"time"

	pb "blockchain/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(targetServerAddress string) {
	conn, err := grpc.NewClient(targetServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBlockchainClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ExecuteTrunsaction(ctx, &pb.TransactionRequest{From: "0000", To: "0001", Amount: 100})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.GetMessage())

}
