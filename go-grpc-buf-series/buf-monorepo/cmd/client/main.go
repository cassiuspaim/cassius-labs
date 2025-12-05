package main

import (
	"context"
	"log"
	"time"

	billingv1 "acme/gen/billingv1"
	userv1 "acme/gen/userv1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Call UserService
	userConn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial user: %v", err)
	}
	defer userConn.Close()

	userClient := userv1.NewUserServiceClient(userConn)
	userRes, err := userClient.GetUser(ctx, &userv1.GetUserRequest{Id: "user-123"})
	if err != nil {
		log.Fatalf("GetUser: %v", err)
	}
	log.Printf("UserService.GetUser => user=%+v error=%+v", userRes.GetUser(), userRes.GetError())

	// Call BillingService
	billingConn, err := grpc.DialContext(ctx, "localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial billing: %v", err)
	}
	defer billingConn.Close()

	billingClient := billingv1.NewBillingServiceClient(billingConn)
	invRes, err := billingClient.GetInvoice(ctx, &billingv1.GetInvoiceRequest{Id: "inv-999"})
	if err != nil {
		log.Fatalf("GetInvoice: %v", err)
	}
	log.Printf("BillingService.GetInvoice => invoice=%+v error=%+v", invRes.GetInvoice(), invRes.GetError())
}
