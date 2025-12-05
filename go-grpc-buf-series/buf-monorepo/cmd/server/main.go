package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	billingv1 "acme/gen/billingv1"
	commonv1 "acme/gen/commonv1"
	userv1 "acme/gen/userv1"

	"google.golang.org/grpc"
)

// userServer implements user.v1.UserService.
type userServer struct {
	userv1.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	if req.GetId() == "" {
		return &userv1.GetUserResponse{
			Error: &commonv1.ErrorStatus{
				Code:    "INVALID_ARGUMENT",
				Message: "id is required",
			},
		}, nil
	}
	return &userv1.GetUserResponse{
		User: &userv1.User{
			Id:    req.GetId(),
			Email: fmt.Sprintf("%s@example.com", req.GetId()),
		},
	}, nil
}

// billingServer implements billing.v1.BillingService.
type billingServer struct {
	billingv1.UnimplementedBillingServiceServer
}

func (s *billingServer) GetInvoice(ctx context.Context, req *billingv1.GetInvoiceRequest) (*billingv1.GetInvoiceResponse, error) {
	if req.GetId() == "" {
		return &billingv1.GetInvoiceResponse{
			Error: &commonv1.ErrorStatus{
				Code:    "INVALID_ARGUMENT",
				Message: "id is required",
			},
		}, nil
	}
	return &billingv1.GetInvoiceResponse{
		Invoice: &billingv1.Invoice{
			Id:          req.GetId(),
			UserId:      "user-123",
			AmountCents: 4999,
			Currency:    "USD",
		},
	}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// User service on :50051
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on 50051: %v", err)
		}
		grpcServer := grpc.NewServer()
		userv1.RegisterUserServiceServer(grpcServer, &userServer{})
		log.Println("UserService listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("UserService failed: %v", err)
		}
	}()

	// Billing service on :50052
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen on 50052: %v", err)
		}
		grpcServer := grpc.NewServer()
		billingv1.RegisterBillingServiceServer(grpcServer, &billingServer{})
		log.Println("BillingService listening on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("BillingService failed: %v", err)
		}
	}()

	wg.Wait()
}
