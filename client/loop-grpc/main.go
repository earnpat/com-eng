package main

import (
	"client-server/helper"
	logsCollection "client-server/repository/v3-logs"
	pb "client-server/services"
	"context"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27017/") // cloud
	dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/") // local
	mongoDB := *dbmg.Database("com-eng-v3")

	logsSvc := logsCollection.NewCollection(mongoDB)

	conn, err := grpc.NewClient(
		// "localhost:9002",
		"159.223.36.152:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)

	// check connection
	_, resErr := client.GetTopics(ctx, &pb.GetRequest{})
	if resErr != nil {
		logrus.Info("resErr: ", resErr)
		return
	}

	loopTimeSec := 5
	startTime := time.Now()
	endTime := startTime.Add(time.Duration(loopTimeSec) * time.Second)

	count := 0
	countSuccess := 0
	countFail := 0
	totalRequestTime := int64(0)
	minTimeNanoSec := float64(1000000)
	maxTimeNanoSec := float64(0)

	logrus.Info("loopTimeSec: ", loopTimeSec)
	logrus.Info("start")
	for time.Now().Before(endTime) {
		count++

		timestamp := time.Now()
		// res, resErr := client.GetTopics(ctx, &pb.GetRequest{})
		_, resErr := client.GetTopics(ctx, &pb.GetRequest{})
		if resErr != nil {
			countFail++
		}
		requestDuration := time.Since(timestamp).Nanoseconds()

		// timestampStart := res.Timestamp
		// timestampEnd := time.Now().UnixNano()
		// nanosecond := timestampEnd - timestampStart

		if float64(requestDuration) > (maxTimeNanoSec) {
			maxTimeNanoSec = float64(requestDuration)
		}

		if float64(requestDuration) < minTimeNanoSec {
			minTimeNanoSec = float64(requestDuration)
		}

		// totalRequestTime += nanosecond
		totalRequestTime += requestDuration
		countSuccess++
	}

	logrus.Info("end of grpc service")
	logrus.Info("count: ", count)
	logrus.Info("countSuccess: ", countSuccess)
	logrus.Info("countFail: ", countFail)

	avgRequestTimeNanoSec := totalRequestTime / int64(count)
	millisecond := float64(avgRequestTimeNanoSec) / float64(1000000)
	logrus.Info("millisecond: ", millisecond)

	logsSvc.InsertOne(ctx, logsCollection.LogsSchema{
		ID:                    primitive.NewObjectID(),
		CreatedTime:           time.Now(),
		Type:                  logsCollection.GRPC,
		Connection:            logsCollection.LOCAL,
		LoopTimeSec:           int64(loopTimeSec),
		Count:                 int64(count),
		CountSuccess:          int64(countSuccess),
		CountFail:             int64(countFail),
		MinTimeNanoSec:        minTimeNanoSec,
		MaxTimeNanoSec:        maxTimeNanoSec,
		AvgRequestTimeNanoSec: avgRequestTimeNanoSec,
		AvgMilliSec:           millisecond,
		AvgMilliSecOneWay:     millisecond / float64(2),
	}, *options.InsertOne())

	defer func() {
		if err := dbmg.Disconnect(dbmgCtx); err != nil {
			panic(err)
		}
	}()
}
