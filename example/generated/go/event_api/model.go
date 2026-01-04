package event_api

import (
	"context"
	"github.com/JacobDoucet/forge/example/generated/go/event"
	"github.com/JacobDoucet/forge/example/generated/go/permissions"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query event.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query event.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj event.Model, projection event.Projection) (event.Model, event.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj event.Model, projection event.Projection) (event.Model, event.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query event.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj event.Model, projection event.Projection) (event.Model, error)
	Update(ctx context.Context, obj event.Model, where event.WhereClause, projection event.Projection) (event.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	event.Model
}

type WhereClause struct {
	Event event.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       event.SortParams
	Limit      int
	Skip       int
}

func (qo *QueryOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type PaginationOptions struct {
	Projection *Projection
	Sort       event.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	event.Projection `json:",inline"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Projection: event.NewProjection(defaultVal),
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = event.ProjectReadPermissions(projection.Projection, actor)

	return projection
}
