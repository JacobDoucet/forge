package project_api

import (
	"context"
	"errors"
	"github.com/JacobDoucet/forge/example/generated/go/project"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj project.Model, projection project.Projection) (project.Model, error) {
	return project.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj project.Model, where project.WhereClause, projection project.Projection) (project.Model, error) {
	return project.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}
