package server

import (
	"blockchain/internal/blockchain"
	"context"
	"log"
	"net"
	"strconv"

	"blockchain/internal/account"
	"blockchain/internal/block"
	pb "blockchain/proto"

	"google.golang.org/grpc"
)

const address = "127.0.0.1:8080"

type server struct {
	pb.UnimplementedBlockchainServer
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

	tr1 := block.CreateNewTransaction(0, in.GetFrom(), in.GetTo(), 100)
	trs := []block.Transaction{*tr1}
	s.bc.MineBlock("First Block", trs, s.accs)

	for _, acc := range s.accs {
		log.Println(acc.String())
	}
	message := "Transaction from " + s.accs[in.GetFrom()].Name + " to " + s.accs[in.GetTo()].Name + " with amount " + strconv.Itoa(int(in.GetAmount()))
	return &pb.TransactionReply{Message: message}, nil
}

func StartServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	bc := blockchain.NewBlockchain()
	accs := InitAccount()
	s := grpc.NewServer()
	pb.RegisterBlockchainServer(s, &server{bc: bc, accs: accs})
	log.Printf("Starting server on %s", address)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
