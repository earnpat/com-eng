package helper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schema struct {
	ID             primitive.ObjectID `json:"id"                bson:"_id"`
	CreatedTime    time.Time          `json:"createdTime"       bson:"createdTime"`
	StartTime      int64              `json:"startTime"         bson:"startTime"`
	EndTime        int64              `json:"endTime"           bson:"endTime"`
	Nanosecond     int64              `json:"nanosecond"        bson:"nanosecond"`
	MilliSec       float64            `json:"millisecond"       bson:"millisecond"`
	MilliSecOneWay float64            `json:"millisecondOneWay" bson:"millisecondOneWay"`
	RefKey         string             `json:"refKey"            bson:"refKey"`
}
