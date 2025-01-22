package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Devisrisamidurai/Grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the RPC call
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	// Log the response message
	log.Printf("Greeting response: %s", res.Message)
}
