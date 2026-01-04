package user

import (
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_model"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_subject"
)

func (m *Model) AsEventSubject() (event_subject.Model, error) {
	if m.Id == "" {
		return event_subject.Model{}, errors.New("user does not have an id")
	}
	return event_subject.Model{
		SubjectId:   m.Id,
		SubjectType: enum_model.User,
	}, nil
}
