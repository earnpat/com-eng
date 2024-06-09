package main

import (
	"client-server/helper"
	grpcCollection "client-server/repository/grpc"
	restCollection "client-server/repository/rest"
	pb "client-server/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("client start")

	dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/")
	mongoDB := *dbmg.Database("com-eng")

	restSvc := restCollection.NewCollection(mongoDB)
	grpcSvc := grpcCollection.NewCollection(mongoDB)

	conn, err := grpc.NewClient(
		"localhost:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)

	app := fiber.New()

	app.Get("/rest", func(c *fiber.Ctx) error {
		ctx := c.Context()
		client := &http.Client{}
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

		logrus.Info("resHttp: ", resHttp.Body)

		var resBody struct {
			Timestamp int64 `json:"timestamp"`
		}
		resBodyErr := json.NewDecoder(resHttp.Body).Decode(&resBody)
		if resBodyErr != nil {
			return resBodyErr
		}

		timestampStart := resBody.Timestamp
		timestampEnd := time.Now().UnixNano()

		logrus.Info("timestamp start: ", timestampStart)
		logrus.Info("timestamp end: ", timestampEnd)

		nanosecond := timestampEnd - timestampStart
		logrus.Info("timestamp diff nanosecond: ", nanosecond)

		millisecond := float64(timestampEnd-timestampStart) / float64(1000000)
		logrus.Info("timestamp diff millisecond: ", millisecond)

		restSvc.InsertOne(ctx, restCollection.RestSchema{
			ID:          primitive.NewObjectID(),
			CreatedTime: time.Now(),
			Nanosecond:  nanosecond,
			Millisecond: millisecond,
		}, *options.InsertOne())

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"millisecond": millisecond,
			"nanosecond":  nanosecond,
		})
	})

	app.Get("/grpc", func(c *fiber.Ctx) error {
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

		logrus.Info("timestamp start: ", timestampStart)
		logrus.Info("timestamp end: ", timestampEnd)

		nanosecond := timestampEnd - timestampStart
		logrus.Info("timestamp diff nanosecond: ", nanosecond)

		millisecond := float64(timestampEnd-timestampStart) / float64(1000000)
		logrus.Info("timestamp diff millisecond: ", millisecond)

		grpcSvc.InsertOne(ctx, grpcCollection.GrpcSchema{
			ID:          primitive.NewObjectID(),
			CreatedTime: time.Now(),
			Nanosecond:  nanosecond,
			Millisecond: millisecond,
		}, *options.InsertOne())

		return c.Status(fiber.StatusOK).JSON(bson.M{
			"millisecond": millisecond,
			"nanosecond":  nanosecond,
		})
	})

	app.Listen(":9003")

	defer func() {
		if err := dbmg.Disconnect(dbmgCtx); err != nil {
			panic(err)
		}
	}()
}
