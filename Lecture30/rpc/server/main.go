package main

import (
	"log"
	"net"

	"L30/news"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen on port 8000: %v", err)
	}
	s := news.Server{}
	grpcServer := grpc.NewServer()

	news.RegisterNewsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8000: %s", err)
	}
}
