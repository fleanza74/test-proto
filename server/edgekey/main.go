package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fleanza74/test-proto/gen/edgekey"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedEdgeKeyServiceServer
}

func (s *server) GenerateKey(ctx context.Context, in *pb.EdgeKeyRequest) (*pb.EdgeKeyResponse, error) {
	log.Printf("Received: %v", in.GetAuthToken())
	return &pb.EdgeKeyResponse{Key: "evvai"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEdgeKeyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
