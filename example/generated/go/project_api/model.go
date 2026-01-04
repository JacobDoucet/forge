package project_api

import (
	"context"
	"github.com/JacobDoucet/forge/example/generated/go/permissions"
	"github.com/JacobDoucet/forge/example/generated/go/project"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query project.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query project.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj project.Model, projection project.Projection) (project.Model, project.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj project.Model, projection project.Projection) (project.Model, project.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query project.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj project.Model, projection project.Projection) (project.Model, error)
	Update(ctx context.Context, obj project.Model, where project.WhereClause, projection project.Projection) (project.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	project.Model
}

type WhereClause struct {
	Project project.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       project.SortParams
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
	Sort       project.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	project.Projection `json:",inline"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Projection: project.NewProjection(defaultVal),
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = project.ProjectReadPermissions(projection.Projection, actor)

	return projection
}
