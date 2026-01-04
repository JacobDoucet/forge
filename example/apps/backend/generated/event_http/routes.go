package event_http

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

	return routes, nil
}
