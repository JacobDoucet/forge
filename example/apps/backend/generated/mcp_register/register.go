package mcp_register

import (
	"errors"
	"context"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_mcp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterProps carries dependencies for registering all generated MCP tools.
type RegisterProps struct {
	RequireAllClients bool
	ResolveActor func(ctx context.Context) (permissions.Actor, error)
	OnError      func(tool string, err error)
}

// RegisterTools registers tools for configured model clients onto the supplied server.
// Each object's tools delegate to the api.Client that Forge already exposes,
// so hooks and permissions apply automatically.
func RegisterTools(server *mcp.Server, client api.Client, props RegisterProps) error {
	if server == nil {
		return errors.New("server is required")
	}
	if client == nil {
		return errors.New("client is required")
	}
	if props.ResolveActor == nil {
		return errors.New("resolveActor is required")
	}
	if props.RequireAllClients && client.Task() == nil {
		return errors.New("missing MCP API client: Task")
	}
	if modelClient := client.Task(); modelClient != nil {
		if err := task_mcp.RegisterTools(server, task_mcp.HandlerProps{
			Api:          modelClient,
			ResolveActor: props.ResolveActor,
			OnError:      props.OnError,
		}); err != nil {
			return err
		}
	}
	return nil
}
