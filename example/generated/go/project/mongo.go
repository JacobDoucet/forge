package project

import (
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id          *primitive.ObjectID      `bson:"_id,omitempty"`
	Created     *actor_trace.MongoRecord `bson:"created,omitempty"`
	Description *string                  `bson:"description,omitempty"`
	Name        *string                  `bson:"name,omitempty"`
	OwnerId     *string                  `bson:"ownerId,omitempty"`
	Updated     *actor_trace.MongoRecord `bson:"updated,omitempty"`
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
	if r.Description != nil {
		elemdescription0 := r.Description
		m.Description = *elemdescription0
	}
	if r.Name != nil {
		elemname0 := r.Name
		m.Name = *elemname0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
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
	// id (Ref<Project>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
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
	// name (string) search options
	NameEq     *string
	NameNe     *string
	NameGt     *string
	NameGte    *string
	NameLt     *string
	NameLte    *string
	NameIn     *[]string
	NameNin    *[]string
	NameExists *bool
	NameLike   *string
	NameNlike  *string
	// ownerId (string) search options
	OwnerIdEq     *string
	OwnerIdNe     *string
	OwnerIdGt     *string
	OwnerIdGte    *string
	OwnerIdLt     *string
	OwnerIdLte    *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	OwnerIdLike   *string
	OwnerIdNlike  *string
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
	if o.NameEq != nil {
		query := bson.M{}
		query["name"] = o.NameEq
		and = append(and, query)
	}
	if o.NameNe != nil {
		query := bson.M{}
		query["name"] = bson.M{"$ne": o.NameNe}
		and = append(and, query)
	}
	if o.NameGt != nil {
		query := bson.M{}
		query["name"] = bson.M{"$gt": o.NameGt}
		and = append(and, query)
	}
	if o.NameGte != nil {
		query := bson.M{}
		query["name"] = bson.M{"$gte": o.NameGte}
		and = append(and, query)
	}
	if o.NameLt != nil {
		query := bson.M{}
		query["name"] = bson.M{"$lt": o.NameLt}
		and = append(and, query)
	}
	if o.NameLte != nil {
		query := bson.M{}
		query["name"] = bson.M{"$lte": o.NameLte}
		and = append(and, query)
	}
	if o.NameIn != nil {
		query := bson.M{}
		query["name"] = bson.M{"$in": o.NameIn}
		and = append(and, query)
	}
	if o.NameNin != nil {
		query := bson.M{}
		query["name"] = bson.M{"$nin": o.NameNin}
		and = append(and, query)
	}
	if o.NameExists != nil {
		query := bson.M{}
		query["name"] = bson.M{"$exists": *o.NameExists}
		and = append(and, query)
	}
	if o.NameLike != nil {
		query := bson.M{}
		query["name"] = bson.M{"$regex": o.NameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.NameNlike != nil {
		query := bson.M{}
		query["name"] = bson.M{"$not": bson.M{"$regex": o.NameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.OwnerIdEq != nil {
		query := bson.M{}
		query["ownerId"] = o.OwnerIdEq
		and = append(and, query)
	}
	if o.OwnerIdNe != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$ne": o.OwnerIdNe}
		and = append(and, query)
	}
	if o.OwnerIdGt != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$gt": o.OwnerIdGt}
		and = append(and, query)
	}
	if o.OwnerIdGte != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$gte": o.OwnerIdGte}
		and = append(and, query)
	}
	if o.OwnerIdLt != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$lt": o.OwnerIdLt}
		and = append(and, query)
	}
	if o.OwnerIdLte != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$lte": o.OwnerIdLte}
		and = append(and, query)
	}
	if o.OwnerIdIn != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$in": o.OwnerIdIn}
		and = append(and, query)
	}
	if o.OwnerIdNin != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$nin": o.OwnerIdNin}
		and = append(and, query)
	}
	if o.OwnerIdExists != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$exists": *o.OwnerIdExists}
		and = append(and, query)
	}
	if o.OwnerIdLike != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$regex": o.OwnerIdLike, "$options": "i"}
		and = append(and, query)
	}
	if o.OwnerIdNlike != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$not": bson.M{"$regex": o.OwnerIdNlike, "$options": "i"}}
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
	CreatedAt int8
	UpdatedAt int8
}
