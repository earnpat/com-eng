// grpc-tutorial-topic-service/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/biskitsx/go-fiber-api/services"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	pb.UnimplementedTopicServiceServer
}

func (s *service) HelloTopic(ctx context.Context, req *pb.GetRequest) (*pb.ResponseHello, error) {
	log.Println("Client call: HelloTopic")
	timestamp := time.Now().Unix()
	timeString := strconv.Itoa(int(timestamp))
	log.Println(timeString)
	return &pb.ResponseHello{Message: timeString}, nil
}

func main() {
	fmt.Println("api grpc")
	c := make(chan os.Signal, 1)
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTopicServiceServer(grpcServer, &service{})
	reflection.Register(grpcServer)

	go func() {
		<-c
		logrus.Info("Gracefully shutting down...")
		grpcServer.GracefulStop()
		_ = lis.Close()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
