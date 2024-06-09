package grpcCollection

import (
	"client-server/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	helper.IMongoSchemaHelper[GrpcSchema]
}

type collectionService struct {
	helper.MongoSchemaHelperService[GrpcSchema]
	collection *mongo.Collection
}

type GrpcSchema struct {
	ID          primitive.ObjectID `json:"id"          bson:"_id"`
	CreatedTime time.Time          `json:"createdTime" bson:"createdTime"`
	Nanosecond  int64              `json:"nanosecond"  bson:"nanosecond"`
	Millisecond float64            `json:"millisecond" bson:"millisecond"`
}

func NewCollection(client mongo.Database) ICollection {
	mgCollection := client.Collection("rest")
	return &collectionService{
		MongoSchemaHelperService: helper.MongoSchemaHelperService[GrpcSchema]{
			Collection: mgCollection,
		},
		collection: mgCollection,
	}
}
