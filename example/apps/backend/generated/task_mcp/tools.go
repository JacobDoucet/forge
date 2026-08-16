package task_mcp

import (
	"context"
	"encoding/json"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HandlerProps carries dependencies required to register the generated MCP tools
// for the Task object. It mirrors the HTTP handler props so the
// same api.Client, actor resolution and error handling can be reused.
type HandlerProps struct {
	Api          task_api.Client
	ResolveActor func(ctx context.Context) (permissions.Actor, error)
	OnError      func(tool string, err error)
}

// GetInput is the input schema for the get_task MCP tool.
type GetInput struct {
	Id string `json:"id" jsonschema:"the id of the Task to fetch"`
}

func handleGet(props HandlerProps) func(ctx context.Context, req *mcp.CallToolRequest, input GetInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input GetInput) (*mcp.CallToolResult, any, error) {
		actor, err := props.ResolveActor(ctx)
		if err != nil {
			if props.OnError != nil {
				props.OnError("get_task", err)
			}
			return nil, nil, err
		}
		model, _, err := props.Api.SelectById(ctx, actor, task.SelectByIdQuery{Id: input.Id}, task_api.NewProjection(true))
		if err != nil {
			if props.OnError != nil {
				props.OnError("get_task", err)
			}
			return nil, nil, err
		}
		payload, err := json.Marshal(model)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, model, nil
	}
}

// SearchInput is the input schema for the search_tasks MCP tool.
type SearchInput struct {
	Limit int `json:"limit,omitempty" jsonschema:"maximum number of records to return"`
	Skip  int `json:"skip,omitempty" jsonschema:"number of records to skip for pagination"`
}

func handleSearch(props HandlerProps) func(ctx context.Context, req *mcp.CallToolRequest, input SearchInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input SearchInput) (*mcp.CallToolResult, any, error) {
		actor, err := props.ResolveActor(ctx)
		if err != nil {
			if props.OnError != nil {
				props.OnError("search_tasks", err)
			}
			return nil, nil, err
		}
		result, _, err := props.Api.Search(ctx, actor, task.WhereClause{}, task_api.QueryOptions{
			Limit: input.Limit,
			Skip:  input.Skip,
		})
		if err != nil {
			if props.OnError != nil {
				props.OnError("search_tasks", err)
			}
			return nil, nil, err
		}
		payload, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, result, nil
	}
}

// CreateInput is the input schema for the create_task MCP tool.
// The data field is a raw JSON object matching the task.Model schema.
type CreateInput struct {
	Data json.RawMessage `json:"data" jsonschema:"the Task model to create, as a JSON object"`
}

func handleCreate(props HandlerProps) func(ctx context.Context, req *mcp.CallToolRequest, input CreateInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input CreateInput) (*mcp.CallToolResult, any, error) {
		actor, err := props.ResolveActor(ctx)
		if err != nil {
			if props.OnError != nil {
				props.OnError("create_task", err)
			}
			return nil, nil, err
		}
		var obj task.Model
		if err := json.Unmarshal(input.Data, &obj); err != nil {
			return nil, nil, err
		}
		model, _, err := props.Api.Create(ctx, actor, obj, task.NewProjection(true))
		if err != nil {
			if props.OnError != nil {
				props.OnError("create_task", err)
			}
			return nil, nil, err
		}
		payload, err := json.Marshal(model)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, model, nil
	}
}

// UpdateInput is the input schema for the update_task MCP tool.
type UpdateInput struct {
	Data json.RawMessage `json:"data" jsonschema:"the Task model to update, as a JSON object including id"`
}

func handleUpdate(props HandlerProps) func(ctx context.Context, req *mcp.CallToolRequest, input UpdateInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input UpdateInput) (*mcp.CallToolResult, any, error) {
		actor, err := props.ResolveActor(ctx)
		if err != nil {
			if props.OnError != nil {
				props.OnError("update_task", err)
			}
			return nil, nil, err
		}
		var obj task.Model
		if err := json.Unmarshal(input.Data, &obj); err != nil {
			return nil, nil, err
		}
		model, _, err := props.Api.Update(ctx, actor, obj, task.NewProjection(true))
		if err != nil {
			if props.OnError != nil {
				props.OnError("update_task", err)
			}
			return nil, nil, err
		}
		payload, err := json.Marshal(model)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, model, nil
	}
}

// DeleteInput is the input schema for the delete_task MCP tool.
type DeleteInput struct {
	Id string `json:"id" jsonschema:"the id of the Task to delete"`
}

// DeleteOutput is a small ack payload returned by the delete tool.
type DeleteOutput struct {
	Ok bool   `json:"ok"`
	Id string `json:"id"`
}

func handleDelete(props HandlerProps) func(ctx context.Context, req *mcp.CallToolRequest, input DeleteInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input DeleteInput) (*mcp.CallToolResult, any, error) {
		actor, err := props.ResolveActor(ctx)
		if err != nil {
			if props.OnError != nil {
				props.OnError("delete_task", err)
			}
			return nil, nil, err
		}
		if err := props.Api.Delete(ctx, actor, input.Id); err != nil {
			if props.OnError != nil {
				props.OnError("delete_task", err)
			}
			return nil, nil, err
		}
		out := DeleteOutput{Ok: true, Id: input.Id}
		payload, _ := json.Marshal(out)
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, out, nil
	}
}

// RegisterTools registers all MCP tools generated for the Task
// object on the supplied server. It delegates to the same api.Client used by
// the HTTP layer, so hooks and permissions are applied uniformly.
func RegisterTools(server *mcp.Server, props HandlerProps) error {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_task",
		Description: "Fetch a single Task by id.",
	}, handleGet(props))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_tasks",
		Description: "Search Task records with optional filters.",
	}, handleSearch(props))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_task",
		Description: "Create a new Task.",
	}, handleCreate(props))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_task",
		Description: "Update an existing Task.",
	}, handleUpdate(props))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_task",
		Description: "Delete a Task by id.",
	}, handleDelete(props))
	return nil
}
