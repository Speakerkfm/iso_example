package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	service_c "github.com/Speakerkfm/test_svc/service_c/api"
)

type server struct {
	service_c.UnimplementedPhoneServiceServer
}

func (s *server) CheckPhone(ctx context.Context, req *service_c.CheckPhoneRequest) (*service_c.CheckPhoneResponse, error) {
	return &service_c.CheckPhoneResponse{
		Exists: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service_c.RegisterPhoneServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
