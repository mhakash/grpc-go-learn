package main

import (
	"context"
	"fmt"
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

func average(c pb.CalculatorServiceClient) {
	log.Printf("Invoked average...")
	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("failed to average: %v", err)
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}
	fmt.Println("average:", res.Result)
}
