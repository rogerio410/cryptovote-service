package presentation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rogerio410/cryptovote-service/internal/application"
	pb "github.com/rogerio410/cryptovote-service/pkg/pb/voteservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VoteServiceServer struct {
	Application application.Application
	Ctx         context.Context
}

func (vs VoteServiceServer) Vote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	if req.Username == "" {
		err := status.New(codes.InvalidArgument, "Username cannot be empty").Err()
		return nil, err
	}

	value := req.Value.String()

	if err := vs.Application.Commands.Vote.Execute(ctx, req.Symbol, value, req.Username); err != nil {
		result_err := status.New(codes.InvalidArgument, err.Error()).Err()
		return nil, result_err
	}

	response := &pb.VoteResponse{Response: true}

	return response, nil
}

func (vs VoteServiceServer) RemoveVote(ctx context.Context, req *pb.RemoveVoteRequest) (*pb.RemoveVoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	if req.Username == "" {
		err := status.New(codes.InvalidArgument, "Username cannot be empty").Err()
		return nil, err
	}

	if err := vs.Application.Commands.RemoveVote.Execute(ctx, req.Symbol, req.Username); err != nil {
		result_err := status.New(codes.InvalidArgument, err.Error()).Err()
		return nil, result_err
	}

	response := &pb.RemoveVoteResponse{Response: true}

	return response, nil
}

func (vs VoteServiceServer) GetAllCrypto(ctx context.Context, req *emptypb.Empty) (*pb.GetAllCryptoResponse, error) {

	cryptos, err := vs.Application.Queries.GetAllCrypto.Execute(ctx)

	if err != nil {
		err_status := status.New(codes.Internal, "Something went wrong!").Err()
		return nil, err_status
	}

	// Manual converting from model.CryptoCurrency to pb.CryptoCurrency
	pb_cryptos := make([]*pb.Cryptocurrency, len(cryptos))
	for i, v := range cryptos {
		pb_cryptos[i] = &pb.Cryptocurrency{Symbol: v.Symbol, Name: v.Name, Votes: v.Votes}
	}

	response := &pb.GetAllCryptoResponse{Cryptocurrencies: pb_cryptos}

	return response, nil
}

func (vs VoteServiceServer) GetAllCryptoStream(req *emptypb.Empty, stream pb.VoteService_GetAllCryptoStreamServer) error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Stream ended! Error:\n", err)
		}
	}()

	for {
		cryptos, err := vs.Application.Queries.GetAllCrypto.Execute(stream.Context())

		if err != nil {
			return status.New(codes.NotFound, "It's was not possivel return all cryptos!").Err()
		}

		// Manual converting from model.CryptoCurrency to pb.CryptoCurrency
		pb_cryptos := make([]*pb.Cryptocurrency, len(cryptos))
		for i, v := range cryptos {
			pb_cryptos[i] = &pb.Cryptocurrency{Symbol: v.Symbol, Name: v.Name, Votes: v.Votes}
		}

		response := &pb.CryptoVotesResponse{Cryptocurrencies: pb_cryptos}
		stream.Send(response)

		time.Sleep(time.Second * 3)
		log.Println("Sending data")
	}
}

func (vs VoteServiceServer) GetCryptoBySymbolStream(req *pb.GetCryptoBySymbolStreamRequest, stream pb.VoteService_GetCryptoBySymbolStreamServer) error {

	for {
		crypto, err := vs.Application.Queries.GetCryptoBySymbol.Execute(stream.Context(), req.Symbol)

		if err != nil {
			return status.New(codes.NotFound, "Crypto not found with given code!").Err()
		}

		// Manual converting from model.CryptoCurrency to pb.CryptoCurrency
		pb_crypto := &pb.Cryptocurrency{
			Symbol: crypto.Symbol,
			Name:   crypto.Name,
			Votes:  crypto.Votes,
		}

		stream.Send(pb_crypto)

		time.Sleep(time.Second * 3)
		log.Println("Sending data")
	}
}
