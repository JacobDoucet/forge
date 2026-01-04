package task_http

import (
	"net/http"
	"strings"
)

func formatEndpoint(rootPath, endpoint string) string {
	return strings.ReplaceAll(rootPath+"/"+endpoint, "//", "/")
}

func RegisterRoutes(handlerProps HandlerProps) (*http.ServeMux, error) {
	if err := handlerProps.Validate(); err != nil {
		return nil, err
	}

	routes := http.NewServeMux()

	searchHandler, err := GetSearchHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/search", searchHandler)

	selectByIdHandler, err := GetSelectByIdHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/id/{id}", selectByIdHandler)

	createHandler, err := GetCreateHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/create", createHandler)

	deleteHandler, err := GetDeleteHandler(handlerProps)
	if err != nil {
		return nil, err
	}
	routes.HandleFunc("/delete/{id}", deleteHandler)

	return routes, nil
}
