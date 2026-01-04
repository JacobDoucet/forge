package task

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_task_priority"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_task_status"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/task_comment"
	"time"
)

type HTTPRecord struct {
	Id          *string                    `json:"id,omitempty"`
	AssigneeId  *string                    `json:"assigneeId,omitempty"`
	Comments    *[]task_comment.HTTPRecord `json:"comments,omitempty"`
	Created     *actor_trace.HTTPRecord    `json:"created,omitempty"`
	Description *string                    `json:"description,omitempty"`
	DueDate     *time.Time                 `json:"dueDate,omitempty"`
	Priority    *enum_task_priority.Value  `json:"priority,omitempty"`
	Status      *enum_task_status.Value    `json:"status,omitempty"`
	Tags        *[]string                  `json:"tags,omitempty"`
	Title       *string                    `json:"title,omitempty"`
	Updated     *actor_trace.HTTPRecord    `json:"updated,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.AssigneeId != nil {
		p.AssigneeId = true
	}
	if r.Comments != nil {
		p.Comments = true
		p.CommentsFields = task_comment.NewProjection(true)
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Description != nil {
		p.Description = true
	}
	if r.DueDate != nil {
		p.DueDate = true
	}
	if r.Priority != nil {
		p.Priority = true
	}
	if r.Status != nil {
		p.Status = true
	}
	if r.Tags != nil {
		p.Tags = true
	}
	if r.Title != nil {
		p.Title = true
	}
	if r.Updated != nil {
		p.Updated = true
		p.UpdatedFields = actor_trace.NewProjection(true)
	}
	return p, nil
}

type HTTPSelectByIdQuery struct {
	Id string `json:"id"`
}

type HTTPWhereClause struct {
	// id (Ref<Task>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// assigneeId (string) search options
	AssigneeIdEq     *string   `json:"assigneeIdEq,omitempty"`
	AssigneeIdNe     *string   `json:"assigneeIdNe,omitempty"`
	AssigneeIdGt     *string   `json:"assigneeIdGt,omitempty"`
	AssigneeIdGte    *string   `json:"assigneeIdGte,omitempty"`
	AssigneeIdLt     *string   `json:"assigneeIdLt,omitempty"`
	AssigneeIdLte    *string   `json:"assigneeIdLte,omitempty"`
	AssigneeIdIn     *[]string `json:"assigneeIdIn,omitempty"`
	AssigneeIdNin    *[]string `json:"assigneeIdNin,omitempty"`
	AssigneeIdExists *bool     `json:"assigneeIdExists,omitempty"`
	AssigneeIdLike   *string   `json:"assigneeIdLike,omitempty"`
	AssigneeIdNlike  *string   `json:"assigneeIdNlike,omitempty"`
	// comments (List<TaskComment>) search options
	Comments      *task_comment.HTTPWhereClause `json:"comments,omitempty"`
	CommentsEmpty *bool                         `json:"commentsEmpty,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// description (string) search options
	DescriptionEq     *string   `json:"descriptionEq,omitempty"`
	DescriptionNe     *string   `json:"descriptionNe,omitempty"`
	DescriptionGt     *string   `json:"descriptionGt,omitempty"`
	DescriptionGte    *string   `json:"descriptionGte,omitempty"`
	DescriptionLt     *string   `json:"descriptionLt,omitempty"`
	DescriptionLte    *string   `json:"descriptionLte,omitempty"`
	DescriptionIn     *[]string `json:"descriptionIn,omitempty"`
	DescriptionNin    *[]string `json:"descriptionNin,omitempty"`
	DescriptionExists *bool     `json:"descriptionExists,omitempty"`
	DescriptionLike   *string   `json:"descriptionLike,omitempty"`
	DescriptionNlike  *string   `json:"descriptionNlike,omitempty"`
	// dueDate (timestamp) search options
	DueDateEq     *time.Time   `json:"dueDateEq,omitempty"`
	DueDateNe     *time.Time   `json:"dueDateNe,omitempty"`
	DueDateGt     *time.Time   `json:"dueDateGt,omitempty"`
	DueDateGte    *time.Time   `json:"dueDateGte,omitempty"`
	DueDateLt     *time.Time   `json:"dueDateLt,omitempty"`
	DueDateLte    *time.Time   `json:"dueDateLte,omitempty"`
	DueDateIn     *[]time.Time `json:"dueDateIn,omitempty"`
	DueDateNin    *[]time.Time `json:"dueDateNin,omitempty"`
	DueDateExists *bool        `json:"dueDateExists,omitempty"`
	// priority (TaskPriority) search options
	PriorityEq     *enum_task_priority.Value   `json:"priorityEq,omitempty"`
	PriorityNe     *enum_task_priority.Value   `json:"priorityNe,omitempty"`
	PriorityGt     *enum_task_priority.Value   `json:"priorityGt,omitempty"`
	PriorityGte    *enum_task_priority.Value   `json:"priorityGte,omitempty"`
	PriorityLt     *enum_task_priority.Value   `json:"priorityLt,omitempty"`
	PriorityLte    *enum_task_priority.Value   `json:"priorityLte,omitempty"`
	PriorityIn     *[]enum_task_priority.Value `json:"priorityIn,omitempty"`
	PriorityNin    *[]enum_task_priority.Value `json:"priorityNin,omitempty"`
	PriorityExists *bool                       `json:"priorityExists,omitempty"`
	// status (TaskStatus) search options
	StatusEq     *enum_task_status.Value   `json:"statusEq,omitempty"`
	StatusNe     *enum_task_status.Value   `json:"statusNe,omitempty"`
	StatusGt     *enum_task_status.Value   `json:"statusGt,omitempty"`
	StatusGte    *enum_task_status.Value   `json:"statusGte,omitempty"`
	StatusLt     *enum_task_status.Value   `json:"statusLt,omitempty"`
	StatusLte    *enum_task_status.Value   `json:"statusLte,omitempty"`
	StatusIn     *[]enum_task_status.Value `json:"statusIn,omitempty"`
	StatusNin    *[]enum_task_status.Value `json:"statusNin,omitempty"`
	StatusExists *bool                     `json:"statusExists,omitempty"`
	// tags (List<string>) search options
	TagsEq     *string   `json:"tagsEq,omitempty"`
	TagsNe     *string   `json:"tagsNe,omitempty"`
	TagsGt     *string   `json:"tagsGt,omitempty"`
	TagsGte    *string   `json:"tagsGte,omitempty"`
	TagsLt     *string   `json:"tagsLt,omitempty"`
	TagsLte    *string   `json:"tagsLte,omitempty"`
	TagsIn     *[]string `json:"tagsIn,omitempty"`
	TagsNin    *[]string `json:"tagsNin,omitempty"`
	TagsExists *bool     `json:"tagsExists,omitempty"`
	TagsLike   *string   `json:"tagsLike,omitempty"`
	TagsNlike  *string   `json:"tagsNlike,omitempty"`
	TagsEmpty  *bool     `json:"tagsEmpty,omitempty"`
	// title (string) search options
	TitleEq     *string   `json:"titleEq,omitempty"`
	TitleNe     *string   `json:"titleNe,omitempty"`
	TitleGt     *string   `json:"titleGt,omitempty"`
	TitleGte    *string   `json:"titleGte,omitempty"`
	TitleLt     *string   `json:"titleLt,omitempty"`
	TitleLte    *string   `json:"titleLte,omitempty"`
	TitleIn     *[]string `json:"titleIn,omitempty"`
	TitleNin    *[]string `json:"titleNin,omitempty"`
	TitleExists *bool     `json:"titleExists,omitempty"`
	TitleLike   *string   `json:"titleLike,omitempty"`
	TitleNlike  *string   `json:"titleNlike,omitempty"`
	// updated (ActorTrace) search options
	Updated *actor_trace.HTTPWhereClause `json:"updated,omitempty"`
}

func (o HTTPSelectByIdQuery) ToSelectByIdQuery() (SelectByIdQuery, error) {
	to := SelectByIdQuery{}
	elemid0 := o.Id
	to.Id = elemid0
	return to, nil
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
	if o.IdEq != nil {
		elemidEq0 := o.IdEq
		to.IdEq = elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]string, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1 := oidIn0
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]string, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1 := oidNin0
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
		elemcomments0, err := o.Comments.ToWhereClause()
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
		elemcreated0, err := o.Created.ToWhereClause()
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
		elemupdated0, err := o.Updated.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	return to, nil
}

type HTTPSortParams struct {
	AssigneeId *int8 `json:"assigneeId,omitempty"`
	CreatedAt  *int8 `json:"createdAt,omitempty"`
	Status     *int8 `json:"status,omitempty"`
	UpdatedAt  *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.AssigneeId != nil {
		to.AssigneeId = *s.AssigneeId
	}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.Status != nil {
		to.Status = *s.Status
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
