package user

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id            *primitive.ObjectID       `bson:"_id,omitempty"`
	ActorRoles    *[]actor_role.MongoRecord `bson:"actorRoles,omitempty"`
	Created       *actor_trace.MongoRecord  `bson:"created,omitempty"`
	Email         *string                   `bson:"email,omitempty"`
	FirstName     *string                   `bson:"firstName,omitempty"`
	Language      *string                   `bson:"language,omitempty"`
	LastName      *string                   `bson:"lastName,omitempty"`
	Role          *enum_role.Value          `bson:"role,omitempty"`
	Updated       *actor_trace.MongoRecord  `bson:"updated,omitempty"`
	UpdatedByUser *actor_trace.MongoRecord  `bson:"updatedByUser,omitempty"`
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
	if r.ActorRoles != nil {
		elemactorRoles0 := make([]actor_role.Model, 0)
		for _, ractorRoles0 := range *r.ActorRoles {
			elemactorRoles1, err := ractorRoles0.ToModel()
			if err != nil {
				return m, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		m.ActorRoles = elemactorRoles0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Email != nil {
		elememail0 := r.Email
		m.Email = *elememail0
	}
	if r.FirstName != nil {
		elemfirstName0 := r.FirstName
		m.FirstName = *elemfirstName0
	}
	if r.Language != nil {
		elemlanguage0 := r.Language
		m.Language = *elemlanguage0
	}
	if r.LastName != nil {
		elemlastName0 := r.LastName
		m.LastName = *elemlastName0
	}
	if r.Role != nil {
		elemrole0 := r.Role
		m.Role = *elemrole0
	}
	if r.Updated != nil {
		elemupdated0, err := r.Updated.ToModel()
		if err != nil {
			return m, err
		}
		m.Updated = elemupdated0
	}
	if r.UpdatedByUser != nil {
		elemupdatedByUser0, err := r.UpdatedByUser.ToModel()
		if err != nil {
			return m, err
		}
		m.UpdatedByUser = elemupdatedByUser0
	}
	return m, nil
}

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}
type MongoSelectByEmailIdxQuery struct {
	Email string
}

type MongoWhereClause struct {
	// id (Ref<User>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.MongoWhereClause
	ActorRolesEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// email (string) search options
	EmailEq     *string
	EmailNe     *string
	EmailGt     *string
	EmailGte    *string
	EmailLt     *string
	EmailLte    *string
	EmailIn     *[]string
	EmailNin    *[]string
	EmailExists *bool
	EmailLike   *string
	EmailNlike  *string
	// firstName (string) search options
	FirstNameEq     *string
	FirstNameNe     *string
	FirstNameGt     *string
	FirstNameGte    *string
	FirstNameLt     *string
	FirstNameLte    *string
	FirstNameIn     *[]string
	FirstNameNin    *[]string
	FirstNameExists *bool
	FirstNameLike   *string
	FirstNameNlike  *string
	// language (string) search options
	LanguageEq     *string
	LanguageNe     *string
	LanguageGt     *string
	LanguageGte    *string
	LanguageLt     *string
	LanguageLte    *string
	LanguageIn     *[]string
	LanguageNin    *[]string
	LanguageExists *bool
	LanguageLike   *string
	LanguageNlike  *string
	// lastName (string) search options
	LastNameEq     *string
	LastNameNe     *string
	LastNameGt     *string
	LastNameGte    *string
	LastNameLt     *string
	LastNameLte    *string
	LastNameIn     *[]string
	LastNameNin    *[]string
	LastNameExists *bool
	LastNameLike   *string
	LastNameNlike  *string
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
	// updated (ActorTrace) search options
	Updated *actor_trace.MongoWhereClause
	// updatedByUser (ActorTrace) search options
	UpdatedByUser *actor_trace.MongoWhereClause
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
	if o.ActorRoles != nil {
		query := bson.M{}
		actorRolesQuery, err := o.ActorRoles.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range actorRolesQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["actorRoles."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.ActorRolesEmpty != nil {
		query := bson.M{}
		if *o.ActorRolesEmpty {
			query["$or"] = bson.A{
				bson.M{"actorRoles": nil},
				bson.M{"actorRoles": bson.A{}},
				bson.M{"actorRoles": bson.M{"$exists": false}},
			}
		} else {
			query["actorRoles"] = bson.M{
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
	if o.EmailEq != nil {
		query := bson.M{}
		query["email"] = o.EmailEq
		and = append(and, query)
	}
	if o.EmailNe != nil {
		query := bson.M{}
		query["email"] = bson.M{"$ne": o.EmailNe}
		and = append(and, query)
	}
	if o.EmailGt != nil {
		query := bson.M{}
		query["email"] = bson.M{"$gt": o.EmailGt}
		and = append(and, query)
	}
	if o.EmailGte != nil {
		query := bson.M{}
		query["email"] = bson.M{"$gte": o.EmailGte}
		and = append(and, query)
	}
	if o.EmailLt != nil {
		query := bson.M{}
		query["email"] = bson.M{"$lt": o.EmailLt}
		and = append(and, query)
	}
	if o.EmailLte != nil {
		query := bson.M{}
		query["email"] = bson.M{"$lte": o.EmailLte}
		and = append(and, query)
	}
	if o.EmailIn != nil {
		query := bson.M{}
		query["email"] = bson.M{"$in": o.EmailIn}
		and = append(and, query)
	}
	if o.EmailNin != nil {
		query := bson.M{}
		query["email"] = bson.M{"$nin": o.EmailNin}
		and = append(and, query)
	}
	if o.EmailExists != nil {
		query := bson.M{}
		query["email"] = bson.M{"$exists": *o.EmailExists}
		and = append(and, query)
	}
	if o.EmailLike != nil {
		query := bson.M{}
		query["email"] = bson.M{"$regex": o.EmailLike, "$options": "i"}
		and = append(and, query)
	}
	if o.EmailNlike != nil {
		query := bson.M{}
		query["email"] = bson.M{"$not": bson.M{"$regex": o.EmailNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.FirstNameEq != nil {
		query := bson.M{}
		query["firstName"] = o.FirstNameEq
		and = append(and, query)
	}
	if o.FirstNameNe != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$ne": o.FirstNameNe}
		and = append(and, query)
	}
	if o.FirstNameGt != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$gt": o.FirstNameGt}
		and = append(and, query)
	}
	if o.FirstNameGte != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$gte": o.FirstNameGte}
		and = append(and, query)
	}
	if o.FirstNameLt != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$lt": o.FirstNameLt}
		and = append(and, query)
	}
	if o.FirstNameLte != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$lte": o.FirstNameLte}
		and = append(and, query)
	}
	if o.FirstNameIn != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$in": o.FirstNameIn}
		and = append(and, query)
	}
	if o.FirstNameNin != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$nin": o.FirstNameNin}
		and = append(and, query)
	}
	if o.FirstNameExists != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$exists": *o.FirstNameExists}
		and = append(and, query)
	}
	if o.FirstNameLike != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$regex": o.FirstNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.FirstNameNlike != nil {
		query := bson.M{}
		query["firstName"] = bson.M{"$not": bson.M{"$regex": o.FirstNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.LanguageEq != nil {
		query := bson.M{}
		query["language"] = o.LanguageEq
		and = append(and, query)
	}
	if o.LanguageNe != nil {
		query := bson.M{}
		query["language"] = bson.M{"$ne": o.LanguageNe}
		and = append(and, query)
	}
	if o.LanguageGt != nil {
		query := bson.M{}
		query["language"] = bson.M{"$gt": o.LanguageGt}
		and = append(and, query)
	}
	if o.LanguageGte != nil {
		query := bson.M{}
		query["language"] = bson.M{"$gte": o.LanguageGte}
		and = append(and, query)
	}
	if o.LanguageLt != nil {
		query := bson.M{}
		query["language"] = bson.M{"$lt": o.LanguageLt}
		and = append(and, query)
	}
	if o.LanguageLte != nil {
		query := bson.M{}
		query["language"] = bson.M{"$lte": o.LanguageLte}
		and = append(and, query)
	}
	if o.LanguageIn != nil {
		query := bson.M{}
		query["language"] = bson.M{"$in": o.LanguageIn}
		and = append(and, query)
	}
	if o.LanguageNin != nil {
		query := bson.M{}
		query["language"] = bson.M{"$nin": o.LanguageNin}
		and = append(and, query)
	}
	if o.LanguageExists != nil {
		query := bson.M{}
		query["language"] = bson.M{"$exists": *o.LanguageExists}
		and = append(and, query)
	}
	if o.LanguageLike != nil {
		query := bson.M{}
		query["language"] = bson.M{"$regex": o.LanguageLike, "$options": "i"}
		and = append(and, query)
	}
	if o.LanguageNlike != nil {
		query := bson.M{}
		query["language"] = bson.M{"$not": bson.M{"$regex": o.LanguageNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.LastNameEq != nil {
		query := bson.M{}
		query["lastName"] = o.LastNameEq
		and = append(and, query)
	}
	if o.LastNameNe != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$ne": o.LastNameNe}
		and = append(and, query)
	}
	if o.LastNameGt != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$gt": o.LastNameGt}
		and = append(and, query)
	}
	if o.LastNameGte != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$gte": o.LastNameGte}
		and = append(and, query)
	}
	if o.LastNameLt != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$lt": o.LastNameLt}
		and = append(and, query)
	}
	if o.LastNameLte != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$lte": o.LastNameLte}
		and = append(and, query)
	}
	if o.LastNameIn != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$in": o.LastNameIn}
		and = append(and, query)
	}
	if o.LastNameNin != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$nin": o.LastNameNin}
		and = append(and, query)
	}
	if o.LastNameExists != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$exists": *o.LastNameExists}
		and = append(and, query)
	}
	if o.LastNameLike != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$regex": o.LastNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.LastNameNlike != nil {
		query := bson.M{}
		query["lastName"] = bson.M{"$not": bson.M{"$regex": o.LastNameNlike, "$options": "i"}}
		and = append(and, query)
	}
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
	if o.UpdatedByUser != nil {
		query := bson.M{}
		updatedByUserQuery, err := o.UpdatedByUser.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range updatedByUserQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["updatedByUser."+k] = v
			}
		}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
	CreatedAt int8
	Email     int8
	UpdatedAt int8
}
