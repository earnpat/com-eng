// grpc-tutorial-topic-service/main.go
package main

import (
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

type Repository interface {
	Create(*pb.Topic) (*pb.Topic, error)
	GetAll() []*pb.Topic
}
type TopicRepository struct {
	topics []*pb.Topic
}

func (repo *TopicRepository) Create(topic *pb.Topic) (*pb.Topic, error) {
	updated := append(repo.topics, topic)
	repo.topics = updated
	return topic, nil
}
func (repo *TopicRepository) GetAll() []*pb.Topic {
	return repo.topics
}

type service struct {
	// repo Repository
	pb.UnimplementedTopicServiceServer
}

// func NewHandler(repo Repository) pb.TopicServiceServer {
// 	return &service{repo: repo}
// }

// func (s *service) CreateTopic(ctx context.Context, req *pb.Topic) (*pb.Response, error) {
// 	topic, err := s.repo.Create(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("Client call: CreateTopic")
// 	return &pb.Response{Created: true, Topic: topic}, nil
// }

// func (s *service) GetTopics(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
// 	topics := s.repo.GetAll()
// 	log.Println("Client call: GetTopics")
// 	return &pb.Response{Topics: topics}, nil
// }

func (s *service) HelloTopic(ctx context.Context, req *pb.GetRequest) (*pb.ResponseHello, error) {
	// topics := s.repo.GetAll()
	log.Println("Client call: HelloTopic")
	timestamp := time.Now().Unix()
	timeString := strconv.Itoa(int(timestamp))
	log.Println(timeString)
	return &pb.ResponseHello{Message: timeString}, nil
}

func main() {
	c := make(chan os.Signal, 1)
	// repo := &TopicRepository{}
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(math.MaxUint32))
	grpcServer := grpc.NewServer()
	// grpcHandler := NewHandler(repo)
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
