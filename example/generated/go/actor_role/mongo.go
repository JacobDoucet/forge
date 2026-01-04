package actor_role

import (
	"github.com/JacobDoucet/forge/example/generated/go/enum_role"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRecord struct {
	Role *enum_role.Value `bson:"role,omitempty"`
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Role != nil {
		elemrole0 := r.Role
		m.Role = *elemrole0
	}
	return m, nil
}

type MongoWhereClause struct {
	// role (Role) search options
	RoleEq     *enum_role.Value
	RoleNe     *enum_role.Value
	RoleGt     *enum_role.Value
	RoleGte    *enum_role.Value
	RoleLt     *enum_role.Value
	RoleLte    *enum_role.Value
	RoleIn     *[]enum_role.Value
	RoleNin    *[]enum_role.Value
	RoleExists *bool
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
	if o.RoleEq != nil {
		query := bson.M{}
		query["role"] = o.RoleEq
		and = append(and, query)
	}
	if o.RoleNe != nil {
		query := bson.M{}
		query["role"] = bson.M{"$ne": o.RoleNe}
		and = append(and, query)
	}
	if o.RoleGt != nil {
		query := bson.M{}
		query["role"] = bson.M{"$gt": o.RoleGt}
		and = append(and, query)
	}
	if o.RoleGte != nil {
		query := bson.M{}
		query["role"] = bson.M{"$gte": o.RoleGte}
		and = append(and, query)
	}
	if o.RoleLt != nil {
		query := bson.M{}
		query["role"] = bson.M{"$lt": o.RoleLt}
		and = append(and, query)
	}
	if o.RoleLte != nil {
		query := bson.M{}
		query["role"] = bson.M{"$lte": o.RoleLte}
		and = append(and, query)
	}
	if o.RoleIn != nil {
		query := bson.M{}
		query["role"] = bson.M{"$in": o.RoleIn}
		and = append(and, query)
	}
	if o.RoleNin != nil {
		query := bson.M{}
		query["role"] = bson.M{"$nin": o.RoleNin}
		and = append(and, query)
	}
	if o.RoleExists != nil {
		query := bson.M{}
		query["role"] = bson.M{"$exists": *o.RoleExists}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
}
