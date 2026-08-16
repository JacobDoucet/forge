package mcp_register

import (
	"context"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_mcp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterProps carries dependencies for registering all generated MCP tools.
type RegisterProps struct {
	ResolveActor func(ctx context.Context) (permissions.Actor, error)
	OnError      func(tool string, err error)
}

// RegisterTools registers all generated MCP tools onto the supplied server.
// Each object's tools delegate to the api.Client that Forge already exposes,
// so hooks and permissions apply automatically.
func RegisterTools(server *mcp.Server, client api.Client, props RegisterProps) error {
	if err := task_mcp.RegisterTools(server, task_mcp.HandlerProps{
		Api:          client.Task(),
		ResolveActor: props.ResolveActor,
		OnError:      props.OnError,
	}); err != nil {
		return err
	}
	return nil
}
