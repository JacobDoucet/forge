package task_mongo

import (
	"context"
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookupOptions struct {
	Projection task.Projection
	Sort       task.MongoSortParams
	Limit      int
	Skip       int
}

func (lo *LookupOptions) limit() int {
	if lo.Limit > 0 {
		return lo.Limit
	}
	return 100
}

type WhereClause struct {
	Task task.MongoWhereClause
}

func aggregateWithRefs(ctx context.Context, where WhereClause, collection *mongo.Collection, lookupOptions LookupOptions) (QueryResult, error) {
	whereTask, err := where.Task.GetLookupQuery()
	if err != nil {
		return QueryResult{}, err
	}

	sortStage := bson.D{}

	if lookupOptions.Sort.AssigneeId > 0 {
		sortStage = append(sortStage, bson.E{Key: "assigneeId", Value: 1})
	} else if lookupOptions.Sort.AssigneeId < 0 {
		sortStage = append(sortStage, bson.E{Key: "assigneeId", Value: -1})
	}
	if lookupOptions.Sort.CreatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: 1})
	} else if lookupOptions.Sort.CreatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: -1})
	}
	if lookupOptions.Sort.Status > 0 {
		sortStage = append(sortStage, bson.E{Key: "status", Value: 1})
	} else if lookupOptions.Sort.Status < 0 {
		sortStage = append(sortStage, bson.E{Key: "status", Value: -1})
	}
	if lookupOptions.Sort.UpdatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: 1})
	} else if lookupOptions.Sort.UpdatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: -1})
	}

	dataPipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: whereTask},
		},
		{
			{Key: "$project", Value: lookupOptions.Projection.ToBson()},
		},
	}

	if len(sortStage) > 0 {
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$sort", Value: sortStage},
		})
	}

	if lookupOptions.Skip > 0 {
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$skip", Value: lookupOptions.Skip},
		})
	}

	dataPipeline = append(dataPipeline, bson.D{
		{Key: "$limit", Value: lookupOptions.limit()},
	})

	// Add the $facet stage
	pipeline := mongo.Pipeline{
		{{
			Key: "$facet",
			Value: bson.D{
				{Key: "data", Value: dataPipeline},
				{Key: "count", Value: bson.A{
					bson.D{
						{
							Key:   "$match",
							Value: whereTask,
						},
					},
					bson.D{
						{
							Key:   "$count",
							Value: "total",
						},
					},
				}},
			},
		}},
	}

	// Execute the aggregation pipeline
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return QueryResult{}, err
	}

	// Decode the result
	var aggregatedResult []struct {
		Data  []Model `bson:"data"`
		Count []struct {
			Total int `bson:"total"`
		} `bson:"count"`
	}
	if err := cur.All(ctx, &aggregatedResult); err != nil {
		return QueryResult{}, errors.Join(cur.Close(ctx), err)
	}

	var totalCount int
	if len(aggregatedResult) > 0 && len(aggregatedResult[0].Count) > 0 {
		totalCount = aggregatedResult[0].Count[0].Total
	}

	return QueryResult{Data: aggregatedResult[0].Data, Count: totalCount}, cur.Close(ctx)
}
