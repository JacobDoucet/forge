package types

import (
	"fmt"
	"strings"
)

// ObjectMCPDef describes MCP (Model Context Protocol) tool generation for an object.
// It is independent of the HTTP config: an object may expose MCP tools without HTTP,
// and vice versa. Both share the same underlying api.Client service layer.
type ObjectMCPDef struct {
	// Methods is the list of MCP tool kinds to generate for this object.
	// Supported values: GET, SEARCH, CREATE, UPDATE, DELETE.
	Methods []string `yaml:"methods"`

	// Per-method overrides for tool name and description.
	Get    ObjectMCPMethodDef `yaml:"get"`
	Search ObjectMCPMethodDef `yaml:"search"`
	Create ObjectMCPMethodDef `yaml:"create"`
	Update ObjectMCPMethodDef `yaml:"update"`
	Delete ObjectMCPMethodDef `yaml:"delete"`
}

// ObjectMCPMethodDef holds optional overrides for a single generated MCP tool.
type ObjectMCPMethodDef struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type MCPMethod string

const (
	MCPGet    MCPMethod = "GET"
	MCPSearch MCPMethod = "SEARCH"
	MCPCreate MCPMethod = "CREATE"
	MCPUpdate MCPMethod = "UPDATE"
	MCPDelete MCPMethod = "DELETE"
)

func sanitizeMCPMethod(s string) (MCPMethod, error) {
	switch strings.ToUpper(s) {
	case "GET":
		return MCPGet, nil
	case "SEARCH":
		return MCPSearch, nil
	case "CREATE":
		return MCPCreate, nil
	case "UPDATE":
		return MCPUpdate, nil
	case "DELETE":
		return MCPDelete, nil
	default:
		return "", fmt.Errorf("invalid MCP method: %s", s)
	}
}
