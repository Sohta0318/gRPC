package main

import (
	"context"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}
