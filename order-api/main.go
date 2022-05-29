package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	address_api "github.com/Speakerkfm/iso_example/address-api/api"
	shipment_api "github.com/Speakerkfm/iso_example/shipment-api/api"
	user_api "github.com/Speakerkfm/iso_example/user-api/api"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	userApiHost     = "user-api.localhost:82"
	addressApiHost  = "address-api.localhost:82"
	shipmentApiHost = "shipment-api.localhost:82"
)

func main() {
	rand.Seed(time.Now().Unix())
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	userApiConn, err := grpc.Dial(userApiHost,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userApiConn.Close()
	userApiClient := user_api.NewUserServiceClient(userApiConn)

	addressApiConn, err := grpc.Dial(addressApiHost, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer addressApiConn.Close()
	addressApiClient := address_api.NewAddressServiceClient(addressApiConn)

	shipmentApiConn, err := grpc.Dial(shipmentApiHost, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer shipmentApiConn.Close()
	shipmentApiClient := shipment_api.NewShipmentServiceClient(shipmentApiConn)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/CreateOrder", func(writer http.ResponseWriter, request *http.Request) {
		queryUserID := request.URL.Query().Get("user_id")
		userID, err := strconv.Atoi(queryUserID)
		if err != nil {
			handleError(writer, err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userApiResp, err := userApiClient.GetUser(ctx, &user_api.GetUserRequest{Id: int64(userID)})
		if err != nil {
			handleError(writer, fmt.Errorf("user-api: %w", err))
			return
		}

		addressApiResp, err := addressApiClient.GetUserAddress(ctx, &address_api.GetUserAddressRequest{UserId: int64(userID)})
		if err != nil {
			handleError(writer, fmt.Errorf("address-api: %w", err))
			return
		}

		shipmentApiResp, err := shipmentApiClient.CreateShipment(ctx, &shipment_api.CreateShipmentRequest{
			UserId:        int64(userID),
			UserName:      userApiResp.GetUser().GetName(),
			OrderId:       int64(rand.Int()),
			UserAddressId: addressApiResp.GetAddress().GetId(),
		})
		if err != nil {
			handleError(writer, fmt.Errorf("shippment-api: %w", err))
			return
		}

		httpResponse := struct {
			User     interface{} `json:"user"`
			Address  interface{} `json:"address"`
			Shipment interface{} `json:"shipment"`
		}{
			User:     userApiResp.GetUser(),
			Address:  addressApiResp.GetAddress(),
			Shipment: shipmentApiResp,
		}

		if err := json.NewEncoder(writer).Encode(httpResponse); err != nil {
			log.Printf("fail to encode http response: %+v, err: %s", httpResponse, err.Error())
		}
	})

	if err := http.ListenAndServe("localhost:8000", http.DefaultServeMux); err != nil {
		log.Fatalf("fail to listen and serve: %s", err.Error())
	}
}

func handleError(writer http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	errResp := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}
	if err := json.NewEncoder(writer).Encode(errResp); err != nil {
		log.Printf("fail to encode err response: %+v, err: %s", errResp, err.Error())
	}
}
