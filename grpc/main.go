package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/biskitsx/go-fiber-api/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	pb.UnimplementedTopicServiceServer
}

func (s *service) GetTopicResponse(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	jsonFile, err := os.Open("../todo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)

	var todoData []*pb.TodoResponse
	err = json.Unmarshal(jsonData, &todoData)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Ok:   true,
		Todo: todoData,
	}, nil
}

func (s *service) GetTopic(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{
		Ok: true,
	}, nil
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
