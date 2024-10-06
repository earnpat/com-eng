package logsCollection

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogsTypeEnum string
type LogsConnectionEnum string

const (
	REST LogsTypeEnum = "REST"
	GRPC LogsTypeEnum = "GRPC"
	WS   LogsTypeEnum = "WS"

	LOCAL  LogsConnectionEnum = "LOCAL"
	DOCKER LogsConnectionEnum = "DOCKER"
)

type LogsSchema struct {
	ID                    primitive.ObjectID `json:"id"                         bson:"_id"`
	CreatedTime           time.Time          `json:"createdTime"                bson:"createdTime"`
	Type                  LogsTypeEnum       `json:"type"                       bson:"type"`
	Connection            LogsConnectionEnum `json:"connection"                 bson:"connection"`
	LoopTimeSec           int64              `json:"loopTimeSec,omitempty"      bson:"loopTimeSec,omitempty"`
	LoopTimeMilliSec      int64              `json:"loopTimeMilliSec,omitempty" bson:"loopTimeMilliSec,omitempty"`
	Count                 int64              `json:"count"                      bson:"count"`
	CountSuccess          int64              `json:"countSuccess"               bson:"countSuccess"`
	CountFail             int64              `json:"countFail"                  bson:"countFail"`
	MinTimeNanoSec        float64            `json:"minTimeNanoSec"             bson:"minTimeNanoSec"`
	MaxTimeNanoSec        float64            `json:"maxTimeNanoSec"             bson:"maxTimeNanoSec"`
	AvgRequestTimeNanoSec int64              `json:"avgRequestTimeNanoSec"      bson:"avgRequestTimeNanoSec"`
	AvgMilliSec           float64            `json:"avgMillisecond"             bson:"avgMillisecond"`
	AvgMilliSecOneWay     float64            `json:"avgMillisecondOneWay"       bson:"avgMillisecondOneWay"`
}
