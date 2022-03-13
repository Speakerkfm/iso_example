package main


import (
	"context"
	"log"
	"net/http"
	"time"

	service_b "github.com/Speakerkfm/iso_example/service_b/api"

	"google.golang.org/grpc"
)

const (
	serviceBHost = "localhost:8001"
	serviceCHost = "localhost:8002"
)

func main() {
	conn, err := grpc.Dial(serviceBHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service_b.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &service_b.GetUserRequest{Id: 3})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	http.DefaultClient.Do()
	log.Printf("Greeting: %+v", r.GetUser())
}
