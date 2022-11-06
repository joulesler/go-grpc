package main

import (
	"context"

	pb "grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Sum: request.Int1 + request.Int2,
	}, nil

}
