package event

import (
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_event_type"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/event_subject"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id       string
	Created  actor_trace.Model
	Subjects []event_subject.Model
	Type     enum_event_type.Value
	Updated  actor_trace.Model
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
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Subjects {
		elemsubjects0 := make([]event_subject.MongoRecord, 0)
		for _, msubjects0 := range m.Subjects {
			elemsubjects1, err := msubjects0.ToMongoRecord(projection.SubjectsFields)
			if err != nil {
				return r, err
			}
			elemsubjects0 = append(elemsubjects0, elemsubjects1)
		}
		r.Subjects = &elemsubjects0
	}
	if projection.Type {
		elemtype0 := m.Type
		r.Type = &elemtype0
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
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Subjects {
		elemsubjects0 := make([]event_subject.HTTPRecord, 0)
		for _, msubjects0 := range m.Subjects {
			elemsubjects1, err := msubjects0.ToHTTPRecord(projection.SubjectsFields)
			if err != nil {
				return r, err
			}
			elemsubjects0 = append(elemsubjects0, elemsubjects1)
		}
		r.Subjects = &elemsubjects0
	}
	if projection.Type {
		elemtype0 := m.Type
		r.Type = &elemtype0
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
	// id (Ref<Event>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
	// subjects (List<EventSubject>) search options
	Subjects      *event_subject.WhereClause
	SubjectsEmpty *bool
	// type (EventType) search options
	TypeEq     *enum_event_type.Value
	TypeNe     *enum_event_type.Value
	TypeGt     *enum_event_type.Value
	TypeGte    *enum_event_type.Value
	TypeLt     *enum_event_type.Value
	TypeLte    *enum_event_type.Value
	TypeIn     *[]enum_event_type.Value
	TypeNin    *[]enum_event_type.Value
	TypeExists *bool
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
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.Subjects != nil {
		elemsubjects0, err := o.Subjects.ToMongoWhereClause()
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
		elemupdated0, err := o.Updated.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	return to, nil
}

type SortParams struct {
	CreatedAt         int8
	SubjectsSubjectId int8
	Type              int8
	UpdatedAt         int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.SubjectsSubjectId = s.SubjectsSubjectId
	to.Type = s.Type
	to.UpdatedAt = s.UpdatedAt
	return to
}
