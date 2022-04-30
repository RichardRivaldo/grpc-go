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

func average(client pb.CalculatorServiceClient) {
	log.Println("Prime Service Invoked!")

	requests := []*pb.AverageRequest{
		{Number: 1},
		{Number: 3},
		{Number: 2},
		{Number: 4},
	}

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("Error when streaming average, reason %v\n", err)
	}

	for i, req := range requests {
		log.Printf("Sending requests of %d: %v\n", i, req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error when receiving result, reason %v\n", err)
	}

	log.Println(res.Average)
}
