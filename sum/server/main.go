package main

import (
	"log"
	"net"

	pb "grpc-go-course/sum/proto"

	"google.golang.org/grpc"
)

var addr string = "localhost:8082"

type Server struct {
	pb.AdditionServer
}

func main() {
	conn, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Could not create connection: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAdditionServer(grpcServer, &Server{})

	lisErr := grpcServer.Serve(conn)

	if lisErr != nil {
		log.Fatalf("Could not start server: %v\n", lisErr)
	}

	// defer conn.Close()

}
