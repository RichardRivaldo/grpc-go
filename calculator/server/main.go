package main

import (
	"log"
	"net"

	pb "github.com/RichardRivaldo/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var server_addr string = "0.0.0.0:9999"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	// TCP Listener for server
	listener, err := net.Listen("tcp", server_addr)
	if err != nil {
		log.Fatalf("Failed listening the server, reason %v\n", err)
	}
	log.Printf("Listening the server on %s\n", server_addr)

	// gRPC Server serving the port listener
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed serving the server, reason %v\n", err)
	}
}
