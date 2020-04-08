/*
 * This implements a gRPC Server that communicates with a gRPC
 * Client and returns a "Hello World" string in reply.
 *
 * Reference:
 * The implementation contained herein has been adapted from gRPC example code.
 * https://grpc.io/docs/quickstart/go/
 */

// Package main implements a gRPC server for Hello World service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alokswaincontact/demo-workspace/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld Server.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld Server
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Message Received from client.")
	return &pb.HelloReply{Message: "Hello World"}, nil
}

func main() {
	//Create a TCP socket and Listen for connection from Client
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Hello World gRPC Server Started")

	//Respond to the connection from Client with "Hello World" message
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
