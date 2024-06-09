package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/biskitsx/go-fiber-api/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	pb.UnimplementedTopicServiceServer
}

func (s *service) GetTopics(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{Timestamp: req.Timestamp}, nil
}

func main() {
	fmt.Println("grpc start")
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTopicServiceServer(grpcServer, &service{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
