package user

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                  bool                   `json:"id"`
	ActorRoles          bool                   `json:"actorRoles"`
	ActorRolesFields    actor_role.Projection  `json:"actorRolesFields,omitempty"`
	Created             bool                   `json:"created"`
	CreatedFields       actor_trace.Projection `json:"createdFields,omitempty"`
	Email               bool                   `json:"email"`
	FirstName           bool                   `json:"firstName"`
	Language            bool                   `json:"language"`
	LastName            bool                   `json:"lastName"`
	Role                bool                   `json:"role"`
	Updated             bool                   `json:"updated"`
	UpdatedFields       actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByUser       bool                   `json:"updatedByUser"`
	UpdatedByUserFields actor_trace.Projection `json:"updatedByUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                  defaultVal,
		ActorRoles:          defaultVal,
		ActorRolesFields:    actor_role.NewProjection(defaultVal),
		Created:             defaultVal,
		CreatedFields:       actor_trace.NewProjection(defaultVal),
		Email:               defaultVal,
		FirstName:           defaultVal,
		Language:            defaultVal,
		LastName:            defaultVal,
		Role:                defaultVal,
		Updated:             defaultVal,
		UpdatedFields:       actor_trace.NewProjection(defaultVal),
		UpdatedByUser:       defaultVal,
		UpdatedByUserFields: actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.ActorRoles {
		if p.ActorRolesFields.Role {
			projection["actorRoles.role"] = 1
		}
	}
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
	if p.Email {
		projection["email"] = 1
	}
	if p.FirstName {
		projection["firstName"] = 1
	}
	if p.Language {
		projection["language"] = 1
	}
	if p.LastName {
		projection["lastName"] = 1
	}
	if p.Role {
		projection["role"] = 1
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
