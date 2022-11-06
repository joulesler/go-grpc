package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello" + request.FirstName,
	}, nil
}
