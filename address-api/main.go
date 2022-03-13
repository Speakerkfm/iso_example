package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	address_api "github.com/Speakerkfm/iso_example/address-api/api"
)

type server struct {
	address_api.UnimplementedAddressServiceServer
}

func (s *server) GetUserAddress(ctx context.Context, req *address_api.GetUserAddressRequest) (*address_api.GetUserAddressResponse, error) {
	return &address_api.GetUserAddressResponse{
		Address: &address_api.Address{
			Id:      99,
			City:    "Perm",
			ZipCode: "614000",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	address_api.RegisterAddressServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
