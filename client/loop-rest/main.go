package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type ResBody struct {
	Timestamp int64 `json:"timestamp"`
}

func main() {
	// ctx := context.Background()

	// // dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27017/") // cloud
	// dbmg, dbmgCtx := helper.MongoConnection("mongodb://localhost:27018/") // local
	// mongoDB := *dbmg.Database("com-eng-v3")

	// logsSvc := logsCollection.NewCollection(mongoDB)

	// check connection
	_, resErr := http.Get("http://localhost:9001?timestamp=" + strconv.Itoa(int(0)))
	if resErr != nil {
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
		// minTimeNanoSec := float64(1000000)
		// maxTimeNanoSec := float64(0)

		logrus.Info("start")
		for time.Now().Before(endTime) {
			count++

			timestamp := time.Now().UnixNano()
			res, resErr := http.Get("http://localhost:9001?timestamp=" + strconv.Itoa(int(timestamp)))
			// _, resErr := http.Get("http://localhost:9001?timestamp=" + strconv.Itoa(int(timestamp)))
			if resErr != nil {
				countFail++
				continue
			}

			defer res.Body.Close()

			var resBody struct {
				Timestamp int64 `json:"timestamp"`
			}
			resBodyErr := json.NewDecoder(res.Body).Decode(&resBody)
			if resBodyErr != nil {
				countFail++
			}

			timestampStart := resBody.Timestamp
			timestampEnd := time.Now().UnixNano()
			nanosecond := timestampEnd - timestampStart

			// if float64(nanosecond) > (maxTimeNanoSec) {
			// 	maxTimeNanoSec = float64(nanosecond)
			// }

			// if float64(nanosecond) < minTimeNanoSec {
			// 	minTimeNanoSec = float64(nanosecond)
			// }

			totalRequestTime += nanosecond
			countSuccess++
		}

		logrus.Info("end of rest service")
		logrus.Info("count: ", count)
		logrus.Info("countSuccess: ", countSuccess)
		logrus.Info("countFail: ", countFail)

		avgRequestTimeNanoSec := totalRequestTime / int64(count)
		millisecond := float64(avgRequestTimeNanoSec) / float64(1000000)
		logrus.Info("millisecond: ", millisecond)

		// logsSvc.InsertOne(ctx, logsCollection.LogsSchema{
		// 	ID:                    primitive.NewObjectID(),
		// 	CreatedTime:           time.Now(),
		// 	Type:                  logsCollection.REST,
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
