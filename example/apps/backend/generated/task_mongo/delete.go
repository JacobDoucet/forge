package task_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(ctx context.Context, db *mongo.Database, id primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func DeleteBulk(ctx context.Context, db *mongo.Database, ids []primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return err
	}

	return nil
}
