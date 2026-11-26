package integration_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JacobDoucet/forge/example/apps/backend/generated/api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/http_server"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/mcp_register"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestHTTPRegistration(t *testing.T) {
	resolve := func(*http.Request) (permissions.Actor, error) { return nil, nil }
	for _, tc := range []struct {
		name string
		client api.Client
		strict bool
		wantError string
		wantEvent bool
		wantUser bool
	}{
		{"empty", &api.CustomClient{}, false, "", false, false},
		{"partial", (&api.CustomClient{}).UseEventClient(event_api.New(&event_api.UnimplementedClient{})), false, "", true, false},
		{"strict partial", (&api.CustomClient{}).UseEventClient(event_api.New(&event_api.UnimplementedClient{})), true, "Project", false, false},
		{"strict empty", &api.CustomClient{}, true, "Event", false, false},
		{"nil client", nil, false, "client is required", false, false},
		{"complete", completeClient(), true, "", true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mux, err := http_server.ServeMux(tc.client, http_server.ServeMuxProps{ResolveActor: resolve, RequireAllClients: tc.strict})
			if tc.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantError) { t.Fatalf("expected %q, got %v", tc.wantError, err) }
				return
			}
			if err != nil { t.Fatal(err) }
			for path, want := range map[string]bool{"/events/": tc.wantEvent, "/users/": tc.wantUser} {
				req := httptest.NewRequest("GET", path, nil)
				_, pattern := mux.Handler(req)
				if (pattern != "") != want { t.Errorf("%s: pattern %q, want registered=%v", path, pattern, want) }
				if !want {
					rec := httptest.NewRecorder()
					mux.ServeHTTP(rec, req)
					if rec.Code != http.StatusNotFound { t.Errorf("%s: got %d, want 404", path, rec.Code) }
				}
			}
		})
	}
	if _, err := http_server.ServeMux(&api.CustomClient{}, http_server.ServeMuxProps{}); err == nil { t.Fatal("expected missing resolver error") }
}

func completeClient() *api.CustomClient {
	c := api.NewUnimplementedClient()
	return &c
}

func TestMCPRegistration(t *testing.T) {
	for _, tc := range []struct {
		name string
		client api.Client
		strict bool
		wantError bool
		wantTools bool
	}{
		{"empty", &api.CustomClient{}, false, false, false},
		{"configured", (&api.CustomClient{}).UseTaskClient(task_api.New(&task_api.UnimplementedClient{})), false, false, true},
		{"strict missing", &api.CustomClient{}, true, true, false},
		{"strict complete", completeClient(), true, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			server := mcp.NewServer(&mcp.Implementation{Name:"test-server", Version:"1"}, nil)
			err := mcp_register.RegisterTools(server, tc.client, mcp_register.RegisterProps{
				RequireAllClients: tc.strict,
				ResolveActor: func(context.Context) (permissions.Actor,error) { return nil,nil },
			})
			if (err != nil) != tc.wantError { t.Fatalf("registration error: %v",err) }
			if tc.wantError { return }
			st, ct := mcp.NewInMemoryTransports()
			ss, err := server.Connect(ctx, st, nil)
			if err != nil { t.Fatal(err) }
			defer ss.Close()
			client := mcp.NewClient(&mcp.Implementation{Name:"test-client",Version:"1"},nil)
			cs, err := client.Connect(ctx, ct,nil)
			if err != nil { t.Fatal(err) }
			defer cs.Close()
			result, err := cs.ListTools(ctx,nil)
			if err != nil { t.Fatal(err) }
			found := false
			for _, tool := range result.Tools { if tool.Name == "get_task" { found = true } }
			if found != tc.wantTools { t.Fatalf("get_task present=%v, want %v",found,tc.wantTools) }
			if !tc.wantTools && len(result.Tools) != 0 { t.Fatalf("unexpected tools: %v",result.Tools) }
		})
	}
}
