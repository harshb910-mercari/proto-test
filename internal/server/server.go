package server

import (
	"context"
	"fmt"

	pb "github.com/harshb910-mercari/proto-test/generated/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements pb.TestServiceServer
type Server struct {
	pb.UnimplementedTestServiceServer
}

// NewTestServiceServer initializes and returns a new Test Service server
func NewTestServiceServer() pb.TestServiceServer {
	return &Server{}
}

// SayHello handles api request and returns a response
func (s *Server) SayHello(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	if err := req.Validate(); err != nil {
		fmt.Println("failed due to validation error", err)
		return nil, status.Errorf(codes.InvalidArgument, "Validation failed: %v", err)
	}
	fmt.Println("request is:", req)

	name := "Anonymous"
	if req.Name != nil && req.Name.Value != "" {
		name = req.Name.Value
	}

	// continue logic if validation passes
	message := fmt.Sprintf("Hello, %s!", name)
	fmt.Println(message)
	return &pb.TestResponse{
		Message: message,
	}, nil
}
