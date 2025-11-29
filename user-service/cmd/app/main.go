package main

import (
	"log"
	"net"
	"user-service/internal/handler"

	userV1 "github.com/ol1mov-dev/protos/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	addr := net.JoinHostPort(HOST, PORT)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userV1.RegisterUserV1ServiceServer(grpcServer, &handler.UserServerHandler{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {

		log.Fatalf("failed to serve: %v", err)
	}
}
