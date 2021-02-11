package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fleanza74/test-proto/gen/edgekey"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEdgeKeyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateKey(ctx, &pb.EdgeKeyRequest{
		AuthToken: "qweeerr",
		Timestamp: 12345,
		Md5:       "123adfsdffd",
		Ip:        "127.0.0.1",
		UserAgent: "Postman"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetKey())
}
