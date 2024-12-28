package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dlinh31/go-grpc/proto"
)

func CallSayHelloClientStream(client pb.GreetServiceClient, nameLists *pb.NamesList){
	 stream, err := client.SayHelloClientStreaming(context.Background())
	 if err != nil {
		log.Fatalf("client cannot connect: %v", err)
	 }
	for _, name := range nameLists.Names {
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Error while sending stream: %v", err)
		}
		log.Printf("Sending stream with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Cannot close connection: %v", err)
	}
	log.Printf("Response: %v", res.Message)
}