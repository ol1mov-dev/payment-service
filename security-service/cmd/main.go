package main

import (
	"context"
	"log"
	"net"
	"sync"

	userPb "user-service/pkg/api/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type userServer struct {
	userPb.UnimplementedUserServiceV1Server
	users map[string]*userPb.User
	mu    sync.RWMutex
}

func newUserServer() *userServer {
	// Инициализируем с тестовыми пользователями
	users := map[string]*userPb.User{
		"user_1": {
			Id:    "user_1",
			Name:  "John Doe",
			Email: "john@example.com",
		},
		"user_2": {
			Id:    "user_2",
			Name:  "Jane Smith",
			Email: "jane@example.com",
		},
	}
	return &userServer{
		users: users,
	}
}

func (s *userServer) GetUser(ctx context.Context, req *userPb.GetUserRequest) (*userPb.GetUserResponse, error) {
	log.Printf("GetUser called with ID: %s", req.UserId)

	s.mu.RLock()
	user, exists := s.users[req.UserId]
	s.mu.RUnlock()

	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", req.UserId)
	}

	return &userPb.GetUserResponse{
		User: user,
	}, nil
}

func main() {
	host := "localhost" // или "0.0.0.0" для доступа снаружи
	port := "50051"
	address := net.JoinHostPort(host, port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Регистрируем сервис
	userPb.RegisterUserServiceV1Server(grpcServer, newUserServer())

	// Включаем рефлексию для тестирования
	reflection.Register(grpcServer)

	log.Printf("User Service starting on %s", address)
	log.Printf("gRPC reflection enabled")
	log.Printf("Available methods:")
	log.Printf("- user.UserService.GetUser")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
