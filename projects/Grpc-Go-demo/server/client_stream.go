package main

import (
	"io"
	"log"

	pb "github.com/Devisrisamidurai/Grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with message: %v", req.Name)
		messages = append(messages, "Hello,", req.Name)
	}
}
