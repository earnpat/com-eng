package main

import (
	"client-server/helper"
	grpcCollection "client-server/repository/grpc"
	restCollection "client-server/repository/rest"
	websocketCollection "client-server/repository/websocket"
	pb "client-server/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("client start")

	dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27017/") // cloud
	// dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/") // local
	mongoDB := *dbmg.Database("com-eng")

	restSvc := restCollection.NewCollection(mongoDB)
	grpcSvc := grpcCollection.NewCollection(mongoDB)
	websocketSvc := websocketCollection.NewCollection(mongoDB)

	conn, err := grpc.NewClient(
		// "178.128.88.107:9002",
		"localhost:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)

	app := fiber.New()

	app.Get("/rest/:refKey", func(c *fiber.Ctx) error {
		ctx := c.Context()
		client := &http.Client{}
		// req, reqErr := http.NewRequest("GET", "http://178.128.88.107:9001", nil)
		req, reqErr := http.NewRequest("GET", "http://localhost:9001", nil)
		if reqErr != nil {
			return reqErr
		}

		q := req.URL.Query()
		timestamp := time.Now().UnixNano()
		q.Add("timestamp", strconv.Itoa(int(timestamp)))
		req.URL.RawQuery = q.Encode()

		resHttp, resHttpErr := client.Do(req)
		if resHttpErr != nil {
			return resHttpErr
		}

		var resBody struct {
			Timestamp int64 `json:"timestamp"`
		}
		resBodyErr := json.NewDecoder(resHttp.Body).Decode(&resBody)
		if resBodyErr != nil {
			return resBodyErr
		}

		timestampStart := resBody.Timestamp
		timestampEnd := time.Now().UnixNano()
		nanosecond := timestampEnd - timestampStart
		millisecond := float64(timestampEnd-timestampStart) / float64(1000000)

		refKey := c.Params("refKey")
		if refKey != "start" {
			restSvc.InsertOne(ctx, helper.Schema{
				ID:                primitive.NewObjectID(),
				CreatedTime:       time.Now(),
				StartTime:         timestampStart,
				EndTime:           timestampEnd,
				Nanosecond:        nanosecond,
				Millisecond:       millisecond,
				MillisecondOneWay: millisecond / float64(2),
				RefKey:            refKey,
			}, *options.InsertOne())
		}

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"millisecond": millisecond,
			"nanosecond":  nanosecond,
		})
	})

	app.Get("/grpc/:refKey", func(c *fiber.Ctx) error {
		ctx := c.Context()
		timestamp := time.Now().UnixNano()
		response, err := client.GetTopics(
			context.Background(),
			&pb.GetRequest{Timestamp: timestamp},
		)
		if err != nil {
			log.Fatalf("Could not list topics: %v", err)
		}

		timestampStart := response.Timestamp
		timestampEnd := time.Now().UnixNano()
		nanosecond := timestampEnd - timestampStart
		millisecond := float64(timestampEnd-timestampStart) / float64(1000000)

		refKey := c.Params("refKey")
		if refKey != "start" {
			grpcSvc.InsertOne(ctx, helper.Schema{
				ID:                primitive.NewObjectID(),
				CreatedTime:       time.Now(),
				StartTime:         timestampStart,
				EndTime:           timestampEnd,
				Nanosecond:        nanosecond,
				Millisecond:       millisecond,
				MillisecondOneWay: millisecond / float64(2),
				RefKey:            refKey,
			}, *options.InsertOne())
		}

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"millisecond": millisecond,
			"nanosecond":  nanosecond,
		})
	})

	app.Get("/websocket/:refKey", func(c *fiber.Ctx) error {
		ctx := c.Context()
		// u := url.URL{Scheme: "ws", Host: "178.128.88.107:9003", Path: "/ws/test"}
		u := url.URL{Scheme: "ws", Host: "localhost:9003", Path: "/ws/test"}

		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("dial:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to WebSocket")
		}
		defer conn.Close()

		timestamp := time.Now().UnixNano()
		err = conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(int(timestamp))))
		if err != nil {
			log.Println("write:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to send message to WebSocket")
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to read message from WebSocket")
		}

		messageStr := string(message[:])
		timestampStart, _ := strconv.ParseInt(messageStr, 10, 64)
		timestampEnd := time.Now().UnixNano()
		nanosecond := timestampEnd - timestampStart
		millisecond := float64(timestampEnd-timestampStart) / float64(1000000)

		refKey := c.Params("refKey")
		if refKey != "start" {
			websocketSvc.InsertOne(ctx, helper.Schema{
				ID:                primitive.NewObjectID(),
				CreatedTime:       time.Now(),
				StartTime:         timestampStart,
				EndTime:           timestampEnd,
				Nanosecond:        nanosecond,
				Millisecond:       millisecond,
				MillisecondOneWay: millisecond / float64(2),
				RefKey:            refKey,
			}, *options.InsertOne())
		}

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"millisecond": millisecond,
			"nanosecond":  nanosecond,
		})
	})

	app.Listen(":9000")

	defer func() {
		if err := dbmg.Disconnect(dbmgCtx); err != nil {
			panic(err)
		}
	}()
}
