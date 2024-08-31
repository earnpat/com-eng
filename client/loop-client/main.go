package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type ResBody struct {
	Timestamp int64 `json:"timestamp"`
}

func main() {

	startTime := time.Now()
	// endTime := startTime.Add(1 * time.Minute)
	endTime := startTime.Add(10 * time.Second)
	count := 0
	countSuccess := 0
	countFail := 0

	// timestamp := time.Now().UnixNano()

	for time.Now().Before(endTime) {
		count++
		logrus.Info("count: ", count)

		// res, resErr := http.Get("http://localhost:9000/rest/start")
		res, resErr := http.Get("http://localhost:9000/grpc/start")
		if resErr != nil {
			countFail++
			continue
		}

		defer res.Body.Close()
		resBody, resBodyErr := io.ReadAll(res.Body)
		if resBodyErr != nil {
			continue
		}

		response := ResBody{}
		unmarshalErr := json.Unmarshal(resBody, &response)
		if unmarshalErr != nil {
			continue
		}

		// logrus.Info("response: ", response.Timestamp)
		countSuccess++
	}

	logrus.Info("count: ", count)
	logrus.Info("countSuccess: ", countSuccess)
	logrus.Info("countFail: ", countFail)

}
