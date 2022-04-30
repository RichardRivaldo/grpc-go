package main

import (
	"context"
	"io"
	"log"
	"time"

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

func spamGreet(c pb.GreetServiceClient) {
	log.Println("Spam Greet Invoked!")

	requests := []*pb.GreetRequest{
		{PersonName: "richard1"},
		{PersonName: "richard2"},
		{PersonName: "richard3"},
	}

	stream, err := c.SpamGreet(context.Background())

	if err != nil {
		log.Fatalf("Error when spamming greet, reason %v\n", err)
	}

	for i, req := range requests {
		log.Printf("Sending request of %d: %v\n", i, req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error when receiving spam reply, reason %v\n", err)
	}

	log.Println(res.Result)
}
