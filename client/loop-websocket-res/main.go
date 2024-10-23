package main

import (
	"client-server/helper"
	logsCollection "client-server/repository/logs"
	"context"
	"encoding/json"
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

type TodoData struct {
	Id        int64  `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"userId"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	ctx := context.Background()

	dbmg, dbmgCtx := helper.MongoConnection(os.Getenv("MONGO_URL"))
	mongoDB := *dbmg.Database("com-eng")

	logsSvc := logsCollection.NewCollection(mongoDB, "logs")

	basePath := os.Getenv("BASE_IP") + ":9003"
	logrus.Info("basePath: ", basePath)

	u := url.URL{Scheme: "ws", Host: basePath, Path: "/ws/response"}

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

	loopTimeSec := 60
	startTime := time.Now()
	endTime := startTime.Add(time.Duration(loopTimeSec) * time.Second)

	count := 0
	countSuccess := 0
	countFail := 0
	totalRequestTime := int64(0)
	totalResponseSize := 0
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

		_, msg, err := conn.ReadMessage()
		if err != nil {
			countFail++
			continue
		}
		requestDuration := time.Since(timestamp).Nanoseconds()

		responseSize := len(msg)

		var response []TodoData
		err = json.Unmarshal(msg, &response)
		if err != nil {
			logrus.Info(err)
			countFail++
			continue
		}

		if float64(requestDuration) > (maxTimeNanoSec) {
			maxTimeNanoSec = float64(requestDuration)
		}

		if float64(requestDuration) < minTimeNanoSec {
			minTimeNanoSec = float64(requestDuration)
		}

		totalRequestTime += requestDuration
		totalResponseSize += responseSize
		countSuccess++
	}

	logrus.Info("end of websocket service")
	logrus.Info("count: ", count)
	logrus.Info("countSuccess: ", countSuccess)
	logrus.Info("countFail: ", countFail)

	avgRequestTimeNanoSec := totalRequestTime / int64(count)
	millisecond := float64(avgRequestTimeNanoSec) / float64(1000000)
	logrus.Info("millisecond: ", millisecond)

	responseSize := totalResponseSize / count

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
		ResponseSize:          float64(responseSize),
	}, *options.InsertOne())

	defer func() {
		if err := dbmg.Disconnect(dbmgCtx); err != nil {
			panic(err)
		}
	}()
}
