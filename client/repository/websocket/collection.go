package websocketCollection

import (
	"client-server/helper"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	helper.IMongoSchemaHelper[helper.Schema]
}

type collectionService struct {
	helper.MongoSchemaHelperService[helper.Schema]
	collection *mongo.Collection
}

func NewCollection(client mongo.Database) ICollection {
	mgCollection := client.Collection("websocket")
	return &collectionService{
		MongoSchemaHelperService: helper.MongoSchemaHelperService[helper.Schema]{
			Collection: mgCollection,
		},
		collection: mgCollection,
	}
}
