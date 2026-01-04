package http_server

import (
	"github.com/JacobDoucet/forge/example/generated/go/api"
	"github.com/JacobDoucet/forge/example/generated/go/event_http"
	"github.com/JacobDoucet/forge/example/generated/go/permissions"
	"github.com/JacobDoucet/forge/example/generated/go/project_http"
	"github.com/JacobDoucet/forge/example/generated/go/task_http"
	"net/http"
)

type ServeMuxProps struct {
	ResolveActor         func(r *http.Request) (permissions.Actor, error)
	EventMetadataHooks   []event_http.MetadataHooks
	ProjectMetadataHooks []project_http.MetadataHooks
	TaskMetadataHooks    []task_http.MetadataHooks
	OnError              func(handler string, e error)
}

func ServeMux(client api.Client, props ServeMuxProps) (*http.ServeMux, error) {
	serveMux := http.NewServeMux()

	eventApi := client.Event()
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

	projectApi := client.Project()
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

	taskApi := client.Task()
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

	return serveMux, nil
}
