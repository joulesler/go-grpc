package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Greet Service Client")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Alice",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
