package server

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/state"
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"blockchain/internal/account"
	"blockchain/internal/block"
	pb "blockchain/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedBlockchainServer
	pb.UnimplementedNodeServer
	bc    *blockchain.Blockchain
	accs  map[string]*account.Account
	nodes map[string]string
}

// Test function
func InitAccount() map[string]*account.Account {
	acc1 := account.CreateNewAccount("0000", "Alice", 1000)
	acc2 := account.CreateNewAccount("0001", "Bob", 1000)

	return map[string]*account.Account{acc1.Address: acc1, acc2.Address: acc2}
}

func (s *server) ExecuteTrunsaction(_ context.Context, in *pb.TransactionRequest) (*pb.TransactionReply, error) {
	log.Printf("Transaction from %s to %s", in.GetFrom(), in.GetTo())
	log.Printf("Amount: %d", in.GetAmount())

	tr1 := block.CreateNewTransaction(0, in.GetFrom(), in.GetTo(), int(in.GetAmount()))
	trs := []block.Transaction{*tr1}
	s.bc.MineBlock("Execute Transaction To Create Block", trs, s.accs)

	message := "Transaction from " + s.accs[in.GetFrom()].Name + " to " + s.accs[in.GetTo()].Name + " with amount " + strconv.Itoa(int(in.GetAmount()))
	return &pb.TransactionReply{Message: message}, nil
}

func (s *server) ResisterNode(_ context.Context, in *pb.ClientInfo) (*pb.JoinReply, error) {
	log.Printf("Node connected: %s", in.GetId())
	log.Printf("Node address: %s", in.GetAddress())
	s.nodes = make(map[string]string)
	s.nodes[in.GetId()] = in.GetAddress()
	return &pb.JoinReply{Message: "Welcome to the blockchain"}, nil
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

func StartServer(clientAddress, nodeAddress string) {
	bc := blockchain.NewBlockchain()
	bc.CreateGenesisBlock()
	accs := InitAccount()
	server := server{bc: bc, accs: accs}

	nodeListener, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		panic(err)
	}
	defer nodeListener.Close()
	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server)
	log.Printf("Starting node server on %s", nodeAddress)
	go s.Serve(nodeListener)

	clientListener, err := net.Listen("tcp", clientAddress)
	if err != nil {
		panic(err)
	}
	defer clientListener.Close()
	c := grpc.NewServer()
	pb.RegisterBlockchainServer(c, &server)
	log.Printf("Starting client server on %s", clientAddress)
	go c.Serve(clientListener)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server exiting")
}

const clientId = "client0123"

func StartClientServer(rootAddress, clientAddress, nodeAddress string) {
	server := server{bc: blockchain.NewBlockchain(), accs: map[string]*account.Account{}}
	// Register node
	conn, err := grpc.NewClient(rootAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNodeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ResisterNode(ctx, &pb.ClientInfo{Id: clientId, Address: nodeAddress})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Resister: %s", r.GetMessage())

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
	s, err := c.Sync(ctx, &pb.SyncInfo{Id: clientId, Type: "full"})
	if err != nil {
		log.Fatalf("could not sync: %v", err)
	}
	log.Printf("Sync: %s", s.GetMessage())

	// Start transaction server
	clientListener, err := net.Listen("tcp", clientAddress)
	if err != nil {
		panic(err)
	}
	defer clientListener.Close()
	toClient := grpc.NewServer()
	pb.RegisterBlockchainServer(toClient, &server)
	log.Printf("Starting client server on %s", clientAddress)
	go toClient.Serve(clientListener)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server exiting")
}
