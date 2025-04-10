package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/harshb910-mercari/proto-test/generated/api"
	"github.com/harshb910-mercari/proto-test/internal/server"
	"google.golang.org/grpc"
)

// startGRPCServer starts the gRPC server
func runGRPCServer(ctx context.Context, grpcServer *grpc.Server, listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		<-ctx.Done()
		log.Println("Initiating graceful shutdown for gRPC server...")
		grpcServer.GracefulStop()
	}()

	log.Println("gRPC server running at", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("gRPC server stopped unexpectedly: %v", err)
	}
	log.Println("gRPC Server stopped.")
}

// startHTTPServer starts the HTTP server
func runHTTPServer(ctx context.Context, httpServer *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		<-ctx.Done()
		log.Println("Initiating graceful shutdown for HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Println("HTTP server graceful shutdown failed:", err)
		}
	}()

	log.Println("HTTP server running at", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("HTTP server stopped unexpectedly: %v", err)
	}
	log.Println("HTTP Server stopped.")
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	// Set up gRPC Server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, server.NewTestServiceServer())
	wg.Add(1)
	go runGRPCServer(ctx, grpcServer, listener, &wg)

	// Set up HTTP Server (gRPC-gateway)
	mux := runtime.NewServeMux()
	if err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, ":50051", []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	wg.Add(1)
	go runHTTPServer(ctx, httpServer, &wg)

	// Wait for SIGINT (Ctrl+C) or SIGTERM
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// Give servers some time for graceful termination
	stop()
	wg.Wait()
	log.Println("Graceful shutdown")
}
