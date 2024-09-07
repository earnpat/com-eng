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
	Todo []*pb.TodoResponse `json:"todo"`
}

func (s *service) GetTopics(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{
		Timestamp: req.Timestamp,
		Todo:      s.Todo,
	}, nil
}

func main() {
	fmt.Println("grpc start")

	////////// mock data //////////
	jsonFile, err := os.Open("../todo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)

	var todoData []*pb.TodoResponse
	err = json.Unmarshal(jsonData, &todoData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//////////////////////////////

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTopicServiceServer(grpcServer, &service{Todo: todoData})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
