package event_api

import (
	"context"
	"errors"
	"github.com/JacobDoucet/forge/example/generated/go/event"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj event.Model, projection event.Projection) (event.Model, error) {
	return event.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj event.Model, where event.WhereClause, projection event.Projection) (event.Model, error) {
	return event.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}
