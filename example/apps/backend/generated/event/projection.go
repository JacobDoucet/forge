package event

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_subject"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id             bool                     `json:"id"`
	Created        bool                     `json:"created"`
	CreatedFields  actor_trace.Projection   `json:"createdFields,omitempty"`
	Subjects       bool                     `json:"subjects"`
	SubjectsFields event_subject.Projection `json:"subjectsFields,omitempty"`
	Type           bool                     `json:"type"`
	Updated        bool                     `json:"updated"`
	UpdatedFields  actor_trace.Projection   `json:"updatedFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:             defaultVal,
		Created:        defaultVal,
		CreatedFields:  actor_trace.NewProjection(defaultVal),
		Subjects:       defaultVal,
		SubjectsFields: event_subject.NewProjection(defaultVal),
		Type:           defaultVal,
		Updated:        defaultVal,
		UpdatedFields:  actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.Created {
		if p.CreatedFields.ActorId {
			projection["created.actorId"] = 1
		}
		if p.CreatedFields.ActorName {
			projection["created.actorName"] = 1
		}
		if p.CreatedFields.ActorType {
			projection["created.actorType"] = 1
		}
		if p.CreatedFields.At {
			projection["created.at"] = 1
		}
	}
	if p.Subjects {
		if p.SubjectsFields.SubjectId {
			projection["subjects.subjectId"] = 1
		}
		if p.SubjectsFields.SubjectType {
			projection["subjects.subjectType"] = 1
		}
	}
	if p.Type {
		projection["type"] = 1
	}
	if p.Updated {
		if p.UpdatedFields.ActorId {
			projection["updated.actorId"] = 1
		}
		if p.UpdatedFields.ActorName {
			projection["updated.actorName"] = 1
		}
		if p.UpdatedFields.ActorType {
			projection["updated.actorType"] = 1
		}
		if p.UpdatedFields.At {
			projection["updated.at"] = 1
		}
	}
	return projection
}
