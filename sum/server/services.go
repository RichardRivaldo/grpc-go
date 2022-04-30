package main

import (
	"context"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/sum/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet with request of %v\n", req)

	return &pb.SumResponse{
		Sum: req.FirstNumber + req.SecondNumber,
	}, nil
}
