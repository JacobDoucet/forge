package task

import (
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
	"github.com/JacobDoucet/forge/example/generated/go/enum_task_priority"
	"github.com/JacobDoucet/forge/example/generated/go/enum_task_status"
	"github.com/JacobDoucet/forge/example/generated/go/task_comment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id          *primitive.ObjectID         `bson:"_id,omitempty"`
	AssigneeId  *string                     `bson:"assigneeId,omitempty"`
	Comments    *[]task_comment.MongoRecord `bson:"comments,omitempty"`
	Created     *actor_trace.MongoRecord    `bson:"created,omitempty"`
	Description *string                     `bson:"description,omitempty"`
	DueDate     *time.Time                  `bson:"dueDate,omitempty"`
	Priority    *enum_task_priority.Value   `bson:"priority,omitempty"`
	Status      *enum_task_status.Value     `bson:"status,omitempty"`
	Tags        *[]string                   `bson:"tags,omitempty"`
	Title       *string                     `bson:"title,omitempty"`
	Updated     *actor_trace.MongoRecord    `bson:"updated,omitempty"`
}

type MongoUpdateWhereClause struct {
	Id primitive.ObjectID
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id.Hex()
		m.Id = elemid0
	}
	if r.AssigneeId != nil {
		elemassigneeId0 := r.AssigneeId
		m.AssigneeId = *elemassigneeId0
	}
	if r.Comments != nil {
		elemcomments0 := make([]task_comment.Model, 0)
		for _, rcomments0 := range *r.Comments {
			elemcomments1, err := rcomments0.ToModel()
			if err != nil {
				return m, err
			}
			elemcomments0 = append(elemcomments0, elemcomments1)
		}
		m.Comments = elemcomments0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Description != nil {
		elemdescription0 := r.Description
		m.Description = *elemdescription0
	}
	if r.DueDate != nil {
		elemdueDate0 := r.DueDate
		m.DueDate = *elemdueDate0
	}
	if r.Priority != nil {
		elempriority0 := r.Priority
		m.Priority = *elempriority0
	}
	if r.Status != nil {
		elemstatus0 := r.Status
		m.Status = *elemstatus0
	}
	if r.Tags != nil {
		elemtags0 := make([]string, 0)
		for _, rtags0 := range *r.Tags {
			elemtags1 := rtags0
			elemtags0 = append(elemtags0, elemtags1)
		}
		m.Tags = elemtags0
	}
	if r.Title != nil {
		elemtitle0 := r.Title
		m.Title = *elemtitle0
	}
	if r.Updated != nil {
		elemupdated0, err := r.Updated.ToModel()
		if err != nil {
			return m, err
		}
		m.Updated = elemupdated0
	}
	return m, nil
}

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}

type MongoWhereClause struct {
	// id (Ref<Task>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// assigneeId (string) search options
	AssigneeIdEq     *string
	AssigneeIdNe     *string
	AssigneeIdGt     *string
	AssigneeIdGte    *string
	AssigneeIdLt     *string
	AssigneeIdLte    *string
	AssigneeIdIn     *[]string
	AssigneeIdNin    *[]string
	AssigneeIdExists *bool
	AssigneeIdLike   *string
	AssigneeIdNlike  *string
	// comments (List<TaskComment>) search options
	Comments      *task_comment.MongoWhereClause
	CommentsEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// description (string) search options
	DescriptionEq     *string
	DescriptionNe     *string
	DescriptionGt     *string
	DescriptionGte    *string
	DescriptionLt     *string
	DescriptionLte    *string
	DescriptionIn     *[]string
	DescriptionNin    *[]string
	DescriptionExists *bool
	DescriptionLike   *string
	DescriptionNlike  *string
	// dueDate (timestamp) search options
	DueDateEq     *time.Time
	DueDateNe     *time.Time
	DueDateGt     *time.Time
	DueDateGte    *time.Time
	DueDateLt     *time.Time
	DueDateLte    *time.Time
	DueDateIn     *[]time.Time
	DueDateNin    *[]time.Time
	DueDateExists *bool
	// priority (TaskPriority) search options
	PriorityEq     *enum_task_priority.Value
	PriorityNe     *enum_task_priority.Value
	PriorityGt     *enum_task_priority.Value
	PriorityGte    *enum_task_priority.Value
	PriorityLt     *enum_task_priority.Value
	PriorityLte    *enum_task_priority.Value
	PriorityIn     *[]enum_task_priority.Value
	PriorityNin    *[]enum_task_priority.Value
	PriorityExists *bool
	// status (TaskStatus) search options
	StatusEq     *enum_task_status.Value
	StatusNe     *enum_task_status.Value
	StatusGt     *enum_task_status.Value
	StatusGte    *enum_task_status.Value
	StatusLt     *enum_task_status.Value
	StatusLte    *enum_task_status.Value
	StatusIn     *[]enum_task_status.Value
	StatusNin    *[]enum_task_status.Value
	StatusExists *bool
	// tags (List<string>) search options
	TagsEq     *string
	TagsNe     *string
	TagsGt     *string
	TagsGte    *string
	TagsLt     *string
	TagsLte    *string
	TagsIn     *[]string
	TagsNin    *[]string
	TagsExists *bool
	TagsLike   *string
	TagsNlike  *string
	TagsEmpty  *bool
	// title (string) search options
	TitleEq     *string
	TitleNe     *string
	TitleGt     *string
	TitleGte    *string
	TitleLt     *string
	TitleLte    *string
	TitleIn     *[]string
	TitleNin    *[]string
	TitleExists *bool
	TitleLike   *string
	TitleNlike  *string
	// updated (ActorTrace) search options
	Updated *actor_trace.MongoWhereClause
}

