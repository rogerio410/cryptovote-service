package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rogerio410/cryptovote-service/internal/presentation"
	"github.com/rogerio410/cryptovote-service/pkg/pb/voteservice"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	server       *grpc.Server
	listener     net.Listener
	globalCancel context.CancelFunc
	gctx         context.Context
)

func main() {
	fmt.Println("Crypto Vote Service")
	server = grpc.NewServer()

	initializeListener()

	// Create new CryptoApp with all Dependencies
	ctx := context.Background()
	gctx, globalCancel = context.WithCancel(ctx)
	application := presentation.NewCryptoApplication(ctx)
	vs := &presentation.VoteServiceServer{
		Application: application,
		Ctx:         gctx,
	}

	voteservice.RegisterVoteServiceServer(server, vs)

	go signalsListener(server)

	if err := server.Serve(listener); err != nil {
		panic("Failed to start gRPC Service")
	}

	fmt.Println("Server Clean Stop OK")

}

func initializeListener() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/deploy/app/cryptovote-service")
	// viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error loading config.env file")
		panic(err)
	}
	port, _ := viper.Get("PORT").(string)
	server, _ := viper.Get("SERVER").(string)

	address := fmt.Sprintf("%v:%v", server, port)
	listener, err = net.Listen("tcp", address)

	if err != nil {
		msg := fmt.Sprintf("Failed to listen %v", err)
		panic(msg)
	}

	fmt.Printf("Server Listening on %v:%v \n", server, port)

}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	globalCancel()
	server.GracefulStop()
}
