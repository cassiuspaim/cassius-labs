package main

import (
	"context"
	"log"
	"time"

	hellov1 "github.com/cassius-labs/go-grpc-buf-series/hello-grpc/gen/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	client := hellov1.NewGreeterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &hellov1.SayHelloRequest{Name: "Cassius"})
	if err != nil {
		log.Fatalf("SayHello error: %v", err)
	}

	log.Println("Response:", resp.GetMessage())
}
