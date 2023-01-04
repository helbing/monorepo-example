package main

import (
	"context"
	"log"
	"net"

	greetv1 "github.com/helbing/monorepo-example/gen/go/greet/v1"
	"github.com/helbing/monorepo-example/packages/tool"
	"google.golang.org/grpc"
)

type Server struct {
	greetv1.UnimplementedGreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	return &greetv1.GreetResponse{
		Greeting: tool.Greeting("world"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8881")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	greetv1.RegisterGreetServiceServer(grpcServer, &Server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
