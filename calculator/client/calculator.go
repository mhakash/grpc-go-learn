package main

import (
	"context"
	"log"

	pb "github.com/mhakash/grpc-go-learn/calculator/proto"
)

func sum(c pb.CalculatorServiceClient) {
	log.Println("Invoked sum...")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 3, SecondNumber: 5})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("sum: %v", res.SumResult)
}
