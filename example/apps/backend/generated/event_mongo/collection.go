package event_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "events"

func CreateIndexes(ctx context.Context, db *mongo.Database) error {
	if err := createCreatedAtIndex(ctx, db); err != nil {
		return err
	}
	if err := createSubjectIndex(ctx, db); err != nil {
		return err
	}
	if err := createSubjectTypeIndex(ctx, db); err != nil {
		return err
	}
	if err := createTypeIndex(ctx, db); err != nil {
		return err
	}
	if err := createUpdatedAtIndex(ctx, db); err != nil {
		return err
	}
	return nil
}

func createCreatedAtIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "created.at", Value: -1},
		},
		Options: options.Index().SetName("createdAt"),
	})
	return err
}

func createSubjectIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "subjects.subjectId", Value: 1},
			{Key: "created.at", Value: 1},
		},
		Options: options.Index().SetName("subject"),
	})
	return err
}

func createSubjectTypeIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "subjects.subjectId", Value: 1},
			{Key: "type", Value: 1},
			{Key: "created.at", Value: 1},
		},
		Options: options.Index().SetName("subject_type"),
	})
	return err
}

func createTypeIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "type", Value: 1},
			{Key: "created.at", Value: 1},
		},
		Options: options.Index().SetName("type"),
	})
	return err
}

func createUpdatedAtIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "updated.at", Value: -1},
		},
		Options: options.Index().SetName("updatedAt"),
	})
	return err
}
