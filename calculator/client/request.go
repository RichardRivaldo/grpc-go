package main

import (
	"context"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/calculator/proto"
)

func add(client pb.CalculatorServiceClient) {
	log.Println("Add Service Invoked!")

	res, err := client.Add(context.Background(), &pb.SumRequest{
		FirstNumber:  17,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Error invoking add service, reason %v\n", err)
	}

	log.Println("Sum Result:", res.Sum)
}
