package permissions_api

import (
	"context"
	"fmt"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/coded_error"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/user"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/user_api"
)

type Client interface {
	SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error)
	UseUserClient(api user_api.Client) Client
	UseUserProjection(projection user_api.Projection) Client
}

func New() Client {
	return &client{}
}

type client struct {
	user           user_api.Client
	userProjection *user_api.Projection
}

func (c *client) UseUserClient(api user_api.Client) Client {
	c.user = api
	return c
}

func (c *client) UseUserProjection(projection user_api.Projection) Client {
	c.userProjection = &projection
	return c
}

func (c *client) SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error) {
	switch actorType {
	case permissions.ActorTypeUser:
		if c.user == nil {
			return nil, coded_error.NewUnexpectedError("user api not provided")
		}
		projection := user_api.Projection{
			Projection: user.NewProjection(true),
		}
		if c.userProjection != nil {
			projection = *c.userProjection
		}
		actor, _, err := c.user.SelectById(
			ctx,
			permissions.NewSuperActor(),
			user.SelectByIdQuery{Id: actorId},
			projection,
		)
		if err != nil {
			return nil, err
		}
		return &actor, nil
	}
	return nil, coded_error.NewUnexpectedError(fmt.Sprintf("unhandled actor type %s", actorType))
}
