package types

import "fmt"

type EndpointRegistry interface {
	Get(name string) (ObjectHTTPDef, bool)
	Register(coll ObjectHTTPDef) error
	List() []ObjectHTTPDef
}

func newEndpointRegistry() EndpointRegistry {
	return &endpointRegistry{
		endpoints: make(map[string]ObjectHTTPDef),
	}
}

type endpointRegistry struct {
	endpoints map[string]ObjectHTTPDef
}

func (r *endpointRegistry) Get(name string) (ObjectHTTPDef, bool) {
	ep, ok := r.endpoints[name]
	return ep, ok
}

func (r *endpointRegistry) Register(ep ObjectHTTPDef) error {
	if _, ok := r.endpoints[ep.Endpoint]; ok {
		return fmt.Errorf("endpoint %s already registered", ep.Endpoint)
	}
	r.endpoints[ep.Endpoint] = ep
	return nil
}

func (r *endpointRegistry) List() []ObjectHTTPDef {
	var endpoints []ObjectHTTPDef
	for _, coll := range r.endpoints {
		endpoints = append(endpoints, coll)
	}
	return endpoints
}
