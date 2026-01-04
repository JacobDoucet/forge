package task_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func Search(ctx context.Context, db *mongo.Database, where WhereClause, lookupOptions LookupOptions) (QueryResult, error) {
	collection := db.Collection(CollectionName)
	return aggregateWithRefs(ctx, where, collection, lookupOptions)
}
