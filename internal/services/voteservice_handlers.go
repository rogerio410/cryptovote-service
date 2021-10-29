package services

import (
	"context"

	pb "github.com/rogerio410/cryptovote-service/pkg/pb/voteservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VoteService struct{}

func (vs VoteService) UpVote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	response := &pb.VoteResponse{Response: true}

	return response, nil
}

func (vs VoteService) DownVote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	if req.Symbol == "" {
		err := status.New(codes.InvalidArgument, "Crypto Symbol cannot be empty").Err()
		return nil, err
	}

	response := &pb.VoteResponse{Response: false}

	return response, nil
}
