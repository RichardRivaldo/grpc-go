package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet with request of %v\n", in)

	return &pb.GreetResponse{
		Result: "Greetings to you, " + in.GetPersonName(),
	}, nil
}

func (s *Server) StreamGreet(in *pb.GreetRequest, stream pb.GreetService_StreamGreetServer) error {
	log.Printf("Greet stream with request of %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Greetings %d, %s!", i, in.PersonName)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
