package handlers

import (
	"context"
	"user-service/databases"

	userV1 "github.com/ol1mov-dev/protos/pkg/user/v1"
)

const DEFAULT_USER_ROLE = "buyer"

type UserServerHandler struct {
	userV1.UnimplementedUserV1ServiceServer
}

func (s *UserServerHandler) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserResponse, error) {
	DB, err := databases.Connect()
	if err != nil {
		return nil, err
	}

	defer DB.Close()

	var defaultUserRole uint32
	err = DB.QueryRow("SELECT id FROM roles WHERE name = $1", DEFAULT_USER_ROLE).Scan(&defaultUserRole)
	if err != nil {
		return nil, err
	}

	var userId uint32
	err = DB.QueryRow("INSERT INTO users (firstname, lastname, email, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.PhoneNumber,
		defaultUserRole,
	).Scan(&userId)

	if err != nil {
		return nil, err
	}

	return &userV1.CreateUserResponse{Id: userId}, nil
}

func (s *UserServerHandler) UpdateUser(ctx context.Context, in *userV1.UpdateUserRequest) (*userV1.UpdateUserResponse, error) {
	return &userV1.UpdateUserResponse{}, nil
}

func (s *UserServerHandler) DeleteUser(ctx context.Context, in *userV1.DeleteUserRequest) (*userV1.DeleteUserResponse, error) {
	return &userV1.DeleteUserResponse{}, nil
}
