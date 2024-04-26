package main

import (
	"context"
	"io"
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

func primes(c pb.CalculatorServiceClient) {
	log.Printf("Invoked greeting many times...")

	req := &pb.PrimesRequest{Number: 30}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to greet many times: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}

		log.Printf("prime: %v", res.Primes)
	}
}
