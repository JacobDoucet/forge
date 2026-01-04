package api

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/project_api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_api"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoBackedClient(db *mongo.Database) Client {
	return &mongoClient{
		event:   event_api.NewMongoBackedClient(db),
		project: project_api.NewMongoBackedClient(db),
		task:    task_api.NewMongoBackedClient(db),
	}
}

type mongoClient struct {
	event   event_api.Client
	project project_api.Client
	task    task_api.Client
}

func (m *mongoClient) ValidateClients() error {
	return nil
}
func (c *mongoClient) Event() event_api.Client {
	return c.event
}
func (c *mongoClient) Project() project_api.Client {
	return c.project
}
func (c *mongoClient) Task() task_api.Client {
	return c.task
}
