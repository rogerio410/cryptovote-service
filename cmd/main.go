package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rogerio410/cryptovote-service/internal/services"
	"github.com/rogerio410/cryptovote-service/pkg/pb/voteservice"
	"google.golang.org/grpc"
)

var (
	server   *grpc.Server
	listener net.Listener
)

func main() {
	fmt.Println("Hello Go")

	server = grpc.NewServer()

	initializeListener()

	voteservice.RegisterVoteServiceServer(server, &services.VoteService{})

	if err := server.Serve(listener); err != nil {
		panic("Failed to start gRPC Service")
	}

	fmt.Println("Server OK")

}

func initializeListener() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := fmt.Sprintf("localhost:%v", port)
	listener, err = net.Listen("tcp", address)

	if err != nil {
		msg := fmt.Sprintf("Failed to listen %v", err)
		panic(msg)
	}

	fmt.Println("Listener OK.")

}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	server.GracefulStop()
}
