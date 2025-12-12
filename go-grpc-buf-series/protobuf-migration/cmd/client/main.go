package main

import (
	"context"
	"log"
	"time"

	billingv1 "acme/gen/billingv1"
	userv1 "acme/gen/userv1"
	userv2 "acme/gen/userv2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Call UserService (v1 and v2) on the same server
	userConn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial user: %v", err)
	}
	defer userConn.Close()

	// v1 call
	userClientV1 := userv1.NewUserServiceClient(userConn)
	userResV1, err := userClientV1.GetUser(ctx, &userv1.GetUserRequest{Id: "user-123"})
	if err != nil {
		log.Fatalf("GetUser v1: %v", err)
	}
	log.Printf("UserService(v1).GetUser => user=%+v error=%+v", userResV1.GetUser(), userResV1.GetError())

	// v2 call
	userClientV2 := userv2.NewUserServiceClient(userConn)
	userResV2, err := userClientV2.GetUser(ctx, &userv2.GetUserRequest{Id: "user-123"})
	if err != nil {
		log.Fatalf("GetUser v2: %v", err)
	}
	log.Printf("UserService(v2).GetUser => user=%+v", userResV2.GetUser())

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
