package main

import (
	"context"
	"log"

	pb "github.com/RichardRivaldo/grpc-go/calculator/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Unary - Sum with request of %v\n", req)

	return &pb.SumResponse{
		Sum: req.FirstNumber + req.SecondNumber,
	}, nil
}

func (s *Server) Prime(req *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Stream - Prime request of %v\n", req)

	factor := 2
	current := req.Number

	for current > 1 {
		if current%int32(factor) == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: int32(factor),
			})
			current = current / int32(factor)
		} else {
			factor += 1
		}
	}

	return nil
}
