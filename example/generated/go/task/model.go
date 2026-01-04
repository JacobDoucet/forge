package task

import (
	"errors"
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
	"github.com/JacobDoucet/forge/example/generated/go/enum_task_priority"
	"github.com/JacobDoucet/forge/example/generated/go/enum_task_status"
	"github.com/JacobDoucet/forge/example/generated/go/task_comment"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	Id          string
	AssigneeId  string
	Comments    []task_comment.Model
	Created     actor_trace.Model
	Description string
	DueDate     time.Time
	Priority    enum_task_priority.Value
	Status      enum_task_status.Value
	Tags        []string
	Title       string
	Updated     actor_trace.Model
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if m.Id != "" {
		elemid0, err := primitive.ObjectIDFromHex(m.Id)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.Id"), err)
		}
		r.Id = &elemid0
	}
	if projection.AssigneeId {
		elemassigneeId0 := m.AssigneeId
		r.AssigneeId = &elemassigneeId0
	}
	if projection.Comments {
		elemcomments0 := make([]task_comment.MongoRecord, 0)
		for _, mcomments0 := range m.Comments {
			elemcomments1, err := mcomments0.ToMongoRecord(projection.CommentsFields)
			if err != nil {
				return r, err
			}
			elemcomments0 = append(elemcomments0, elemcomments1)
		}
		r.Comments = &elemcomments0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Description {
		elemdescription0 := m.Description
		r.Description = &elemdescription0
	}
	if projection.DueDate {
		elemdueDate0 := m.DueDate
		r.DueDate = &elemdueDate0
	}
	if projection.Priority {
		elempriority0 := m.Priority
		r.Priority = &elempriority0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
	}
	if projection.Tags {
		elemtags0 := make([]string, 0)
		for _, mtags0 := range m.Tags {
			elemtags1 := mtags0
			elemtags0 = append(elemtags0, elemtags1)
		}
		r.Tags = &elemtags0
	}
	if projection.Title {
		elemtitle0 := m.Title
		r.Title = &elemtitle0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToMongoRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if m.Id != "" {
		elemid0 := m.Id
		r.Id = &elemid0
	}
	if projection.AssigneeId {
		elemassigneeId0 := m.AssigneeId
		r.AssigneeId = &elemassigneeId0
	}
	if projection.Comments {
		elemcomments0 := make([]task_comment.HTTPRecord, 0)
		for _, mcomments0 := range m.Comments {
			elemcomments1, err := mcomments0.ToHTTPRecord(projection.CommentsFields)
			if err != nil {
				return r, err
			}
			elemcomments0 = append(elemcomments0, elemcomments1)
		}
		r.Comments = &elemcomments0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Description {
		elemdescription0 := m.Description
		r.Description = &elemdescription0
	}
	if projection.DueDate {
		elemdueDate0 := m.DueDate
		r.DueDate = &elemdueDate0
	}
	if projection.Priority {
		elempriority0 := m.Priority
		r.Priority = &elempriority0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
	}
	if projection.Tags {
		elemtags0 := make([]string, 0)
		for _, mtags0 := range m.Tags {
			elemtags1 := mtags0
			elemtags0 = append(elemtags0, elemtags1)
		}
		r.Tags = &elemtags0
	}
	if projection.Title {
		elemtitle0 := m.Title
		r.Title = &elemtitle0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToHTTPRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	return r, nil
}

type SelectByIdQuery struct {
	Id string
}

type WhereClause struct {
	// id (Ref<Task>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
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
	Comments      *task_comment.WhereClause
	CommentsEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
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
	Updated *actor_trace.WhereClause
}

func (o SelectByIdQuery) ToMongoSelectByIdQuery() (MongoSelectByIdQuery, error) {
	to := MongoSelectByIdQuery{}
	elemid0, err := primitive.ObjectIDFromHex(o.Id)
	if err != nil {
		return to, errors.Join(errors.New("invalid o.Id"), err)
	}
	to.Id = elemid0
	return to, nil
}

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.IdEq != nil {
		elemidEq0, err := primitive.ObjectIDFromHex(*o.IdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.IdEq"), err)
		}
		to.IdEq = &elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]primitive.ObjectID, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1, err := primitive.ObjectIDFromHex(oidIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidIn0"), err)
			}
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]primitive.ObjectID, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1, err := primitive.ObjectIDFromHex(oidNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidNin0"), err)
			}
			elemidNin0 = append(elemidNin0, elemidNin1)
		}
		to.IdNin = &elemidNin0
	}
	if o.IdExists != nil {
		elemidExists0 := o.IdExists
		to.IdExists = elemidExists0
	}
	if o.AssigneeIdEq != nil {
		elemassigneeIdEq0 := o.AssigneeIdEq
		to.AssigneeIdEq = elemassigneeIdEq0
	}
	if o.AssigneeIdNe != nil {
		elemassigneeIdNe0 := o.AssigneeIdNe
		to.AssigneeIdNe = elemassigneeIdNe0
	}
	if o.AssigneeIdGt != nil {
		elemassigneeIdGt0 := o.AssigneeIdGt
		to.AssigneeIdGt = elemassigneeIdGt0
	}
	if o.AssigneeIdGte != nil {
		elemassigneeIdGte0 := o.AssigneeIdGte
		to.AssigneeIdGte = elemassigneeIdGte0
	}
	if o.AssigneeIdLt != nil {
		elemassigneeIdLt0 := o.AssigneeIdLt
		to.AssigneeIdLt = elemassigneeIdLt0
	}
	if o.AssigneeIdLte != nil {
		elemassigneeIdLte0 := o.AssigneeIdLte
		to.AssigneeIdLte = elemassigneeIdLte0
	}
	if o.AssigneeIdIn != nil {
		elemassigneeIdIn0 := make([]string, 0)
		for _, oassigneeIdIn0 := range *o.AssigneeIdIn {
			elemassigneeIdIn1 := oassigneeIdIn0
			elemassigneeIdIn0 = append(elemassigneeIdIn0, elemassigneeIdIn1)
		}
		to.AssigneeIdIn = &elemassigneeIdIn0
	}
	if o.AssigneeIdNin != nil {
		elemassigneeIdNin0 := make([]string, 0)
		for _, oassigneeIdNin0 := range *o.AssigneeIdNin {
			elemassigneeIdNin1 := oassigneeIdNin0
			elemassigneeIdNin0 = append(elemassigneeIdNin0, elemassigneeIdNin1)
		}
		to.AssigneeIdNin = &elemassigneeIdNin0
	}
	if o.AssigneeIdExists != nil {
		elemassigneeIdExists0 := o.AssigneeIdExists
		to.AssigneeIdExists = elemassigneeIdExists0
	}
	if o.AssigneeIdLike != nil {
		elemassigneeIdLike0 := o.AssigneeIdLike
		to.AssigneeIdLike = elemassigneeIdLike0
	}
	if o.AssigneeIdNlike != nil {
		elemassigneeIdNlike0 := o.AssigneeIdNlike
		to.AssigneeIdNlike = elemassigneeIdNlike0
	}
	if o.Comments != nil {
		elemcomments0, err := o.Comments.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Comments = &elemcomments0
	}
	if o.CommentsEmpty != nil {
		elemcommentsEmpty0 := o.CommentsEmpty
		to.CommentsEmpty = elemcommentsEmpty0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.DescriptionEq != nil {
		elemdescriptionEq0 := o.DescriptionEq
		to.DescriptionEq = elemdescriptionEq0
	}
	if o.DescriptionNe != nil {
		elemdescriptionNe0 := o.DescriptionNe
		to.DescriptionNe = elemdescriptionNe0
	}
	if o.DescriptionGt != nil {
		elemdescriptionGt0 := o.DescriptionGt
		to.DescriptionGt = elemdescriptionGt0
	}
	if o.DescriptionGte != nil {
		elemdescriptionGte0 := o.DescriptionGte
		to.DescriptionGte = elemdescriptionGte0
	}
	if o.DescriptionLt != nil {
		elemdescriptionLt0 := o.DescriptionLt
		to.DescriptionLt = elemdescriptionLt0
	}
	if o.DescriptionLte != nil {
		elemdescriptionLte0 := o.DescriptionLte
		to.DescriptionLte = elemdescriptionLte0
	}
	if o.DescriptionIn != nil {
		elemdescriptionIn0 := make([]string, 0)
		for _, odescriptionIn0 := range *o.DescriptionIn {
			elemdescriptionIn1 := odescriptionIn0
			elemdescriptionIn0 = append(elemdescriptionIn0, elemdescriptionIn1)
		}
		to.DescriptionIn = &elemdescriptionIn0
	}
	if o.DescriptionNin != nil {
		elemdescriptionNin0 := make([]string, 0)
		for _, odescriptionNin0 := range *o.DescriptionNin {
			elemdescriptionNin1 := odescriptionNin0
			elemdescriptionNin0 = append(elemdescriptionNin0, elemdescriptionNin1)
		}
		to.DescriptionNin = &elemdescriptionNin0
	}
	if o.DescriptionExists != nil {
		elemdescriptionExists0 := o.DescriptionExists
		to.DescriptionExists = elemdescriptionExists0
	}
	if o.DescriptionLike != nil {
		elemdescriptionLike0 := o.DescriptionLike
		to.DescriptionLike = elemdescriptionLike0
	}
	if o.DescriptionNlike != nil {
		elemdescriptionNlike0 := o.DescriptionNlike
		to.DescriptionNlike = elemdescriptionNlike0
	}
	if o.DueDateEq != nil {
		elemdueDateEq0 := o.DueDateEq
		to.DueDateEq = elemdueDateEq0
	}
	if o.DueDateNe != nil {
		elemdueDateNe0 := o.DueDateNe
		to.DueDateNe = elemdueDateNe0
	}
	if o.DueDateGt != nil {
		elemdueDateGt0 := o.DueDateGt
		to.DueDateGt = elemdueDateGt0
	}
	if o.DueDateGte != nil {
		elemdueDateGte0 := o.DueDateGte
		to.DueDateGte = elemdueDateGte0
	}
	if o.DueDateLt != nil {
		elemdueDateLt0 := o.DueDateLt
		to.DueDateLt = elemdueDateLt0
	}
	if o.DueDateLte != nil {
		elemdueDateLte0 := o.DueDateLte
		to.DueDateLte = elemdueDateLte0
	}
	if o.DueDateIn != nil {
		elemdueDateIn0 := make([]time.Time, 0)
		for _, odueDateIn0 := range *o.DueDateIn {
			elemdueDateIn1 := odueDateIn0
			elemdueDateIn0 = append(elemdueDateIn0, elemdueDateIn1)
		}
		to.DueDateIn = &elemdueDateIn0
	}
	if o.DueDateNin != nil {
		elemdueDateNin0 := make([]time.Time, 0)
		for _, odueDateNin0 := range *o.DueDateNin {
			elemdueDateNin1 := odueDateNin0
			elemdueDateNin0 = append(elemdueDateNin0, elemdueDateNin1)
		}
		to.DueDateNin = &elemdueDateNin0
	}
	if o.DueDateExists != nil {
		elemdueDateExists0 := o.DueDateExists
		to.DueDateExists = elemdueDateExists0
	}
	if o.PriorityEq != nil {
		elempriorityEq0 := o.PriorityEq
		to.PriorityEq = elempriorityEq0
	}
	if o.PriorityNe != nil {
		elempriorityNe0 := o.PriorityNe
		to.PriorityNe = elempriorityNe0
	}
	if o.PriorityGt != nil {
		elempriorityGt0 := o.PriorityGt
		to.PriorityGt = elempriorityGt0
	}
	if o.PriorityGte != nil {
		elempriorityGte0 := o.PriorityGte
		to.PriorityGte = elempriorityGte0
	}
	if o.PriorityLt != nil {
		elempriorityLt0 := o.PriorityLt
		to.PriorityLt = elempriorityLt0
	}
	if o.PriorityLte != nil {
		elempriorityLte0 := o.PriorityLte
		to.PriorityLte = elempriorityLte0
	}
	if o.PriorityIn != nil {
		elempriorityIn0 := make([]enum_task_priority.Value, 0)
		for _, opriorityIn0 := range *o.PriorityIn {
			elempriorityIn1 := opriorityIn0
			elempriorityIn0 = append(elempriorityIn0, elempriorityIn1)
		}
		to.PriorityIn = &elempriorityIn0
	}
	if o.PriorityNin != nil {
		elempriorityNin0 := make([]enum_task_priority.Value, 0)
		for _, opriorityNin0 := range *o.PriorityNin {
			elempriorityNin1 := opriorityNin0
			elempriorityNin0 = append(elempriorityNin0, elempriorityNin1)
		}
		to.PriorityNin = &elempriorityNin0
	}
	if o.PriorityExists != nil {
		elempriorityExists0 := o.PriorityExists
		to.PriorityExists = elempriorityExists0
	}
	if o.StatusEq != nil {
		elemstatusEq0 := o.StatusEq
		to.StatusEq = elemstatusEq0
	}
	if o.StatusNe != nil {
		elemstatusNe0 := o.StatusNe
		to.StatusNe = elemstatusNe0
	}
	if o.StatusGt != nil {
		elemstatusGt0 := o.StatusGt
		to.StatusGt = elemstatusGt0
	}
	if o.StatusGte != nil {
		elemstatusGte0 := o.StatusGte
		to.StatusGte = elemstatusGte0
	}
	if o.StatusLt != nil {
		elemstatusLt0 := o.StatusLt
		to.StatusLt = elemstatusLt0
	}
	if o.StatusLte != nil {
		elemstatusLte0 := o.StatusLte
		to.StatusLte = elemstatusLte0
	}
	if o.StatusIn != nil {
		elemstatusIn0 := make([]enum_task_status.Value, 0)
		for _, ostatusIn0 := range *o.StatusIn {
			elemstatusIn1 := ostatusIn0
			elemstatusIn0 = append(elemstatusIn0, elemstatusIn1)
		}
		to.StatusIn = &elemstatusIn0
	}
	if o.StatusNin != nil {
		elemstatusNin0 := make([]enum_task_status.Value, 0)
		for _, ostatusNin0 := range *o.StatusNin {
			elemstatusNin1 := ostatusNin0
			elemstatusNin0 = append(elemstatusNin0, elemstatusNin1)
		}
		to.StatusNin = &elemstatusNin0
	}
	if o.StatusExists != nil {
		elemstatusExists0 := o.StatusExists
		to.StatusExists = elemstatusExists0
	}
	if o.TagsEq != nil {
		elemtagsEq0 := o.TagsEq
		to.TagsEq = elemtagsEq0
	}
	if o.TagsNe != nil {
		elemtagsNe0 := o.TagsNe
		to.TagsNe = elemtagsNe0
	}
	if o.TagsGt != nil {
		elemtagsGt0 := o.TagsGt
		to.TagsGt = elemtagsGt0
	}
	if o.TagsGte != nil {
		elemtagsGte0 := o.TagsGte
		to.TagsGte = elemtagsGte0
	}
	if o.TagsLt != nil {
		elemtagsLt0 := o.TagsLt
		to.TagsLt = elemtagsLt0
	}
	if o.TagsLte != nil {
		elemtagsLte0 := o.TagsLte
		to.TagsLte = elemtagsLte0
	}
	if o.TagsIn != nil {
		elemtagsIn0 := make([]string, 0)
		for _, otagsIn0 := range *o.TagsIn {
			elemtagsIn1 := otagsIn0
			elemtagsIn0 = append(elemtagsIn0, elemtagsIn1)
		}
		to.TagsIn = &elemtagsIn0
	}
	if o.TagsNin != nil {
		elemtagsNin0 := make([]string, 0)
		for _, otagsNin0 := range *o.TagsNin {
			elemtagsNin1 := otagsNin0
			elemtagsNin0 = append(elemtagsNin0, elemtagsNin1)
		}
		to.TagsNin = &elemtagsNin0
	}
	if o.TagsExists != nil {
		elemtagsExists0 := o.TagsExists
		to.TagsExists = elemtagsExists0
	}
	if o.TagsLike != nil {
		elemtagsLike0 := o.TagsLike
		to.TagsLike = elemtagsLike0
	}
	if o.TagsNlike != nil {
		elemtagsNlike0 := o.TagsNlike
		to.TagsNlike = elemtagsNlike0
	}
	if o.TagsEmpty != nil {
		elemtagsEmpty0 := o.TagsEmpty
		to.TagsEmpty = elemtagsEmpty0
	}
	if o.TitleEq != nil {
		elemtitleEq0 := o.TitleEq
		to.TitleEq = elemtitleEq0
	}
	if o.TitleNe != nil {
		elemtitleNe0 := o.TitleNe
		to.TitleNe = elemtitleNe0
	}
	if o.TitleGt != nil {
		elemtitleGt0 := o.TitleGt
		to.TitleGt = elemtitleGt0
	}
	if o.TitleGte != nil {
		elemtitleGte0 := o.TitleGte
		to.TitleGte = elemtitleGte0
	}
	if o.TitleLt != nil {
		elemtitleLt0 := o.TitleLt
		to.TitleLt = elemtitleLt0
	}
	if o.TitleLte != nil {
		elemtitleLte0 := o.TitleLte
		to.TitleLte = elemtitleLte0
	}
	if o.TitleIn != nil {
		elemtitleIn0 := make([]string, 0)
		for _, otitleIn0 := range *o.TitleIn {
			elemtitleIn1 := otitleIn0
			elemtitleIn0 = append(elemtitleIn0, elemtitleIn1)
		}
		to.TitleIn = &elemtitleIn0
	}
	if o.TitleNin != nil {
		elemtitleNin0 := make([]string, 0)
		for _, otitleNin0 := range *o.TitleNin {
			elemtitleNin1 := otitleNin0
			elemtitleNin0 = append(elemtitleNin0, elemtitleNin1)
		}
		to.TitleNin = &elemtitleNin0
	}
	if o.TitleExists != nil {
		elemtitleExists0 := o.TitleExists
		to.TitleExists = elemtitleExists0
	}
	if o.TitleLike != nil {
		elemtitleLike0 := o.TitleLike
		to.TitleLike = elemtitleLike0
	}
	if o.TitleNlike != nil {
		elemtitleNlike0 := o.TitleNlike
		to.TitleNlike = elemtitleNlike0
	}
	if o.Updated != nil {
		elemupdated0, err := o.Updated.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	return to, nil
}

type SortParams struct {
	AssigneeId int8
	CreatedAt  int8
	Status     int8
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.AssigneeId = s.AssigneeId
	to.CreatedAt = s.CreatedAt
	to.Status = s.Status
	to.UpdatedAt = s.UpdatedAt
	return to
}
