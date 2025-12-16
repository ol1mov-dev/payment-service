package main

import (
	"log"
	"net"
	"order-service/databases"
	"order-service/internal/handlers"

	orderv1 "github.com/ol1mov-dev/protos/pkg/order/v1"

	_ "slices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	HOST = "localhost"
	PORT = "8082"
)

func main() {
	err := databases.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	addr := net.JoinHostPort(HOST, PORT)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderv1.RegisterOrderV1ServiceServer(grpcServer, &handlers.OrderServer{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
