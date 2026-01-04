package actor_trace

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MongoRecord struct {
	ActorId   *string    `bson:"actorId,omitempty"`
	ActorName *string    `bson:"actorName,omitempty"`
	ActorType *string    `bson:"actorType,omitempty"`
	At        *time.Time `bson:"at,omitempty"`
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.ActorId != nil {
		elemactorId0 := r.ActorId
		m.ActorId = *elemactorId0
	}
	if r.ActorName != nil {
		elemactorName0 := r.ActorName
		m.ActorName = *elemactorName0
	}
	if r.ActorType != nil {
		elemactorType0 := r.ActorType
		m.ActorType = *elemactorType0
	}
	if r.At != nil {
		elemat0 := r.At
		m.At = *elemat0
	}
	return m, nil
}

type MongoWhereClause struct {
	// actorId (string) search options
	ActorIdEq     *string
	ActorIdNe     *string
	ActorIdGt     *string
	ActorIdGte    *string
	ActorIdLt     *string
	ActorIdLte    *string
	ActorIdIn     *[]string
	ActorIdNin    *[]string
	ActorIdExists *bool
	ActorIdLike   *string
	ActorIdNlike  *string
	// actorName (string) search options
	ActorNameEq     *string
	ActorNameNe     *string
	ActorNameGt     *string
	ActorNameGte    *string
	ActorNameLt     *string
	ActorNameLte    *string
	ActorNameIn     *[]string
	ActorNameNin    *[]string
	ActorNameExists *bool
	ActorNameLike   *string
	ActorNameNlike  *string
	// actorType (string) search options
	ActorTypeEq     *string
	ActorTypeNe     *string
	ActorTypeGt     *string
	ActorTypeGte    *string
	ActorTypeLt     *string
	ActorTypeLte    *string
	ActorTypeIn     *[]string
	ActorTypeNin    *[]string
	ActorTypeExists *bool
	ActorTypeLike   *string
	ActorTypeNlike  *string
	// at (timestamp) search options
	AtEq     *time.Time
	AtNe     *time.Time
	AtGt     *time.Time
	AtGte    *time.Time
	AtLt     *time.Time
	AtLte    *time.Time
	AtIn     *[]time.Time
	AtNin    *[]time.Time
	AtExists *bool
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
	if o.ActorIdEq != nil {
		query := bson.M{}
		query["actorId"] = o.ActorIdEq
		and = append(and, query)
	}
	if o.ActorIdNe != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$ne": o.ActorIdNe}
		and = append(and, query)
	}
	if o.ActorIdGt != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$gt": o.ActorIdGt}
		and = append(and, query)
	}
	if o.ActorIdGte != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$gte": o.ActorIdGte}
		and = append(and, query)
	}
	if o.ActorIdLt != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$lt": o.ActorIdLt}
		and = append(and, query)
	}
	if o.ActorIdLte != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$lte": o.ActorIdLte}
		and = append(and, query)
	}
	if o.ActorIdIn != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$in": o.ActorIdIn}
		and = append(and, query)
	}
	if o.ActorIdNin != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$nin": o.ActorIdNin}
		and = append(and, query)
	}
	if o.ActorIdExists != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$exists": *o.ActorIdExists}
		and = append(and, query)
	}
	if o.ActorIdLike != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$regex": o.ActorIdLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ActorIdNlike != nil {
		query := bson.M{}
		query["actorId"] = bson.M{"$not": bson.M{"$regex": o.ActorIdNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ActorNameEq != nil {
		query := bson.M{}
		query["actorName"] = o.ActorNameEq
		and = append(and, query)
	}
	if o.ActorNameNe != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$ne": o.ActorNameNe}
		and = append(and, query)
	}
	if o.ActorNameGt != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$gt": o.ActorNameGt}
		and = append(and, query)
	}
	if o.ActorNameGte != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$gte": o.ActorNameGte}
		and = append(and, query)
	}
	if o.ActorNameLt != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$lt": o.ActorNameLt}
		and = append(and, query)
	}
	if o.ActorNameLte != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$lte": o.ActorNameLte}
		and = append(and, query)
	}
	if o.ActorNameIn != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$in": o.ActorNameIn}
		and = append(and, query)
	}
	if o.ActorNameNin != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$nin": o.ActorNameNin}
		and = append(and, query)
	}
	if o.ActorNameExists != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$exists": *o.ActorNameExists}
		and = append(and, query)
	}
	if o.ActorNameLike != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$regex": o.ActorNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ActorNameNlike != nil {
		query := bson.M{}
		query["actorName"] = bson.M{"$not": bson.M{"$regex": o.ActorNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ActorTypeEq != nil {
		query := bson.M{}
		query["actorType"] = o.ActorTypeEq
		and = append(and, query)
	}
	if o.ActorTypeNe != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$ne": o.ActorTypeNe}
		and = append(and, query)
	}
	if o.ActorTypeGt != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$gt": o.ActorTypeGt}
		and = append(and, query)
	}
	if o.ActorTypeGte != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$gte": o.ActorTypeGte}
		and = append(and, query)
	}
	if o.ActorTypeLt != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$lt": o.ActorTypeLt}
		and = append(and, query)
	}
	if o.ActorTypeLte != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$lte": o.ActorTypeLte}
		and = append(and, query)
	}
	if o.ActorTypeIn != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$in": o.ActorTypeIn}
		and = append(and, query)
	}
	if o.ActorTypeNin != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$nin": o.ActorTypeNin}
		and = append(and, query)
	}
	if o.ActorTypeExists != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$exists": *o.ActorTypeExists}
		and = append(and, query)
	}
	if o.ActorTypeLike != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$regex": o.ActorTypeLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ActorTypeNlike != nil {
		query := bson.M{}
		query["actorType"] = bson.M{"$not": bson.M{"$regex": o.ActorTypeNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.AtEq != nil {
		query := bson.M{}
		query["at"] = o.AtEq
		and = append(and, query)
	}
	if o.AtNe != nil {
		query := bson.M{}
		query["at"] = bson.M{"$ne": o.AtNe}
		and = append(and, query)
	}
	if o.AtGt != nil {
		query := bson.M{}
		query["at"] = bson.M{"$gt": o.AtGt}
		and = append(and, query)
	}
	if o.AtGte != nil {
		query := bson.M{}
		query["at"] = bson.M{"$gte": o.AtGte}
		and = append(and, query)
	}
	if o.AtLt != nil {
		query := bson.M{}
		query["at"] = bson.M{"$lt": o.AtLt}
		and = append(and, query)
	}
	if o.AtLte != nil {
		query := bson.M{}
		query["at"] = bson.M{"$lte": o.AtLte}
		and = append(and, query)
	}
	if o.AtIn != nil {
		query := bson.M{}
		query["at"] = bson.M{"$in": o.AtIn}
		and = append(and, query)
	}
	if o.AtNin != nil {
		query := bson.M{}
		query["at"] = bson.M{"$nin": o.AtNin}
		and = append(and, query)
	}
	if o.AtExists != nil {
		query := bson.M{}
		query["at"] = bson.M{"$exists": *o.AtExists}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
}
