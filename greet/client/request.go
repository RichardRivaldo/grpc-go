package main

import (
	"context"
	"io"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/greet/proto"
)

func greet(c pb.GreetServiceClient) {
	log.Println("Greet Invoked!")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		PersonName: "Richard",
	})

	if err != nil {
		log.Fatalf("Error when greeting, reason %v\n", err)
	}

	log.Println(res.Result)
}

func streamGreet(c pb.GreetServiceClient) {
	log.Println("Greet Stream Invoked!")

	req := &pb.GreetRequest{
		PersonName: "Richard",
	}

	stream, err := c.StreamGreet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error when streaming greet, reason %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when streaming greet, reason %v\n", err)
		}

		log.Println(msg)
	}
}
