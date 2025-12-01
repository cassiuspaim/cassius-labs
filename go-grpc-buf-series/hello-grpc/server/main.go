package main

import (
	"context"
	"log"
	"net"

	hellov1 "github.com/cassius-labs/go-grpc-buf-series/hello-grpc/gen/hellov1"
	"google.golang.org/grpc"
)

type greeterServer struct {
	hellov1.UnimplementedGreeterServiceServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *hellov1.SayHelloRequest) (*hellov1.SayHelloResponse, error) {
	msg := "Hello, " + req.GetName() + "!"
	return &hellov1.SayHelloResponse{Message: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	hellov1.RegisterGreeterServiceServer(grpcServer, &greeterServer{})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}