package main

import (
	"io"
	"log"
	"time"

	pb "github.com/dlinh31/go-grpc/proto"
)

func (*helloServer) SayHelloBidirectionalStreaming (stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v", err)
		}
		log.Printf("Received stream: %v", req.Name)
		time.Sleep(2 * time.Second)
		if err := stream.Send(&pb.HelloResponse{Message: "Hello " + req.Name}); err != nil {
			log.Fatalf("Error while sending stream from server: %v", err)
		}

	}
}