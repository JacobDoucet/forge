package types

import "testing"

func TestSanitizeMCPMethod(t *testing.T) {
	cases := []struct {
		in      string
		want    MCPMethod
		wantErr bool
	}{
		{"GET", MCPGet, false},
		{"search", MCPSearch, false},
		{"Create", MCPCreate, false},
		{"UPDATE", MCPUpdate, false},
		{"delete", MCPDelete, false},
		{"PATCH", "", true},
		{"", "", true},
	}
	for _, c := range cases {
		got, err := sanitizeMCPMethod(c.in)
		if c.wantErr {
			if err == nil {
				t.Errorf("sanitizeMCPMethod(%q) expected error, got %q", c.in, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("sanitizeMCPMethod(%q) unexpected error: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("sanitizeMCPMethod(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestObjectMCPHelpers_Defaults(t *testing.T) {
	o := Object{
		Name: "Device",
		MCP: ObjectMCPDef{
			Methods: []string{"GET", "SEARCH", "CREATE", "UPDATE", "DELETE"},
		},
	}
	if err := o.validateMCP(nil); err != nil {
		t.Fatalf("validateMCP: %v", err)
	}
	if !o.HasMCPMethods() {
		t.Fatal("expected HasMCPMethods to be true")
	}
	for _, m := range []MCPMethod{MCPGet, MCPSearch, MCPCreate, MCPUpdate, MCPDelete} {
		if !o.HasMCPMethod(m) {
			t.Errorf("expected HasMCPMethod(%s) true", m)
		}
	}

	// Deterministic default tool names.
	tests := map[MCPMethod]string{
		MCPGet:    "get_device",
		MCPSearch: "search_devices",
		MCPCreate: "create_device",
		MCPUpdate: "update_device",
		MCPDelete: "delete_device",
	}
	for m, want := range tests {
		if got := o.GetMCPToolName(m); got != want {
			t.Errorf("GetMCPToolName(%s) = %q, want %q", m, got, want)
		}
	}

	// Default descriptions should be non-empty.
	for _, m := range []MCPMethod{MCPGet, MCPSearch, MCPCreate, MCPUpdate, MCPDelete} {
		if d := o.GetMCPToolDescription(m); d == "" {
			t.Errorf("GetMCPToolDescription(%s) is empty", m)
		}
	}
}

func TestObjectMCPHelpers_Overrides(t *testing.T) {
	o := Object{
		Name: "Device",
		MCP: ObjectMCPDef{
			Methods: []string{"GET"},
			Get: ObjectMCPMethodDef{
				Name:        "device_lookup",
				Description: "Look up a device by id.",
			},
		},
	}
	if err := o.validateMCP(nil); err != nil {
		t.Fatalf("validateMCP: %v", err)
	}
	if got := o.GetMCPToolName(MCPGet); got != "device_lookup" {
		t.Errorf("expected override name, got %q", got)
	}
	if got := o.GetMCPToolDescription(MCPGet); got != "Look up a device by id." {
		t.Errorf("expected override description, got %q", got)
	}
}

func TestValidateMCP_InvalidMethod(t *testing.T) {
	o := Object{
		Name: "Device",
		MCP:  ObjectMCPDef{Methods: []string{"NOPE"}},
	}
	if err := o.validateMCP(nil); err == nil {
		t.Fatal("expected error for invalid mcp method")
	}
}
