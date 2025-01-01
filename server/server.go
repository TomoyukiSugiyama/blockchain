package server

import (
	"blockchain/internal/blockchain"
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"blockchain/internal/account"
	"blockchain/internal/block"
	pb "blockchain/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const clientAddress = "127.0.0.1:8080"
const nodeAddress = "127.0.0.1:9090"

type server struct {
	pb.UnimplementedBlockchainServer
	pb.UnimplementedNodeServer
	bc   *blockchain.Blockchain
	accs map[string]*account.Account
}

// Test function
func InitAccount() map[string]*account.Account {
	acc1 := account.CreateNewAccount("0000", "Alice", 1000)
	acc2 := account.CreateNewAccount("0001", "Bob", 1000)
	log.Println(acc1.String())
	log.Println(acc2.String())

	return map[string]*account.Account{acc1.Id: acc1, acc2.Id: acc2}
}

func (s *server) ExecuteTrunsaction(_ context.Context, in *pb.TransactionRequest) (*pb.TransactionReply, error) {
	log.Printf("Transaction from %s to %s", in.GetFrom(), in.GetTo())
	log.Printf("Amount: %d", in.GetAmount())

	tr1 := block.CreateNewTransaction(0, in.GetFrom(), in.GetTo(), int(in.GetAmount()))
	trs := []block.Transaction{*tr1}
	s.bc.MineBlock("First Block", trs, s.accs)

	for _, acc := range s.accs {
		log.Println(acc.String())
	}
	message := "Transaction from " + s.accs[in.GetFrom()].Name + " to " + s.accs[in.GetTo()].Name + " with amount " + strconv.Itoa(int(in.GetAmount()))
	return &pb.TransactionReply{Message: message}, nil
}

func (s *server) ResisterNode(_ context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Node connected: %s", in.GetId())
	return &pb.JoinReply{Message: "Welcome to the blockchain"}, nil
}

func StartServer() {
	clientListener, err := net.Listen("tcp", clientAddress)
	if err != nil {
		panic(err)
	}
	defer clientListener.Close()

	bc := blockchain.NewBlockchain()
	accs := InitAccount()
	c := grpc.NewServer()
	pb.RegisterBlockchainServer(c, &server{bc: bc, accs: accs})
	log.Printf("Starting client server on %s", clientAddress)
	go c.Serve(clientListener)

	nodeListener, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		panic(err)
	}
	defer nodeListener.Close()
	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server{bc: bc, accs: accs})
	log.Printf("Starting node server on %s", nodeAddress)
	if err := s.Serve(nodeListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func StartClientServer() {
	conn, err := grpc.NewClient(nodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNodeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ResisterNode(ctx, &pb.JoinRequest{Id: "client0123"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.GetMessage())
}
