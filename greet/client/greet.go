package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("Invoked greeting many times...")

	req := &pb.GreetRequest{FirstName: "Hakim"}

	stream, err := c.GreetManyTimes(context.Background(), req)
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

		log.Printf("greeting: %v", res.Result)
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("Invoked long greeting...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Hakim"},
		{FirstName: "Jordan"},
		{FirstName: "Kobe"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("failed to long greet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("sending: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}

	log.Printf("long greeting:\n%v", res.Result)
}
