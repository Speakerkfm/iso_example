package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	shipment_api "github.com/Speakerkfm/iso_example/shipment-api/api"
)

type server struct {
	shipment_api.UnimplementedShipmentServiceServer
}

func (s *server) CreateShipment(ctx context.Context, req *shipment_api.CreateShipmentRequest) (*shipment_api.CreateShipmentResponse, error) {
	return &shipment_api.CreateShipmentResponse{
		Id: 15,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	shipment_api.RegisterShipmentServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
