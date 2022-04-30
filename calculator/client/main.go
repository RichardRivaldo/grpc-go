package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/RichardRivaldo/grpc-go/calculator/proto"
)

var server_addr string = "0.0.0.0:9999"

func main() {
	// Dial the server and close the client with defer
	// gRPC expects SSL by default, we can bypass this by passing
	// the second parameter is to create credentials without SSL
	conn, err := grpc.Dial(server_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed dialing the gRPC server, reason %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)
	// add(client)
	// prime(client)
	// average(client)
	max(client)
}
