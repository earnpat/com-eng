package main

import (
	"encoding/json"
	"log"

	pb "github.com/biskitsx/go-fiber-api/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)

	var topic *pb.Topic
	data := map[string]interface{}{
		"title":       "Used computer for sale",
		"description": "Lenevo ThinkPad T440",
		"price":       300,
		"category_id": "category001",
	}
	jsonStr, _ := json.Marshal(data)
	json.Unmarshal(jsonStr, &topic)

	response, err := client.HelloTopic(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list topics: %v", err)
	}

	log.Println("response: ", response.Message)
}
