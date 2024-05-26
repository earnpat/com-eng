// grpc-tutorial-topic-service/client/client.go
// สำหรับการทดลองนี้ ให้แยกใส่โฟลเดอร์ client เพื่อไม่ให้มีปัญหาเรื่อง package
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
	// defaultFilename = "topic.json"
)

// func parseFile(file string) (*pb.Topic, error) {
// 	var topic *pb.Topic
// 	// data, err := ioutil.ReadFile(file)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	data := map[string]interface{}{
// 		"title":       "Used computer for sale",
// 		"description": "Lenevo ThinkPad T440",
// 		"price":       300,
// 		"category_id": "category001",
// 	}
// 	jsonStr, _ := json.Marshal(data)
// 	json.Unmarshal(jsonStr, &topic)
// 	return topic, nil
// }

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)
	// file := defaultFilename
	// if len(os.Args) > 1 {
	// 	file = os.Args[1]
	// }

	// topic, err := parseFile(file)
	// if err != nil {
	// 	log.Fatalf("Could not parse file: %v", err)
	// }

	var topic *pb.Topic
	data := map[string]interface{}{
		"title":       "Used computer for sale",
		"description": "Lenevo ThinkPad T440",
		"price":       300,
		"category_id": "category001",
	}
	jsonStr, _ := json.Marshal(data)
	json.Unmarshal(jsonStr, &topic)

	// r, err := client.CreateTopic(context.Background(), topic)
	// if err != nil {
	// 	log.Fatalf("Could not greet: %v", err)
	// }

	// log.Printf("Created: %t", r.Created)
	// getAll, err := client.GetTopics(context.Background(), &pb.GetRequest{})
	// if err != nil {
	// 	log.Fatalf("Could not list topics: %v", err)
	// }

	// for _, topic := range getAll.Topics {
	// 	log.Println("topic: ", topic)
	// }

	hello, err := client.HelloTopic(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list topics: %v", err)
	}

	log.Println("hello: ", hello.Message)

}
