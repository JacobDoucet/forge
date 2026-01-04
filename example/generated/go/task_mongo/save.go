package task_mongo

import (
	"context"
	"github.com/JacobDoucet/forge/example/generated/go/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, db *mongo.Database, m task.MongoRecord) (primitive.ObjectID, error) {
	coll := db.Collection(CollectionName)
	result, err := coll.InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func CreateBulk(ctx context.Context, db *mongo.Database, m []task.MongoRecord) ([]primitive.ObjectID, error) {
	coll := db.Collection(CollectionName)
	var docs []interface{}
	for _, v := range m {
		docs = append(docs, v)
	}
	result, err := coll.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID))
	}
	return ids, nil
}

func Update(ctx context.Context, db *mongo.Database, m task.MongoRecord, where task.MongoWhereClause) error {
	filter, err := where.GetLookupQuery()
	if err != nil {
		return err
	}
	m.Id = nil

	coll := db.Collection(CollectionName)
	_, err = coll.UpdateOne(ctx, filter, bson.M{"$set": m})
	return err
}

func UpdateMany(ctx context.Context, db *mongo.Database, m task.MongoRecord, ids []primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	filter := bson.M{"_id": bson.M{"$in": ids}}
	m.Id = nil
	_, err := coll.UpdateMany(ctx, filter, bson.M{"$set": m})
	return err
}
