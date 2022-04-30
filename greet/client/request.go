package main

import (
	"context"
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
