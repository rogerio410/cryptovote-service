package presentation

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/application"
	pb "github.com/rogerio410/cryptovote-service/pkg/pb/voteservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VoteService struct {
	Application application.Application
}

func (vs VoteService) Vote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	value := req.Value.String()

	if err := vs.Application.Commands.Vote.Execute(ctx, req.Symbol, value); err != nil {
		result_err := status.New(codes.InvalidArgument, err.Error()).Err()
		return nil, result_err
	}

	response := &pb.VoteResponse{Response: true}

	return response, nil
}

func (vs VoteService) RemoveVote(ctx context.Context, req *pb.RemoveVoteRequest) (*pb.RemoveVoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	response := &pb.RemoveVoteResponse{Response: false}

	return response, nil
}

func (vs VoteService) GetAllCripto(ctx context.Context, req *pb.GetAllCriptoRequest) (*pb.GetAllCriptoResponse, error) {
	cryptos, err := vs.Application.Queries.AllCrypto.Execute(ctx)

	if err != nil {
		err_status := status.New(codes.Internal, "Something went wrong!").Err()
		return nil, err_status
	}

	// Manual convert model.CryptoCurrency into pb.CryptoCurrency
	pb_cryptos := make([]*pb.Cryptocurrency, len(cryptos))
	for i, v := range cryptos {
		pb_cryptos[i] = &pb.Cryptocurrency{Symbol: v.Symbol, Name: v.Name}
	}

	response := &pb.GetAllCriptoResponse{Cryptocurrencies: pb_cryptos}

	return response, nil
}
