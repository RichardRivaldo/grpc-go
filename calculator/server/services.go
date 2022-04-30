package main

import (
	"context"
	"io"
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
	log.Printf("Server Stream - Prime request of %v\n", req)

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

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Client Stream - Average request")

	sum := 0
	n := 0

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: float32(sum) / float32(n),
			})
		}

		if err != nil {
			log.Fatalf("Failed receiving stream, reason %v\n", err)
		}

		sum += int(msg.Number)
		n++
	}
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Client Stream - Average request")

	var res int32 = 0

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed receiving stream, reason %v\n", err)
		}

		if msg.Number > res {
			res = msg.Number

			err := stream.Send(&pb.MaxResponse{
				Max: res,
			})

			if err != nil {
				log.Fatalf("Failed sending response, reason %v\n", err)
			}
		}
	}
}
