package api

import (
	"github.com/JacobDoucet/forge/example/generated/go/event_api"
	"github.com/JacobDoucet/forge/example/generated/go/project_api"
	"github.com/JacobDoucet/forge/example/generated/go/task_api"
)

type Client interface {
	Event() event_api.Client
	Project() project_api.Client
	Task() task_api.Client
	ValidateClients() error
}

type CustomClient struct {
	event   event_api.Client
	project project_api.Client
	task    task_api.Client
}

func (c *CustomClient) ValidateClients() error {
	if c.event == nil {
		c.event = event_api.New(&event_api.UnimplementedClient{})
	}
	if c.project == nil {
		c.project = project_api.New(&project_api.UnimplementedClient{})
	}
	if c.task == nil {
		c.task = task_api.New(&task_api.UnimplementedClient{})
	}
	return nil
}
func (c *CustomClient) UseEventClient(client event_api.Client) *CustomClient {
	if client == nil {
		c.event = event_api.New(&event_api.UnimplementedClient{})
		return c
	}
	c.event = client
	return c
}

func (c *CustomClient) Event() event_api.Client {
	return c.event
}
func (c *CustomClient) UseProjectClient(client project_api.Client) *CustomClient {
	if client == nil {
		c.project = project_api.New(&project_api.UnimplementedClient{})
		return c
	}
	c.project = client
	return c
}

func (c *CustomClient) Project() project_api.Client {
	return c.project
}
func (c *CustomClient) UseTaskClient(client task_api.Client) *CustomClient {
	if client == nil {
		c.task = task_api.New(&task_api.UnimplementedClient{})
		return c
	}
	c.task = client
	return c
}

func (c *CustomClient) Task() task_api.Client {
	return c.task
}
