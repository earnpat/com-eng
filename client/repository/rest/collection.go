package restCollection

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

func NewCollection(client mongo.Database, connectionName string) ICollection {
	mgCollection := client.Collection("rest")
	return &collectionService{
		MongoSchemaHelperService: helper.MongoSchemaHelperService[helper.Schema]{
			Collection: mgCollection,
		},
		collection: mgCollection,
	}
}
