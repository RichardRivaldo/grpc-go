package main

import (
	"context"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet with request of %v\n", in)

	return &pb.GreetResponse{
		Result: "Greetings to you, " + in.GetPersonName(),
	}, nil
}
