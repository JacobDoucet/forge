package event

import (
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
	"github.com/JacobDoucet/forge/example/generated/go/enum_event_type"
	"github.com/JacobDoucet/forge/example/generated/go/event_subject"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id       *primitive.ObjectID          `bson:"_id,omitempty"`
	Created  *actor_trace.MongoRecord     `bson:"created,omitempty"`
	Subjects *[]event_subject.MongoRecord `bson:"subjects,omitempty"`
	Type     *enum_event_type.Value       `bson:"type,omitempty"`
	Updated  *actor_trace.MongoRecord     `bson:"updated,omitempty"`
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

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}

type MongoWhereClause struct {
	// id (Ref<Event>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// subjects (List<EventSubject>) search options
	Subjects      *event_subject.MongoWhereClause
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
	if o.Subjects != nil {
		query := bson.M{}
		subjectsQuery, err := o.Subjects.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range subjectsQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["subjects."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.SubjectsEmpty != nil {
		query := bson.M{}
		if *o.SubjectsEmpty {
			query["$or"] = bson.A{
				bson.M{"subjects": nil},
				bson.M{"subjects": bson.A{}},
				bson.M{"subjects": bson.M{"$exists": false}},
			}
		} else {
			query["subjects"] = bson.M{
				"$ne":     nil,
				"$not":    bson.M{"$size": 0},
				"$exists": true,
			}
		}
		and = append(and, query)
	}
	if o.TypeEq != nil {
		query := bson.M{}
		query["type"] = o.TypeEq
		and = append(and, query)
	}
	if o.TypeNe != nil {
		query := bson.M{}
		query["type"] = bson.M{"$ne": o.TypeNe}
		and = append(and, query)
	}
	if o.TypeGt != nil {
		query := bson.M{}
		query["type"] = bson.M{"$gt": o.TypeGt}
		and = append(and, query)
	}
	if o.TypeGte != nil {
		query := bson.M{}
		query["type"] = bson.M{"$gte": o.TypeGte}
		and = append(and, query)
	}
	if o.TypeLt != nil {
		query := bson.M{}
		query["type"] = bson.M{"$lt": o.TypeLt}
		and = append(and, query)
	}
	if o.TypeLte != nil {
		query := bson.M{}
		query["type"] = bson.M{"$lte": o.TypeLte}
		and = append(and, query)
	}
	if o.TypeIn != nil {
		query := bson.M{}
		query["type"] = bson.M{"$in": o.TypeIn}
		and = append(and, query)
	}
	if o.TypeNin != nil {
		query := bson.M{}
		query["type"] = bson.M{"$nin": o.TypeNin}
		and = append(and, query)
	}
	if o.TypeExists != nil {
		query := bson.M{}
		query["type"] = bson.M{"$exists": *o.TypeExists}
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
	CreatedAt         int8
	SubjectsSubjectId int8
	Type              int8
	UpdatedAt         int8
}
