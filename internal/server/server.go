package server

import (
	"context"
	"fmt"

	pb "github.com/harshb910-mercari/proto-test/generated/api"
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
	message := fmt.Sprintf("Hello, %s!", req.Name)
	fmt.Println(message)
	return &pb.TestResponse{
		Message: message,
	}, nil
}
