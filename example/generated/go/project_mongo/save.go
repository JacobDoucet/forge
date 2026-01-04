package project_mongo

import (
	"context"
	"github.com/JacobDoucet/forge/example/generated/go/project"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, db *mongo.Database, m project.MongoRecord) (primitive.ObjectID, error) {
	coll := db.Collection(CollectionName)
	result, err := coll.InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func CreateBulk(ctx context.Context, db *mongo.Database, m []project.MongoRecord) ([]primitive.ObjectID, error) {
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

func Update(ctx context.Context, db *mongo.Database, m project.MongoRecord, where project.MongoWhereClause) error {
	filter, err := where.GetLookupQuery()
	if err != nil {
		return err
	}
	m.Id = nil

	coll := db.Collection(CollectionName)
	_, err = coll.UpdateOne(ctx, filter, bson.M{"$set": m})
	return err
}

func UpdateMany(ctx context.Context, db *mongo.Database, m project.MongoRecord, ids []primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	filter := bson.M{"_id": bson.M{"$in": ids}}
	m.Id = nil
	_, err := coll.UpdateMany(ctx, filter, bson.M{"$set": m})
	return err
}
