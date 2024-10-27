package main

import (
	"client-server/helper"
	logsCollection "client-server/repository/logs"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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

	logsSvc := logsCollection.NewCollection(mongoDB, "logs-count")

	url := "http://" + os.Getenv("BASE_IP") + ":9001/response"
	logrus.Info("url: ", url)

	// check connection
	_, resErr := http.Get(url)
	if resErr != nil {
		logrus.Info("resErr: ", resErr.Error())
		return
	}

	loopCount := 3000

	count := 0
	countSuccess := 0
	countFail := 0
	totalRequestTime := int64(0)
	totalResponseSize := 0
	minTimeNanoSec := float64(1000000000)
	maxTimeNanoSec := float64(0)

	logrus.Info("loopCount: ", loopCount)
	logrus.Info("start")
	startTime := time.Now()
	for count < loopCount {
		count++

		timestamp := time.Now()
		res, resErr := http.Get(url)
		if resErr != nil {
			logrus.Info("resErr: ", resErr)
			countFail++
			continue
		}
		requestDuration := time.Since(timestamp).Nanoseconds()

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			countFail++
			continue
		}

		responseSize := len(body)

		var response struct {
			Todo []TodoData `json:"todo"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
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
	endTime := time.Since(startTime).Milliseconds()

	logrus.Info("end of rest service")
	logrus.Info("endTime: ", endTime)
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
		Type:                  logsCollection.REST,
		Connection:            logsCollection.LOCAL,
		LoopTimeMilliSec:      endTime,
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
