package task_comment

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	AuthorId  bool `json:"authorId"`
	CreatedAt bool `json:"createdAt"`
	Text      bool `json:"text"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		AuthorId:  defaultVal,
		CreatedAt: defaultVal,
		Text:      defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	if p.AuthorId {
		projection["authorId"] = 1
	}
	if p.CreatedAt {
		projection["createdAt"] = 1
	}
	if p.Text {
		projection["text"] = 1
	}
	return projection
}
