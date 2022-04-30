package main

import (
	"context"
	"fmt"
	"io"
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

func (s *Server) SpamGreet(stream pb.GreetService_SpamGreetServer) error {
	log.Printf("Spam Greet\n")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error when streaming request, reason %v\n", err)
		}

		log.Printf("Receiving request of %v\n", req)
		res += fmt.Sprintf("Greetings, %s\n", req.PersonName)
	}
}

func (s *Server) MultiGreet(stream pb.GreetService_MultiGreetServer) error {
	log.Printf("Multi Greet\n")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error when streaming request, reason %v\n", err)
		}

		res := "Greetings, " + req.PersonName
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error when sending response, reason %v\n", err)
		}
	}
}
