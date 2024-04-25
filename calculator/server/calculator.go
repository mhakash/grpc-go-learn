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
