package logsCollection

import (
	"client-server/helper"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	helper.IMongoSchemaHelper[LogsSchema]
}

type collectionService struct {
	helper.MongoSchemaHelperService[LogsSchema]
	collection *mongo.Collection
}

func NewCollection(client mongo.Database, connectionName string) ICollection {
	mgCollection := client.Collection(connectionName)
	return &collectionService{
		MongoSchemaHelperService: helper.MongoSchemaHelperService[LogsSchema]{
			Collection: mgCollection,
		},
		collection: mgCollection,
	}
}
