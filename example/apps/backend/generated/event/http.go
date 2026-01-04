package event

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_event_type"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_subject"
)

type HTTPRecord struct {
	Id       *string                     `json:"id,omitempty"`
	Created  *actor_trace.HTTPRecord     `json:"created,omitempty"`
	Subjects *[]event_subject.HTTPRecord `json:"subjects,omitempty"`
	Type     *enum_event_type.Value      `json:"type,omitempty"`
	Updated  *actor_trace.HTTPRecord     `json:"updated,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Subjects != nil {
		elemsubjects0 := make([]event_subject.Model, 0)
		for _, rsubjects0 := range *r.Subjects {
			elemsubjects1, err := rsubjects0.ToModel()
			if err != nil {
				return m, err
			}
			elemsubjects0 = append(elemsubjects0, elemsubjects1)
		}
		m.Subjects = elemsubjects0
	}
	if r.Type != nil {
		elemtype0 := r.Type
		m.Type = *elemtype0
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
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Subjects != nil {
		p.Subjects = true
		p.SubjectsFields = event_subject.NewProjection(true)
	}
	if r.Type != nil {
		p.Type = true
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
	// id (Ref<Event>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// subjects (List<EventSubject>) search options
	Subjects      *event_subject.HTTPWhereClause `json:"subjects,omitempty"`
	SubjectsEmpty *bool                          `json:"subjectsEmpty,omitempty"`
	// type (EventType) search options
	TypeEq     *enum_event_type.Value   `json:"typeEq,omitempty"`
	TypeNe     *enum_event_type.Value   `json:"typeNe,omitempty"`
	TypeGt     *enum_event_type.Value   `json:"typeGt,omitempty"`
	TypeGte    *enum_event_type.Value   `json:"typeGte,omitempty"`
	TypeLt     *enum_event_type.Value   `json:"typeLt,omitempty"`
	TypeLte    *enum_event_type.Value   `json:"typeLte,omitempty"`
	TypeIn     *[]enum_event_type.Value `json:"typeIn,omitempty"`
	TypeNin    *[]enum_event_type.Value `json:"typeNin,omitempty"`
	TypeExists *bool                    `json:"typeExists,omitempty"`
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
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.Subjects != nil {
		elemsubjects0, err := o.Subjects.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Subjects = &elemsubjects0
	}
	if o.SubjectsEmpty != nil {
		elemsubjectsEmpty0 := o.SubjectsEmpty
		to.SubjectsEmpty = elemsubjectsEmpty0
	}
	if o.TypeEq != nil {
		elemtypeEq0 := o.TypeEq
		to.TypeEq = elemtypeEq0
	}
	if o.TypeNe != nil {
		elemtypeNe0 := o.TypeNe
		to.TypeNe = elemtypeNe0
	}
	if o.TypeGt != nil {
		elemtypeGt0 := o.TypeGt
		to.TypeGt = elemtypeGt0
	}
	if o.TypeGte != nil {
		elemtypeGte0 := o.TypeGte
		to.TypeGte = elemtypeGte0
	}
	if o.TypeLt != nil {
		elemtypeLt0 := o.TypeLt
		to.TypeLt = elemtypeLt0
	}
	if o.TypeLte != nil {
		elemtypeLte0 := o.TypeLte
		to.TypeLte = elemtypeLte0
	}
	if o.TypeIn != nil {
		elemtypeIn0 := make([]enum_event_type.Value, 0)
		for _, otypeIn0 := range *o.TypeIn {
			elemtypeIn1 := otypeIn0
			elemtypeIn0 = append(elemtypeIn0, elemtypeIn1)
		}
		to.TypeIn = &elemtypeIn0
	}
	if o.TypeNin != nil {
		elemtypeNin0 := make([]enum_event_type.Value, 0)
		for _, otypeNin0 := range *o.TypeNin {
			elemtypeNin1 := otypeNin0
			elemtypeNin0 = append(elemtypeNin0, elemtypeNin1)
		}
		to.TypeNin = &elemtypeNin0
	}
	if o.TypeExists != nil {
		elemtypeExists0 := o.TypeExists
		to.TypeExists = elemtypeExists0
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
	CreatedAt         *int8 `json:"createdAt,omitempty"`
	SubjectsSubjectId *int8 `json:"subjectsSubjectId,omitempty"`
	Type              *int8 `json:"type,omitempty"`
	UpdatedAt         *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.SubjectsSubjectId != nil {
		to.SubjectsSubjectId = *s.SubjectsSubjectId
	}
	if s.Type != nil {
		to.Type = *s.Type
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
