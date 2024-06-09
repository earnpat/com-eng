package helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoSchemaHelper[T any] interface {
	FindOne(ctx context.Context, filter bson.M, opts options.FindOneOptions) (T, error)
	Find(ctx context.Context, filter bson.M, opts options.FindOptions) ([]T, error)
	InsertOne(
		ctx context.Context,
		document T,
		opts options.InsertOneOptions,
	) (primitive.ObjectID, error)
	InsertMany(ctx context.Context, document []T, opts options.BulkWriteOptions) (int64, error)
	UpdateOne(ctx context.Context, docID primitive.ObjectID, updateSet bson.M) error
	UpdateMany(ctx context.Context, filter bson.M, updateSet bson.M) (int64, int64, error)
	DeleteOne(ctx context.Context, docID primitive.ObjectID) error
	DeleteMany(ctx context.Context, filter bson.M) (int64, error)
	Count(ctx context.Context, filter bson.M) (int64, error)
}

type MongoSchemaHelperService[T any] struct {
	Collection *mongo.Collection
}

func (s *MongoSchemaHelperService[T]) FindOne(
	ctx context.Context,
	filter bson.M,
	opts options.FindOneOptions,
) (T, error) {
	var result T
	if err := s.Collection.FindOne(ctx, filter, &opts).Decode(&result); err != nil {
		return result, fmt.Errorf("%v : %v", s.Collection.Name(), err.Error())
	}
	return result, nil
}

func (s *MongoSchemaHelperService[T]) Find(
	ctx context.Context,
	filter bson.M,
	opts options.FindOptions,
) ([]T, error) {
	var result []T
	cur, curErr := s.Collection.Find(ctx, filter, &opts)
	if curErr != nil {
		return nil, curErr
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *MongoSchemaHelperService[T]) InsertOne(
	ctx context.Context,
	document T,
	opts options.InsertOneOptions,
) (primitive.ObjectID, error) {
	insertResult, insertResultErr := s.Collection.InsertOne(ctx, document, &opts)
	if insertResultErr != nil {
		return primitive.NilObjectID, insertResultErr
	}

	insertID, _ := insertResult.InsertedID.(primitive.ObjectID)
	return insertID, nil
}

func (s *MongoSchemaHelperService[T]) InsertMany(
	ctx context.Context,
	document []T,
	opts options.BulkWriteOptions,
) (int64, error) {
	models := []mongo.WriteModel{}
	for _, doc := range document {
		models = append(models, mongo.NewInsertOneModel().SetDocument(doc))
	}
	result, err := s.Collection.BulkWrite(ctx, models, &opts)
	if err != nil {
		return 0, err
	}

	return result.InsertedCount, nil
}

func (s *MongoSchemaHelperService[T]) UpdateOne(
	ctx context.Context,
	docID primitive.ObjectID,
	updateSet bson.M,
) error {
	opts := options.Update()
	opts.SetUpsert(true)

	_, updateResultErr := s.Collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{"$set": updateSet},
		opts,
	)
	if updateResultErr != nil {
		return updateResultErr
	}

	return nil
}

func (s *MongoSchemaHelperService[T]) UpdateMany(
	ctx context.Context,
	filter bson.M,
	updateSet bson.M,
) (int64, int64, error) {
	updateResult, updateResultErr := s.Collection.UpdateMany(ctx, filter, updateSet)
	if updateResultErr != nil {
		return 0, 0, updateResultErr
	}

	return updateResult.MatchedCount, updateResult.ModifiedCount, nil
}

func (s *MongoSchemaHelperService[T]) DeleteOne(
	ctx context.Context,
	docID primitive.ObjectID,
) error {
	_, err := s.Collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoSchemaHelperService[T]) DeleteMany(
	ctx context.Context,
	filter bson.M,
) (int64, error) {
	result, err := s.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (s *MongoSchemaHelperService[T]) Count(ctx context.Context, filter bson.M) (int64, error) {
	return s.Collection.CountDocuments(ctx, filter)
}
