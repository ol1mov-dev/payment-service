package handler

import (
	"context"
	"user-service/database"

	userV1 "github.com/ol1mov-dev/protos/pkg/user/v1"
)

type UserServerHandler struct {
	userV1.UnimplementedUserV1ServiceServer
}

func (s *UserServerHandler) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserResponse, error) {
	DB, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer DB.Close()

	var newUserUUID string

	err = DB.QueryRow("INSERT INTO users (firstname, lastname, fathername, email, password, phone_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		req.FirstName,
		req.LastName,
		req.Email,
		req.User.PhoneNumber,
	).Scan(&newUserId)

	if err != nil {
		return nil, err
	}

	return &userV1.CreateUserResponse{UserId: newUserId}, nil
}

func (s *UserServerHandler) UpdateUser(ctx context.Context, in *userV1.UpdateUserRequest) (*userV1.UpdateUserResponse, error) {
	return &userV1.UpdateUserResponse{}, nil
}

func (s *UserServerHandler) DeleteUser(ctx context.Context, in *userV1.DeleteUserRequest) (*userV1.DeleteUserResponse, error) {
	return &userV1.DeleteUserResponse{}, nil
}
