package http_server

import (
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_http"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/project_http"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_http"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/user_http"
	"net/http"
)

type ServeMuxProps struct {
	RequireAllClients    bool
	ResolveActor         func(r *http.Request) (permissions.Actor, error)
	EventMetadataHooks   []event_http.MetadataHooks
	ProjectMetadataHooks []project_http.MetadataHooks
	TaskMetadataHooks    []task_http.MetadataHooks
	UserMetadataHooks    []user_http.MetadataHooks
	OnError              func(handler string, e error)
}

func ServeMux(client api.Client, props ServeMuxProps) (*http.ServeMux, error) {
	if client == nil {
		return nil, errors.New("client is required")
	}
	if props.ResolveActor == nil {
		return nil, errors.New("resolveActor is required")
	}
	serveMux := http.NewServeMux()

	eventApi := client.Event()
	if eventApi != nil {
		eventServeMux, err := event_http.RegisterRoutes(event_http.HandlerProps{
			Api:           eventApi,
			ResolveActor:  props.ResolveActor,
			MetadataHooks: props.EventMetadataHooks,
			OnError:       props.OnError,
		})
		if err != nil {
			return nil, err
		}
		serveMux.Handle("/events/", http.StripPrefix("/events", eventServeMux))
	} else if props.RequireAllClients {
		return nil, errors.New("missing HTTP API client: Event")
	}

	projectApi := client.Project()
	if projectApi != nil {
		projectServeMux, err := project_http.RegisterRoutes(project_http.HandlerProps{
			Api:           projectApi,
			ResolveActor:  props.ResolveActor,
			MetadataHooks: props.ProjectMetadataHooks,
			OnError:       props.OnError,
		})
		if err != nil {
			return nil, err
		}
		serveMux.Handle("/projects/", http.StripPrefix("/projects", projectServeMux))
	} else if props.RequireAllClients {
		return nil, errors.New("missing HTTP API client: Project")
	}

	taskApi := client.Task()
	if taskApi != nil {
		taskServeMux, err := task_http.RegisterRoutes(task_http.HandlerProps{
			Api:           taskApi,
			ResolveActor:  props.ResolveActor,
			MetadataHooks: props.TaskMetadataHooks,
			OnError:       props.OnError,
		})
		if err != nil {
			return nil, err
		}
		serveMux.Handle("/tasks/", http.StripPrefix("/tasks", taskServeMux))
	} else if props.RequireAllClients {
		return nil, errors.New("missing HTTP API client: Task")
	}

	userApi := client.User()
	if userApi != nil {
		userServeMux, err := user_http.RegisterRoutes(user_http.HandlerProps{
			Api:           userApi,
			ResolveActor:  props.ResolveActor,
			MetadataHooks: props.UserMetadataHooks,
			OnError:       props.OnError,
		})
		if err != nil {
			return nil, err
		}
		serveMux.Handle("/users/", http.StripPrefix("/users", userServeMux))
	} else if props.RequireAllClients {
		return nil, errors.New("missing HTTP API client: User")
	}

	return serveMux, nil
}
