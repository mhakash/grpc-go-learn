package main

import (
	"context"
	"log"

	pb "github.com/mhakash/grpc-go-learn/greet/proto"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("received: %v", req)
	return &pb.GreetResponse{Result: "Hello " + req.FirstName}, nil
}
