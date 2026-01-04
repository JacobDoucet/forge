package types

import (
	"fmt"
	"sort"
)

type EventRegistry interface {
	RegisterEvent(event string) error
	GetEnums() []Enum
	GetObjects(permissions *ObjectPermissions) []Object
}

func NewEventRegistry() EventRegistry {
	return &eventRegistry{
		events: make(map[string]struct{}),
	}
}

type eventRegistry struct {
	events map[string]struct{}
}

func (e *eventRegistry) RegisterEvent(event string) error {
	if _, ok := e.events[event]; ok {
		return fmt.Errorf("event %s already registered", event)
	}
	e.events[event] = struct{}{}
	return nil
}

func (e *eventRegistry) listEvents() []string {
	var events []string
	for event := range e.events {
		events = append(events, event)
	}
	sort.Strings(events)
	return events
}

func (e *eventRegistry) GetEnums() []Enum {
	return []Enum{
		{
			Name:   "EventType",
			Type:   FieldTypeString,
			Values: e.listEvents(),
		},
	}
}

func (e *eventRegistry) GetObjects(permissions *ObjectPermissions) []Object {
	var eventPermissions ObjectPermissions
	if permissions != nil {
		eventPermissions = *permissions
	}
	return []Object{
		{
			Name: "EventSubject",
			Fields: []Field{
				{Name: "subjectType", Type: "Model"},
				{Name: "subjectId", Type: "string"},
			},
		},
		{
			Name:        "Event",
			Permissions: eventPermissions,
			HTTP: ObjectHTTPDef{
				Endpoint: "events",
				Methods:  []string{"GET"},
			},
			Collection: []Collection{
				{Type: "mongo", Name: "events"},
			},
			Fields: []Field{
				{Name: "type", Type: "EventType"},
				{Name: "subjects", Type: "List<EventSubject>"},
			},
			Indexes: []Index{
				{Name: "type", Fields: []IndexField{
					{Name: "type"},
					{Name: "created.at"},
				}},
				{
					Name: "subject",
					Fields: []IndexField{
						{Name: "subjects.subjectId"},
						{Name: "created.at"},
					},
				},
				{
					Name: "subject_type",
					Fields: []IndexField{
						{Name: "subjects.subjectId"},
						{Name: "type"},
						{Name: "created.at"},
					},
				},
			},
		},
	}
}
