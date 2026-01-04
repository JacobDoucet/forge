package task_api

import (
	"context"
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoBackedClient(db *mongo.Database, hooks ...Hooks) Client {
	return New(&mongoClient{db: db}, hooks...)
}

type mongoClient struct {
	db *mongo.Database
}

func (m *mongoClient) Search(ctx context.Context, where WhereClause, options QueryOptions) (QueryResult, error) {
	projection := options.GetProjection()
	mongoWhereClause, err := where.Task.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := task_mongo.Search(
		ctx,
		m.db,
		task_mongo.WhereClause{
			Task: mongoWhereClause,
		},
		task_mongo.LookupOptions{
			Projection: projection.Projection,
			Sort:       options.Sort.ToMongoSortParams(),
			Limit:      options.Limit,
			Skip:       options.Skip,
		},
	)
	if err != nil {
		return QueryResult{}, err
	}

	modelRecords, err := FromMongoQueryResultDataList(searchResult.Data)
	if err != nil {
		return QueryResult{}, err
	}

	return QueryResult{
		Data:  modelRecords,
		Total: searchResult.Count,
		Skip:  options.Skip,
	}, nil
}

func (m *mongoClient) Create(ctx context.Context, obj task.Model, projection task.Projection) (task.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return task.Model{}, err
	}
	var id primitive.ObjectID
	id, err = task_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return task.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj task.Model, where task.WhereClause, projection task.Projection) (task.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return task.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return task.Model{}, err
	}

	err = task_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return task.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = task_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r task_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []task_mongo.Model) ([]Model, error) {
	ms := make([]Model, len(dbRecords))
	var err error
	for i, r := range dbRecords {
		var iErr error
		ms[i], iErr = FromMongoQueryResultData(r)
		if iErr != nil {
			err = errors.Join(err, iErr)
		}
	}
	return ms, err
}
