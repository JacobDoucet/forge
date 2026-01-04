package user_api

import (
	"context"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/user"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query user.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query user.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectByEmailIdx(ctx context.Context, actor permissions.Actor, query user.SelectByEmailIdxQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj user.Model, projection user.Projection) (user.Model, user.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj user.Model, projection user.Projection) (user.Model, user.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query user.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj user.Model, projection user.Projection) (user.Model, error)
	Update(ctx context.Context, obj user.Model, where user.WhereClause, projection user.Projection) (user.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	user.Model
}

type WhereClause struct {
	User user.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       user.SortParams
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
	Sort       user.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	user.Projection `json:",inline"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Projection: user.NewProjection(defaultVal),
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = user.ProjectReadPermissions(projection.Projection, actor)

	return projection
}
