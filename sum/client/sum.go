package main

import (
	"context"
	pb "grpc-go-course/sum/proto"
	"log"
)

func Sum(c pb.AdditionClient) {
	log.Println("Sum Service Client")
	res, err := c.Sum(context.Background(), &pb.Request{
		Int1: 10,
		Int2: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Greeting: %d\n", res.Sum)
}
