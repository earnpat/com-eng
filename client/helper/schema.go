package helper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schema struct {
	ID                primitive.ObjectID `json:"id"                bson:"_id"`
	CreatedTime       time.Time          `json:"createdTime"       bson:"createdTime"`
	Nanosecond        int64              `json:"nanosecond"        bson:"nanosecond"`
	Millisecond       float64            `json:"millisecond"       bson:"millisecond"`
	MillisecondOneWay float64            `json:"millisecondOneWay" bson:"millisecondOneWay"`
}
