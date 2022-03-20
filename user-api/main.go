package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	user_api "github.com/Speakerkfm/iso_example/user-api/api"
)

type server struct {
	user_api.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *user_api.GetUserRequest) (*user_api.GetUserResponse, error) {
	return &user_api.GetUserResponse{
		User: &user_api.User{
			Id:      1,
			Name:    "test_user",
			Surname: "test_surname",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	user_api.RegisterUserServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
