package permissions_api

import (
	"context"
	"fmt"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/coded_error"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
)

type Client interface {
	SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error)
}

func New() Client {
	return &client{}
}

type client struct {
}

func (c *client) SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error) {
	switch actorType {
	}
	return nil, coded_error.NewUnexpectedError(fmt.Sprintf("unhandled actor type %s", actorType))
}
