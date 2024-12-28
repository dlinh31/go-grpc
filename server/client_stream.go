package main

import (
	"io"
	"log"

	pb "github.com/dlinh31/go-grpc/proto"
)


func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Message: messages})
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v", err)
			return err
		}
		log.Printf("Received name: %v", req.Name)
		messages = append(messages, "Hello " + req.Name)
	}

}