package main

import (
	"client-server/helper"
	logsCollection "client-server/repository/v3-logs"
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	ctx := context.Background()

	dbmg, dbmgCtx := helper.MongoConnection(os.Getenv("MONGO_URL"))
	mongoDB := *dbmg.Database("com-eng-v3")

	logsSvc := logsCollection.NewCollection(mongoDB)

	basePath := os.Getenv("BASE_IP") + ":9003"
	logrus.Info("basePath: ", basePath)

	u := url.URL{Scheme: "ws", Host: basePath, Path: "/ws/test"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	// check connection
	err = conn.WriteMessage(websocket.TextMessage, nil)
	if err != nil {
		return
	}

	_, _, err = conn.ReadMessage()
	if err != nil {
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
		err = conn.WriteMessage(websocket.TextMessage, nil)
		if err != nil {
			countFail++
			continue
		}

		_, _, err := conn.ReadMessage()
		if err != nil {
			countFail++
			continue
		}
		requestDuration := time.Since(timestamp).Nanoseconds()

		if float64(requestDuration) > (maxTimeNanoSec) {
			maxTimeNanoSec = float64(requestDuration)
		}

		if float64(requestDuration) < minTimeNanoSec {
			minTimeNanoSec = float64(requestDuration)
		}

		totalRequestTime += requestDuration
		countSuccess++
	}

	logrus.Info("end of rest service")
	logrus.Info("count: ", count)
	logrus.Info("countSuccess: ", countSuccess)
	logrus.Info("countFail: ", countFail)

	avgRequestTimeNanoSec := totalRequestTime / int64(count)
	millisecond := float64(avgRequestTimeNanoSec) / float64(1000000)
	logrus.Info("millisecond: ", millisecond)

	logsSvc.InsertOne(ctx, logsCollection.LogsSchema{
		ID:                    primitive.NewObjectID(),
		CreatedTime:           time.Now(),
		Type:                  logsCollection.WS,
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
