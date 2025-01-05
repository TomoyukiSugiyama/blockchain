package server

import (
	"blockchain/internal/block"
	"blockchain/internal/transaction"
	pb "blockchain/proto"
	"context"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *server) ExecuteTrunsaction(_ context.Context, in *pb.TransactionRequest) (*pb.TransactionReply, error) {
	log.Printf("Transaction from %s to %s", in.GetFrom(), in.GetTo())
	log.Printf("Amount: %d", in.GetAmount())

	tr1 := transaction.CreateNewTransaction(0, in.GetFrom(), in.GetTo(), int(in.GetAmount()))
	s.bloadcastTransaction(*tr1)
	s.tp.Push(tr1)
	trs := []transaction.Transaction{*tr1}
	b := s.bc.MineBlock("Execute Transaction To Create Block", trs, s.accs)
	s.bloadcastVerifyBlock(b)
	s.bc.AddBlock(b, b.Hash, s.accs)
	message := "Transaction from " + s.accs[in.GetFrom()].Name + " to " + s.accs[in.GetTo()].Name + " with amount " + strconv.Itoa(int(in.GetAmount()))
	return &pb.TransactionReply{Message: message}, nil
}

func (s *server) bloadcastTransaction(tr transaction.Transaction) {
	for _, node := range s.nodes {
		go func() {
			// Connect to client node
			conn, err := grpc.NewClient(node, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			n := pb.NewNodeClient(conn)
			b, err := n.Bloadcast(context.Background(), &pb.Transaction{Content: tr.ToJson()})
			if err != nil {
				log.Fatalf("could not bloadcast: %v", err)
			}
			log.Printf("Bloadcast: %v", b.GetValid())
		}()
	}
}

func (s *server) bloadcastVerifyBlock(b *block.Block) {
	for _, node := range s.nodes {
		go func() {
			// Connect to client node
			conn, err := grpc.NewClient(node, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			n := pb.NewNodeClient(conn)
			bloadcast, err := n.BloadcastBlock(context.Background(), &pb.Block{Content: b.ToJson()})
			if err != nil {
				log.Fatalf("could not bloadcast: %v", err)
			}
			log.Printf("Bloadcast: %v", bloadcast.GetValid())
		}()
	}
}
