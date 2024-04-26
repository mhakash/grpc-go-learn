package main

import (
	"context"
	"log"

	pb "github.com/mhakash/grpc-go-learn/calculator/proto"
)

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("received: %v", req)
	return &pb.SumResponse{SumResult: req.FirstNumber + req.SecondNumber}, nil
}

func (s *server) Primes(req *pb.PrimesRequest, steam pb.CalculatorService_PrimesServer) error {
	log.Printf("Calculator.Primes invoked with: %v", req)
	// return all prime factors
	n := req.Number
	var k int32 = 2
	for n > 1 {
		if n%k == 0 {
			if err := steam.Send(&pb.PrimesResponse{Primes: k}); err != nil {
				return err
			}
			n = n / k
		} else {
			k++
		}
	}

	return nil
}
