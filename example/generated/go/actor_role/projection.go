package actor_role

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Role bool `json:"role"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Role: defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	if p.Role {
		projection["role"] = 1
	}
	return projection
}
