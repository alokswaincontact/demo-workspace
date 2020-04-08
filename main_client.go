/*
 * This implements a gRPC Client that communicates with a gRPC
 * Server which returns a "Hello World" string
 *
 * Reference:
 * The implementation contained herein has been adapted from gRPC example code.
 * https://grpc.io/docs/quickstart/go/
 */

// Package main implements a Client for Hello World service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/alokswaincontact/demo-workspace/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Message Received: %s", r.GetMessage())
}
