package main

import (
	"context"
	"log"

	pb "github.com/mhakash/grpc-go-learn/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Invoked greeting...")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Hakim"})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("greeting: %v", res.Result)
}
