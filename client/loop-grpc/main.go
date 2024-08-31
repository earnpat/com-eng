package main

import (
	pb "client-server/services"
	"context"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ResBody struct {
	Timestamp int64 `json:"timestamp"`
}

func main() {
	ctx := context.Background()

	// // dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27017/") // cloud
	// dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/") // local
	// mongoDB := *dbmg.Database("com-eng-v3")

	// logsSvc := logsCollection.NewCollection(mongoDB)

	conn, err := grpc.NewClient(
		"localhost:9002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)

	// check connection
	_, resErr := client.GetTopics(ctx, &pb.GetRequest{Timestamp: 0})
	if resErr != nil {
		logrus.Info("resErr: ", resErr)
		return
	}

	loopTimeSecs := []int{300}
	// loopTimeSecs := helper.GetLoopTime()
	for _, loopTimeSec := range loopTimeSecs {
		startTime := time.Now()
		endTime := startTime.Add(time.Duration(loopTimeSec) * time.Second)

		count := 0
		countSuccess := 0
		countFail := 0
		totalRequestTime := int64(0)
		minTimeNanoSec := float64(1000000)
		maxTimeNanoSec := float64(0)

		logrus.Info("start")
		for time.Now().Before(endTime) {
			count++

			timestamp := time.Now().UnixNano()
			res, resErr := client.GetTopics(ctx, &pb.GetRequest{Timestamp: timestamp})
			if resErr != nil {
				countFail++
			}

			timestampStart := res.Timestamp
			timestampEnd := time.Now().UnixNano()
			nanosecond := timestampEnd - timestampStart

			if float64(nanosecond) > (maxTimeNanoSec) {
				maxTimeNanoSec = float64(nanosecond)
			}

			if float64(nanosecond) < minTimeNanoSec {
				minTimeNanoSec = float64(nanosecond)
			}

			totalRequestTime += nanosecond
			countSuccess++
		}

		logrus.Info("end of grpc service")
		logrus.Info("count: ", count)
		logrus.Info("countSuccess: ", countSuccess)
		logrus.Info("countFail: ", countFail)

		avgRequestTimeNanoSec := totalRequestTime / int64(count)
		millisecond := float64(avgRequestTimeNanoSec) / float64(1000000)
		logrus.Info("millisecond: ", millisecond)

		// logsSvc.InsertOne(ctx, logsCollection.LogsSchema{
		// 	ID:                    primitive.NewObjectID(),
		// 	CreatedTime:           time.Now(),
		// 	Type:                  logsCollection.GRPC,
		// 	LoopTimeSec:           int64(loopTimeSec),
		// 	Count:                 int64(count),
		// 	CountSuccess:          int64(countSuccess),
		// 	CountFail:             int64(countFail),
		// 	MinTimeNanoSec:        minTimeNanoSec,
		// 	MaxTimeNanoSec:        maxTimeNanoSec,
		// 	AvgRequestTimeNanoSec: avgRequestTimeNanoSec,
		// 	AvgMilliSec:           millisecond,
		// 	AvgMilliSecOneWay:     millisecond / float64(2),
		// }, *options.InsertOne())

		// time.Sleep(10 * time.Second)
	}

	// defer func() {
	// 	if err := dbmg.Disconnect(dbmgCtx); err != nil {
	// 		panic(err)
	// 	}
	// }()
}
