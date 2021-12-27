package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	service_b "github.com/Speakerkfm/test_svc/service_b/api"
)

type server struct {
	service_b.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *service_b.GetUserRequest) (*service_b.GetUserResponse, error) {
	return &service_b.GetUserResponse{
		User: &service_b.User{
			Id: req.GetId(),
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
	service_b.RegisterUserServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
