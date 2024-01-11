package main

import (
	"log"
	"net"

	pb "github.com/satyajitnayk/simple-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", listner.Addr())
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
