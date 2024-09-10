package main

import (
	"client-server/helper"
	logsCollection "client-server/repository/v3-logs"
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResBody struct {
	Timestamp int64 `json:"timestamp"`
}

type TodoData struct {
	Id        int64  `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"userId"`
}

func main() {
	ctx := context.Background()

	// dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27017/") // cloud
	dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/") // local
	mongoDB := *dbmg.Database("com-eng-v3")

	logsSvc := logsCollection.NewCollection(mongoDB)

	url := "https://159.223.36.152:9001"

	// check connection
	_, resErr := http.Get(url)
	if resErr != nil {
		return
	}

	// loopTimeSecs := []int{3}
	loopTimeSecs := helper.GetLoopTime()
	for _, loopTimeSec := range loopTimeSecs {
		startTime := time.Now()
		endTime := startTime.Add(time.Duration(loopTimeSec) * time.Second)

		count := 0
		countSuccess := 0
		countFail := 0
		totalRequestTime := int64(0)
		minTimeNanoSec := float64(1000000000)
		maxTimeNanoSec := float64(0)

		logrus.Info("loopTimeSec: ", loopTimeSec)
		logrus.Info("start")
		for time.Now().Before(endTime) {
			count++

			timestamp := time.Now()
			// res, resErr := http.Get(url)
			_, resErr := http.Get(url)
			if resErr != nil {
				logrus.Info("resErr: ", resErr)
				countFail++
				continue
			}
			requestDuration := time.Since(timestamp).Nanoseconds()

			// defer res.Body.Close()

			// var resBody struct {
			// 	Timestamp int64      `json:"timestamp"`
			// 	Todo      []TodoData `json:"todo"`
			// }
			// resBodyErr := json.NewDecoder(res.Body).Decode(&resBody)
			// if resBodyErr != nil {
			// 	countFail++
			// }

			// timestampStart := resBody.Timestamp
			// // timestampStart := int64(0)
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
			Type:                  logsCollection.REST,
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

		time.Sleep(10 * time.Second)
	}

	defer func() {
		if err := dbmg.Disconnect(dbmgCtx); err != nil {
			panic(err)
		}
	}()
}
