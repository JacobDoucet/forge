package task

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_comment"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id             bool                    `json:"id"`
	AssigneeId     bool                    `json:"assigneeId"`
	Comments       bool                    `json:"comments"`
	CommentsFields task_comment.Projection `json:"commentsFields,omitempty"`
	Created        bool                    `json:"created"`
	CreatedFields  actor_trace.Projection  `json:"createdFields,omitempty"`
	Description    bool                    `json:"description"`
	DueDate        bool                    `json:"dueDate"`
	Priority       bool                    `json:"priority"`
	Status         bool                    `json:"status"`
	Tags           bool                    `json:"tags"`
	Title          bool                    `json:"title"`
	Updated        bool                    `json:"updated"`
	UpdatedFields  actor_trace.Projection  `json:"updatedFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:             defaultVal,
		AssigneeId:     defaultVal,
		Comments:       defaultVal,
		CommentsFields: task_comment.NewProjection(defaultVal),
		Created:        defaultVal,
		CreatedFields:  actor_trace.NewProjection(defaultVal),
		Description:    defaultVal,
		DueDate:        defaultVal,
		Priority:       defaultVal,
		Status:         defaultVal,
		Tags:           defaultVal,
		Title:          defaultVal,
		Updated:        defaultVal,
		UpdatedFields:  actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.AssigneeId {
		projection["assigneeId"] = 1
	}
	if p.Comments {
		if p.CommentsFields.AuthorId {
			projection["comments.authorId"] = 1
		}
		if p.CommentsFields.CreatedAt {
			projection["comments.createdAt"] = 1
		}
		if p.CommentsFields.Text {
			projection["comments.text"] = 1
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
	if p.Description {
		projection["description"] = 1
	}
	if p.DueDate {
		projection["dueDate"] = 1
	}
	if p.Priority {
		projection["priority"] = 1
	}
	if p.Status {
		projection["status"] = 1
	}
	if p.Tags {
		projection["tags"] = 1
	}
	if p.Title {
		projection["title"] = 1
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
