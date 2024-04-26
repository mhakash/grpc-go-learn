package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/mhakash/grpc-go-learn/greet/proto"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("received: %v", req)
	return &pb.GreetResponse{Result: "Hello " + req.FirstName}, nil
}

func (s *server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greeter.GreetManyTimes invoked with: %v", req)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %v, number %v", req.FirstName, i)
		if err := stream.Send(&pb.GreetResponse{Result: res}); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("Greeter.LongGreet invoked...")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			return err
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
