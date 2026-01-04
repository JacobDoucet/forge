package task_api

import (
	"context"
	"github.com/JacobDoucet/forge/example/generated/go/permissions"
	"github.com/JacobDoucet/forge/example/generated/go/task"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query task.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query task.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj task.Model, projection task.Projection) (task.Model, task.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj task.Model, projection task.Projection) (task.Model, task.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query task.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj task.Model, projection task.Projection) (task.Model, error)
	Update(ctx context.Context, obj task.Model, where task.WhereClause, projection task.Projection) (task.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	task.Model
}

type WhereClause struct {
	Task task.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       task.SortParams
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
	Sort       task.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	task.Projection `json:",inline"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Projection: task.NewProjection(defaultVal),
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = task.ProjectReadPermissions(projection.Projection, actor)

	return projection
}
