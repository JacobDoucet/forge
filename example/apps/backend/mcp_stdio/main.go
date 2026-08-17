// Package main provides a minimal MCP stdio server that exposes all Forge-generated
// MCP tools for this example backend. It runs against the same api.Client as the
// HTTP server, so hooks and permissions apply identically.
//
// Run with:
//
//	go run ./mcp_stdio
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JacobDoucet/forge/example/apps/backend/generated/api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/mcp_register"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := getEnv("DB_NAME", "forge_example")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() { _ = client.Disconnect(context.Background()) }()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	apiClient := api.NewMongoBackedClient(client.Database(dbName))

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "forge-example",
		Version: "v0.1.0",
	}, nil)

	if err := mcp_register.RegisterTools(server, apiClient, mcp_register.RegisterProps{
		ResolveActor: func(_ context.Context) (permissions.Actor, error) {
			// In a real deployment resolve actor from OAuth/session/JWT.
			return &permissions.SuperActor{
				Name:      "mcp admin",
				Username:  "mcp.admin",
				AdminName: "mcp admin",
			}, nil
		},
		OnError: func(tool string, err error) {
			log.Printf("mcp tool %s error: %v", tool, err)
		},
	}); err != nil {
		log.Fatalf("Failed to register MCP tools: %v", err)
	}

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("MCP server exited: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
