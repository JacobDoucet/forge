package actor_trace

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	ActorId   bool `json:"actorId"`
	ActorName bool `json:"actorName"`
	ActorType bool `json:"actorType"`
	At        bool `json:"at"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		ActorId:   defaultVal,
		ActorName: defaultVal,
		ActorType: defaultVal,
		At:        defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	if p.ActorId {
		projection["actorId"] = 1
	}
	if p.ActorName {
		projection["actorName"] = 1
	}
	if p.ActorType {
		projection["actorType"] = 1
	}
	if p.At {
		projection["at"] = 1
	}
	return projection
}
