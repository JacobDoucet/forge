package event_subject

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	SubjectId   bool `json:"subjectId"`
	SubjectType bool `json:"subjectType"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		SubjectId:   defaultVal,
		SubjectType: defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	if p.SubjectId {
		projection["subjectId"] = 1
	}
	if p.SubjectType {
		projection["subjectType"] = 1
	}
	return projection
}
