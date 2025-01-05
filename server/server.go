package server

import (
	"blockchain/internal/block"
	"blockchain/internal/blockchain"
	"blockchain/internal/state"
	"context"
	"io"
	"log"

	"blockchain/internal/account"
	"blockchain/internal/transaction"
	pb "blockchain/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedBlockchainServer
	pb.UnimplementedNodeServer
	bc          *blockchain.Blockchain
	accs        map[string]*account.Account
	nodes       map[string]string
	appAddress  string
	nodeAddress string
	tp          *transaction.TransactionPool
}

// Test function
func InitAccount() map[string]*account.Account {
	acc1 := account.CreateNewAccount("0000", "Alice", 1000)
	acc2 := account.CreateNewAccount("0001", "Bob", 1000)

	return map[string]*account.Account{acc1.Address: acc1, acc2.Address: acc2}
}

func (s *server) Bloadcast(_ context.Context, in *pb.Transaction) (*pb.Verify, error) {
	log.Printf("Received: %s", in.GetContent())
	tr := transaction.Transaction{}
	tr.FromJson(in.GetContent())
	log.Printf("Transaction: %s", tr.String())
	if !validateTransaction(tr) {
		return &pb.Verify{Valid: false}, nil
	}

	s.tp.Push(&tr)
	return &pb.Verify{Valid: true}, nil
}

func (s *server) BloadcastBlock(_ context.Context, in *pb.Block) (*pb.VerifyBlock, error) {
	log.Printf("Received: %s", in.GetContent())
	b := block.Block{}
	b.FromJson(in.GetContent())
	log.Printf("Block: %s", b.String())
	// TODO: Validate block
	s.bc.AddBlock(&b, b.Hash, s.accs)
	return &pb.VerifyBlock{Valid: true}, nil
}

func validateTransaction(tr transaction.Transaction) bool {
	if tr.From == "" || tr.To == "" || tr.Amount <= 0 {
		return false
	}
	return true
}

func (s *server) ResisterNode(_ context.Context, in *pb.ClientInfo) (*pb.NodeInfo, error) {
	log.Printf("Node id: %s", in.GetId())
	log.Printf("Node address: %s", in.GetAddress())
	s.nodes[in.GetId()] = in.GetAddress()
	return &pb.NodeInfo{Id: "server@0123", Address: s.nodeAddress}, nil
}

func (s *server) Sync(_ context.Context, in *pb.SyncInfo) (*pb.SyncReply, error) {
	log.Printf("Client id: %s", in.GetId())
	log.Printf("Sync mode: %s", in.GetType())

	clientAddress := s.nodes[in.GetId()]

	go func() {
		// Connect to client node
		conn, err := grpc.NewClient(clientAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewNodeClient(conn)
		stream, err := c.Upload(context.Background())
		if err != nil {
			log.Fatalf("could not upload: %v", err)
		}
		for _, state := range s.bc.State {
			log.Printf("Syncing: %s", state.String())
			err := stream.Send(&pb.FileChunk{Content: state.ToJson()})
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
		}
		_, err = stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("could not close: %v", err)
		}
		log.Printf("Sync Success")
	}()

	return &pb.SyncReply{Message: "Start to Sync"}, nil
}

func (s *server) Upload(stream pb.Node_UploadServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadStatus{Message: "Upload Success"})
		}

		if err != nil {
			log.Fatalf("could not upload: %v", err)
		}
		s.bc.State = append(s.bc.State, &state.State{})
		s.bc.State[len(s.bc.State)-1].FromJson(in.GetContent())
		s.accs = s.bc.State[len(s.bc.State)-1].Accounts
		log.Printf("Received: %s", in.GetContent())
		log.Printf("State: %s", s.bc.State[len(s.bc.State)-1].String())
	}

}
