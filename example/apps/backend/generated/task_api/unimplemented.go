package task_api

import (
	"context"
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj task.Model, projection task.Projection) (task.Model, error) {
	return task.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj task.Model, where task.WhereClause, projection task.Projection) (task.Model, error) {
	return task.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}
