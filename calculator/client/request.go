package main

import (
	"context"
	"io"
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

func prime(client pb.CalculatorServiceClient) {
	log.Println("Prime Service Invoked!")

	stream, err := client.Prime(context.Background(), &pb.PrimeRequest{
		Number: 210,
	})

	if err != nil {
		log.Fatalf("Error when streaming prime, reason %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when streaming prime, reason %v\n", err)
		}

		log.Println(msg.Prime)
	}
}
