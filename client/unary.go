package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dlinh31/go-grpc/proto"
)

func _(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("client cannot greet: %v", err)
	}
	log.Printf("%s", res.Message)
}