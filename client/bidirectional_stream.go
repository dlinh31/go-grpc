package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dlinh31/go-grpc/proto"
)

func CallSayHelloBidirectionalStream (client pb.GreetServiceClient, namesList *pb.NamesList){
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("client cannot connect: %v", err)
	}

	waitc := make(chan struct{})

	go func(){

		for _, name := range namesList.Names{
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("Error while sending stream from client: %v", err)
			}
			log.Printf("Sending request with name: %v", name)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
		<-waitc
		log.Printf("Finished sending request")
		
		
	}()

	for{
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving stream in client: %v", err)
		}
		log.Printf("Received message: %v", message.Message)
	}
	close(waitc)

}