package server

import (
	"blockchain/internal/block"
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
	"blockchain/internal/transaction"
	pb "blockchain/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedBlockchainServer
	pb.UnimplementedNodeServer
	bc            *blockchain.Blockchain
	accs          map[string]*account.Account
	nodes         map[string]string
	clientAddress string
	nodeAddress   string
	tp            *transaction.TransactionPool
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

func StartServer(clientAddress, nodeAddress string) {
	bc := blockchain.NewBlockchain()
	bc.CreateGenesisBlock()
	server := server{
		bc:            bc,
		accs:          InitAccount(),
		clientAddress: clientAddress,
		nodeAddress:   nodeAddress,
		nodes:         map[string]string{},
		tp:            transaction.NewTransactionPool(),
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

const clientId = "client@0123"

func StartClientServer(rootAddress, clientAddress, nodeAddress string) {
	server := server{
		bc:            blockchain.NewBlockchain(),
		accs:          map[string]*account.Account{},
		clientAddress: clientAddress,
		nodeAddress:   nodeAddress,
		nodes:         map[string]string{},
		tp:            transaction.NewTransactionPool(),
	}
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
	log.Printf("Node id: %s", r.GetId())
	log.Printf("Node address: %s", r.GetAddress())
	server.nodes[r.GetId()] = r.GetAddress()

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
