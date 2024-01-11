package main

import (
	"log"
	"time"

	pb "github.com/satyajitnayk/simple-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		// simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
