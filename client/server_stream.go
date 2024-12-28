package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dlinh31/go-grpc/proto"
)


func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Sending names started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")

}