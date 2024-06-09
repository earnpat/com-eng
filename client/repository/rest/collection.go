package restCollection

import (
	"client-server/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	helper.IMongoSchemaHelper[RestSchema]
}

type collectionService struct {
	helper.MongoSchemaHelperService[RestSchema]
	collection *mongo.Collection
}

type RestSchema struct {
	ID          primitive.ObjectID `json:"id"          bson:"_id"`
	CreatedTime time.Time          `json:"createdTime" bson:"createdTime"`
	Nanosecond  int64              `json:"nanosecond"  bson:"nanosecond"`
	Millisecond float64            `json:"millisecond" bson:"millisecond"`
}

func NewCollection(client mongo.Database) ICollection {
	mgCollection := client.Collection("rest")
	return &collectionService{
		MongoSchemaHelperService: helper.MongoSchemaHelperService[RestSchema]{
			Collection: mgCollection,
		},
		collection: mgCollection,
	}
}
