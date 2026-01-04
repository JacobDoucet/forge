package project

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                  bool                   `json:"id"`
	Created             bool                   `json:"created"`
	CreatedFields       actor_trace.Projection `json:"createdFields,omitempty"`
	Description         bool                   `json:"description"`
	Name                bool                   `json:"name"`
	OwnerId             bool                   `json:"ownerId"`
	Updated             bool                   `json:"updated"`
	UpdatedFields       actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByUser       bool                   `json:"updatedByUser"`
	UpdatedByUserFields actor_trace.Projection `json:"updatedByUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                  defaultVal,
		Created:             defaultVal,
		CreatedFields:       actor_trace.NewProjection(defaultVal),
		Description:         defaultVal,
		Name:                defaultVal,
		OwnerId:             defaultVal,
		Updated:             defaultVal,
		UpdatedFields:       actor_trace.NewProjection(defaultVal),
		UpdatedByUser:       defaultVal,
		UpdatedByUserFields: actor_trace.NewProjection(defaultVal),
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
	if p.Description {
		projection["description"] = 1
	}
	if p.Name {
		projection["name"] = 1
	}
	if p.OwnerId {
		projection["ownerId"] = 1
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
	if p.UpdatedByUser {
		if p.UpdatedByUserFields.ActorId {
			projection["updatedByUser.actorId"] = 1
		}
		if p.UpdatedByUserFields.ActorName {
			projection["updatedByUser.actorName"] = 1
		}
		if p.UpdatedByUserFields.ActorType {
			projection["updatedByUser.actorType"] = 1
		}
		if p.UpdatedByUserFields.At {
			projection["updatedByUser.at"] = 1
		}
	}
	return projection
}
