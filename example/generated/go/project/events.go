package project

import (
	"errors"
	"github.com/JacobDoucet/forge/example/generated/go/enum_model"
	"github.com/JacobDoucet/forge/example/generated/go/event_subject"
)

func (m *Model) AsEventSubject() (event_subject.Model, error) {
	if m.Id == "" {
		return event_subject.Model{}, errors.New("project does not have an id")
	}
	return event_subject.Model{
		SubjectId:   m.Id,
		SubjectType: enum_model.Project,
	}, nil
}
