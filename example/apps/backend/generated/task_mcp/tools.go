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
		projection := task_api.NewProjection(true)
		model, resultProjection, err := props.Api.SelectById(ctx, actor, task.SelectByIdQuery{Id: input.Id}, projection)
		if err != nil {
			if props.OnError != nil {
				props.OnError("get_task", err)
			}
			return nil, nil, err
		}
		httpRec, err := model.Model.ToHTTPRecord(resultProjection.Projection)
		if err != nil {
			return nil, nil, err
		}
		payload, err := json.Marshal(httpRec)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, httpRec, nil
	}
}

// SearchInput is the input schema for the search_tasks MCP tool.
//
// The Query field is the generated task.HTTPWhereClause. Because the
// MCP Go SDK infers a JSON Schema from this struct, every supported filter
// field (Eq/Ne/Gt/Gte/Lt/Lte/In/Nin/Like/Exists/nested clauses, ...) is
// advertised to the agent via the tool schema, so the LLM can discover the
// search vocabulary without any extra documentation. All fields use pointer
// types with `,omitempty` so any field omitted from the JSON input stays
// unfiltered and does not appear as required in the schema.
type SearchInput struct {
	Query task.HTTPWhereClause `json:"query,omitempty" jsonschema:"filters matching the generated Task WhereClause schema; omit fields to leave them unfiltered"`
	Limit int                  `json:"limit,omitempty" jsonschema:"maximum number of records to return"`
	Skip  int                  `json:"skip,omitempty" jsonschema:"number of records to skip for pagination"`
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
		where, err := input.Query.ToWhereClause()
		if err != nil {
			return nil, nil, err
		}
		result, resultProjection, err := props.Api.Search(ctx, actor, where, task_api.QueryOptions{
			Limit: input.Limit,
			Skip:  input.Skip,
		})
		if err != nil {
			if props.OnError != nil {
				props.OnError("search_tasks", err)
			}
			return nil, nil, err
		}
		httpResult, err := task_api.ToHTTPQueryResult(result, resultProjection)
		if err != nil {
			return nil, nil, err
		}
		payload, err := json.Marshal(httpResult)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, httpResult, nil
	}
}

// CreateInput is the input schema for the create_task MCP tool.
// The data field is the generated task.HTTPRecord, so the MCP schema
// advertises the full Task shape (camelCase field names, matching
// the HTTP API). All fields are optional in the schema; required semantics are
// enforced by the underlying service layer.
type CreateInput struct {
	Data task.HTTPRecord `json:"data" jsonschema:"the Task to create, matching the HTTPRecord schema"`
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
		obj, err := input.Data.ToModel()
		if err != nil {
			return nil, nil, err
		}
		projection := task.NewProjection(true)
		model, resultProjection, err := props.Api.Create(ctx, actor, obj, projection)
		if err != nil {
			if props.OnError != nil {
				props.OnError("create_task", err)
			}
			return nil, nil, err
		}
		httpRec, err := model.ToHTTPRecord(resultProjection)
		if err != nil {
			return nil, nil, err
		}
		payload, err := json.Marshal(httpRec)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, httpRec, nil
	}
}

// UpdateInput is the input schema for the update_task MCP tool.
// The data field is the generated task.HTTPRecord (must include id).
type UpdateInput struct {
	Data task.HTTPRecord `json:"data" jsonschema:"the Task to update; must include id"`
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
		obj, err := input.Data.ToModel()
		if err != nil {
			return nil, nil, err
		}
		projection := task.NewProjection(true)
		model, resultProjection, err := props.Api.Update(ctx, actor, obj, projection)
		if err != nil {
			if props.OnError != nil {
				props.OnError("update_task", err)
			}
			return nil, nil, err
		}
		httpRec, err := model.ToHTTPRecord(resultProjection)
		if err != nil {
			return nil, nil, err
		}
		payload, err := json.Marshal(httpRec)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(payload)}},
		}, httpRec, nil
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
