package event_subject

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_model"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRecord struct {
	SubjectId   *string           `bson:"subjectId,omitempty"`
	SubjectType *enum_model.Value `bson:"subjectType,omitempty"`
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.SubjectId != nil {
		elemsubjectId0 := r.SubjectId
		m.SubjectId = *elemsubjectId0
	}
	if r.SubjectType != nil {
		elemsubjectType0 := r.SubjectType
		m.SubjectType = *elemsubjectType0
	}
	return m, nil
}

type MongoWhereClause struct {
	// subjectId (string) search options
	SubjectIdEq     *string
	SubjectIdNe     *string
	SubjectIdGt     *string
	SubjectIdGte    *string
	SubjectIdLt     *string
	SubjectIdLte    *string
	SubjectIdIn     *[]string
	SubjectIdNin    *[]string
	SubjectIdExists *bool
	SubjectIdLike   *string
	SubjectIdNlike  *string
	// subjectType (Model) search options
	SubjectTypeEq     *enum_model.Value
	SubjectTypeNe     *enum_model.Value
	SubjectTypeGt     *enum_model.Value
	SubjectTypeGte    *enum_model.Value
	SubjectTypeLt     *enum_model.Value
	SubjectTypeLte    *enum_model.Value
	SubjectTypeIn     *[]enum_model.Value
	SubjectTypeNin    *[]enum_model.Value
	SubjectTypeExists *bool
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
	if o.SubjectIdEq != nil {
		query := bson.M{}
		query["subjectId"] = o.SubjectIdEq
		and = append(and, query)
	}
	if o.SubjectIdNe != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$ne": o.SubjectIdNe}
		and = append(and, query)
	}
	if o.SubjectIdGt != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$gt": o.SubjectIdGt}
		and = append(and, query)
	}
	if o.SubjectIdGte != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$gte": o.SubjectIdGte}
		and = append(and, query)
	}
	if o.SubjectIdLt != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$lt": o.SubjectIdLt}
		and = append(and, query)
	}
	if o.SubjectIdLte != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$lte": o.SubjectIdLte}
		and = append(and, query)
	}
	if o.SubjectIdIn != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$in": o.SubjectIdIn}
		and = append(and, query)
	}
	if o.SubjectIdNin != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$nin": o.SubjectIdNin}
		and = append(and, query)
	}
	if o.SubjectIdExists != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$exists": *o.SubjectIdExists}
		and = append(and, query)
	}
	if o.SubjectIdLike != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$regex": o.SubjectIdLike, "$options": "i"}
		and = append(and, query)
	}
	if o.SubjectIdNlike != nil {
		query := bson.M{}
		query["subjectId"] = bson.M{"$not": bson.M{"$regex": o.SubjectIdNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.SubjectTypeEq != nil {
		query := bson.M{}
		query["subjectType"] = o.SubjectTypeEq
		and = append(and, query)
	}
	if o.SubjectTypeNe != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$ne": o.SubjectTypeNe}
		and = append(and, query)
	}
	if o.SubjectTypeGt != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$gt": o.SubjectTypeGt}
		and = append(and, query)
	}
	if o.SubjectTypeGte != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$gte": o.SubjectTypeGte}
		and = append(and, query)
	}
	if o.SubjectTypeLt != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$lt": o.SubjectTypeLt}
		and = append(and, query)
	}
	if o.SubjectTypeLte != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$lte": o.SubjectTypeLte}
		and = append(and, query)
	}
	if o.SubjectTypeIn != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$in": o.SubjectTypeIn}
		and = append(and, query)
	}
	if o.SubjectTypeNin != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$nin": o.SubjectTypeNin}
		and = append(and, query)
	}
	if o.SubjectTypeExists != nil {
		query := bson.M{}
		query["subjectType"] = bson.M{"$exists": *o.SubjectTypeExists}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
}
