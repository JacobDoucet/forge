package types

import (
	"fmt"
	"strings"
)

type ObjectHTTPDef struct {
	Endpoint string   `yaml:"endpoint"`
	Methods  []string `yaml:"methods"`
}

type HttpMethod string

const (
	HttpGET    HttpMethod = "GET"
	HttpPOST   HttpMethod = "POST"
	HttpPUT    HttpMethod = "PUT"
	HttpDELETE HttpMethod = "DELETE"
	HttpPATCH  HttpMethod = "PATCH"
)

func sanitizeHTTPMethod(s string) (HttpMethod, error) {
	switch strings.ToUpper(s) {
	case "GET":
		return HttpGET, nil
	case "POST":
		return HttpPOST, nil
	case "PUT":
		return HttpPUT, nil
	case "DELETE":
		return HttpDELETE, nil
	case "PATCH":
		return HttpPATCH, nil
	default:
		return "", fmt.Errorf("invalid HTTP method: %s", s)
	}
}
