package main

import (
	"log"
	"net"

	pb "github.com/mhakash/grpc-go-learn/calculator/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

type server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