type MongoLookup interface {
	GetQueryParts() (bson.A, error)
	GetLookupQuery() (bson.M, error)
}

func (o MongoWhereClause) GetLookupQuery() (bson.M, error) {
	query := bson.M{}
	and, err := o.GetQueryParts()
	if err != nil {
		return nil, err
	}
	if len(and) > 0 {
		query["$and"] = and
	}
	return query, nil
}

func (o MongoWhereClause) GetQueryParts() (bson.A, error) {
	and := bson.A{}
	if o.IdEq != nil {
		query := bson.M{}
		query["_id"] = o.IdEq
		and = append(and, query)
	}
	if o.IdIn != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$in": o.IdIn}
		and = append(and, query)
	}
	if o.IdNin != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$nin": o.IdNin}
		and = append(and, query)
	}
	if o.IdExists != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$exists": *o.IdExists}
		and = append(and, query)
	}
	if o.AssigneeIdEq != nil {
		query := bson.M{}
		query["assigneeId"] = o.AssigneeIdEq
		and = append(and, query)
	}
	if o.AssigneeIdNe != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$ne": o.AssigneeIdNe}
		and = append(and, query)
	}
	if o.AssigneeIdGt != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$gt": o.AssigneeIdGt}
		and = append(and, query)
	}
	if o.AssigneeIdGte != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$gte": o.AssigneeIdGte}
		and = append(and, query)
	}
	if o.AssigneeIdLt != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$lt": o.AssigneeIdLt}
		and = append(and, query)
	}
	if o.AssigneeIdLte != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$lte": o.AssigneeIdLte}
		and = append(and, query)
	}
	if o.AssigneeIdIn != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$in": o.AssigneeIdIn}
		and = append(and, query)
	}
	if o.AssigneeIdNin != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$nin": o.AssigneeIdNin}
		and = append(and, query)
	}
	if o.AssigneeIdExists != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$exists": *o.AssigneeIdExists}
		and = append(and, query)
	}
	if o.AssigneeIdLike != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$regex": o.AssigneeIdLike, "$options": "i"}
		and = append(and, query)
	}
	if o.AssigneeIdNlike != nil {
		query := bson.M{}
		query["assigneeId"] = bson.M{"$not": bson.M{"$regex": o.AssigneeIdNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.Comments != nil {
		query := bson.M{}
		commentsQuery, err := o.Comments.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range commentsQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["comments."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.CommentsEmpty != nil {
		query := bson.M{}
		if *o.CommentsEmpty {
			query["$or"] = bson.A{
				bson.M{"comments": nil},
				bson.M{"comments": bson.A{}},
				bson.M{"comments": bson.M{"$exists": false}},
			}
		} else {
			query["comments"] = bson.M{
				"$ne":     nil,
				"$not":    bson.M{"$size": 0},
				"$exists": true,
			}
		}
		and = append(and, query)
	}
	if o.Created != nil {
		query := bson.M{}
		createdQuery, err := o.Created.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range createdQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["created."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.DescriptionEq != nil {
		query := bson.M{}
		query["description"] = o.DescriptionEq
		and = append(and, query)
	}
	if o.DescriptionNe != nil {
		query := bson.M{}
		query["description"] = bson.M{"$ne": o.DescriptionNe}
		and = append(and, query)
	}
	if o.DescriptionGt != nil {
		query := bson.M{}
		query["description"] = bson.M{"$gt": o.DescriptionGt}
		and = append(and, query)
	}
	if o.DescriptionGte != nil {
		query := bson.M{}
		query["description"] = bson.M{"$gte": o.DescriptionGte}
		and = append(and, query)
	}
	if o.DescriptionLt != nil {
		query := bson.M{}
		query["description"] = bson.M{"$lt": o.DescriptionLt}
		and = append(and, query)
	}
	if o.DescriptionLte != nil {
		query := bson.M{}
		query["description"] = bson.M{"$lte": o.DescriptionLte}
		and = append(and, query)
	}
	if o.DescriptionIn != nil {
		query := bson.M{}
		query["description"] = bson.M{"$in": o.DescriptionIn}
		and = append(and, query)
	}
	if o.DescriptionNin != nil {
		query := bson.M{}
		query["description"] = bson.M{"$nin": o.DescriptionNin}
		and = append(and, query)
	}
	if o.DescriptionExists != nil {
		query := bson.M{}
		query["description"] = bson.M{"$exists": *o.DescriptionExists}
		and = append(and, query)
	}
	if o.DescriptionLike != nil {
		query := bson.M{}
		query["description"] = bson.M{"$regex": o.DescriptionLike, "$options": "i"}
		and = append(and, query)
	}
	if o.DescriptionNlike != nil {
		query := bson.M{}
		query["description"] = bson.M{"$not": bson.M{"$regex": o.DescriptionNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.DueDateEq != nil {
		query := bson.M{}
		query["dueDate"] = o.DueDateEq
		and = append(and, query)
	}
	if o.DueDateNe != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$ne": o.DueDateNe}
		and = append(and, query)
	}
	if o.DueDateGt != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$gt": o.DueDateGt}
		and = append(and, query)
	}
	if o.DueDateGte != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$gte": o.DueDateGte}
		and = append(and, query)
	}
	if o.DueDateLt != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$lt": o.DueDateLt}
		and = append(and, query)
	}
	if o.DueDateLte != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$lte": o.DueDateLte}
		and = append(and, query)
	}
	if o.DueDateIn != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$in": o.DueDateIn}
		and = append(and, query)
	}
	if o.DueDateNin != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$nin": o.DueDateNin}
		and = append(and, query)
	}
	if o.DueDateExists != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$exists": *o.DueDateExists}
		and = append(and, query)
	}
	if o.PriorityEq != nil {
		query := bson.M{}
		query["priority"] = o.PriorityEq
		and = append(and, query)
	}
	if o.PriorityNe != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$ne": o.PriorityNe}
		and = append(and, query)
	}
	if o.PriorityGt != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$gt": o.PriorityGt}
		and = append(and, query)
	}
	if o.PriorityGte != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$gte": o.PriorityGte}
		and = append(and, query)
	}
	if o.PriorityLt != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$lt": o.PriorityLt}
		and = append(and, query)
	}
	if o.PriorityLte != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$lte": o.PriorityLte}
		and = append(and, query)
	}
	if o.PriorityIn != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$in": o.PriorityIn}
		and = append(and, query)
	}
	if o.PriorityNin != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$nin": o.PriorityNin}
		and = append(and, query)
	}
	if o.PriorityExists != nil {
		query := bson.M{}
		query["priority"] = bson.M{"$exists": *o.PriorityExists}
		and = append(and, query)
	}
	if o.StatusEq != nil {
		query := bson.M{}
		query["status"] = o.StatusEq
		and = append(and, query)
	}
	if o.StatusNe != nil {
		query := bson.M{}
		query["status"] = bson.M{"$ne": o.StatusNe}
		and = append(and, query)
	}
	if o.StatusGt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gt": o.StatusGt}
		and = append(and, query)
	}
	if o.StatusGte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gte": o.StatusGte}
		and = append(and, query)
	}
	if o.StatusLt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lt": o.StatusLt}
		and = append(and, query)
	}
	if o.StatusLte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lte": o.StatusLte}
		and = append(and, query)
	}
	if o.StatusIn != nil {
		query := bson.M{}
		query["status"] = bson.M{"$in": o.StatusIn}
		and = append(and, query)
	}
	if o.StatusNin != nil {
		query := bson.M{}
		query["status"] = bson.M{"$nin": o.StatusNin}
		and = append(and, query)
	}
	if o.StatusExists != nil {
		query := bson.M{}
		query["status"] = bson.M{"$exists": *o.StatusExists}
		and = append(and, query)
	}
	if o.TagsEq != nil {
		query := bson.M{}
		query["tags"] = o.TagsEq
		and = append(and, query)
	}
	if o.TagsNe != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$ne": o.TagsNe}
		and = append(and, query)
	}
	if o.TagsGt != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$gt": o.TagsGt}
		and = append(and, query)
	}
	if o.TagsGte != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$gte": o.TagsGte}
		and = append(and, query)
	}
	if o.TagsLt != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$lt": o.TagsLt}
		and = append(and, query)
	}
	if o.TagsLte != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$lte": o.TagsLte}
		and = append(and, query)
	}
	if o.TagsIn != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$in": o.TagsIn}
		and = append(and, query)
	}
	if o.TagsNin != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$nin": o.TagsNin}
		and = append(and, query)
	}
	if o.TagsExists != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$exists": *o.TagsExists}
		and = append(and, query)
	}
	if o.TagsLike != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$regex": o.TagsLike, "$options": "i"}
		and = append(and, query)
	}
	if o.TagsNlike != nil {
		query := bson.M{}
		query["tags"] = bson.M{"$not": bson.M{"$regex": o.TagsNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.TagsEmpty != nil {
		query := bson.M{}
		if *o.TagsEmpty {
			query["$or"] = bson.A{
				bson.M{"tags": nil},
				bson.M{"tags": bson.A{}},
				bson.M{"tags": bson.M{"$exists": false}},
			}
		} else {
			query["tags"] = bson.M{
				"$ne":     nil,
				"$not":    bson.M{"$size": 0},
				"$exists": true,
			}
		}
		and = append(and, query)
	}
	if o.TitleEq != nil {
		query := bson.M{}
		query["title"] = o.TitleEq
		and = append(and, query)
	}
	if o.TitleNe != nil {
		query := bson.M{}
		query["title"] = bson.M{"$ne": o.TitleNe}
		and = append(and, query)
	}
	if o.TitleGt != nil {
		query := bson.M{}
		query["title"] = bson.M{"$gt": o.TitleGt}
		and = append(and, query)
	}
	if o.TitleGte != nil {
		query := bson.M{}
		query["title"] = bson.M{"$gte": o.TitleGte}
		and = append(and, query)
	}
	if o.TitleLt != nil {
		query := bson.M{}
		query["title"] = bson.M{"$lt": o.TitleLt}
		and = append(and, query)
	}
	if o.TitleLte != nil {
		query := bson.M{}
		query["title"] = bson.M{"$lte": o.TitleLte}
		and = append(and, query)
	}
	if o.TitleIn != nil {
		query := bson.M{}
		query["title"] = bson.M{"$in": o.TitleIn}
		and = append(and, query)
	}
	if o.TitleNin != nil {
		query := bson.M{}
		query["title"] = bson.M{"$nin": o.TitleNin}
		and = append(and, query)
	}
	if o.TitleExists != nil {
		query := bson.M{}
		query["title"] = bson.M{"$exists": *o.TitleExists}
		and = append(and, query)
	}
	if o.TitleLike != nil {
		query := bson.M{}
		query["title"] = bson.M{"$regex": o.TitleLike, "$options": "i"}
		and = append(and, query)
	}
	if o.TitleNlike != nil {
		query := bson.M{}
		query["title"] = bson.M{"$not": bson.M{"$regex": o.TitleNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.Updated != nil {
		query := bson.M{}
		updatedQuery, err := o.Updated.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range updatedQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["updated."+k] = v
			}
		}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
	AssigneeId int8
	CreatedAt  int8
	Status     int8
	UpdatedAt  int8
}
