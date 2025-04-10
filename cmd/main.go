package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/harshb910-mercari/proto-test/generated/api"
	"github.com/harshb910-mercari/proto-test/internal/server"
	"google.golang.org/grpc"
)

// startGRPCServer starts the gRPC server
func startGRPCServer(grpcServer pb.TestServiceServer, listener net.Listener) {
	server := grpc.NewServer()
	pb.RegisterTestServiceServer(server, grpcServer)
	log.Println("Starting gRPC server on :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

// startHTTPServer starts the HTTP server using gRPC-Gateway
func startHTTPServer(ctx context.Context, address string) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, ":50051", opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	log.Printf("Starting HTTP server on %s", address)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

func main() {
	ctx := context.Background()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := server.NewTestServiceServer()

	// Run gRPC server in a goroutine
	go startGRPCServer(grpcServer, listener)

	// Run HTTP server
	startHTTPServer(ctx, ":8080")
}
