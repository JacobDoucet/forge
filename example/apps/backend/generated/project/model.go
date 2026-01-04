package project

import (
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id          string
	Created     actor_trace.Model
	Description string
	Name        string
	OwnerId     string
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
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
	}
	if projection.OwnerId {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
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
	if projection.Description {
		elemdescription0 := m.Description
		r.Description = &elemdescription0
	}
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
	}
	if projection.OwnerId {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
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
	// id (Ref<Project>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
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
	if o.NameEq != nil {
		elemnameEq0 := o.NameEq
		to.NameEq = elemnameEq0
	}
	if o.NameNe != nil {
		elemnameNe0 := o.NameNe
		to.NameNe = elemnameNe0
	}
	if o.NameGt != nil {
		elemnameGt0 := o.NameGt
		to.NameGt = elemnameGt0
	}
	if o.NameGte != nil {
		elemnameGte0 := o.NameGte
		to.NameGte = elemnameGte0
	}
	if o.NameLt != nil {
		elemnameLt0 := o.NameLt
		to.NameLt = elemnameLt0
	}
	if o.NameLte != nil {
		elemnameLte0 := o.NameLte
		to.NameLte = elemnameLte0
	}
	if o.NameIn != nil {
		elemnameIn0 := make([]string, 0)
		for _, onameIn0 := range *o.NameIn {
			elemnameIn1 := onameIn0
			elemnameIn0 = append(elemnameIn0, elemnameIn1)
		}
		to.NameIn = &elemnameIn0
	}
	if o.NameNin != nil {
		elemnameNin0 := make([]string, 0)
		for _, onameNin0 := range *o.NameNin {
			elemnameNin1 := onameNin0
			elemnameNin0 = append(elemnameNin0, elemnameNin1)
		}
		to.NameNin = &elemnameNin0
	}
	if o.NameExists != nil {
		elemnameExists0 := o.NameExists
		to.NameExists = elemnameExists0
	}
	if o.NameLike != nil {
		elemnameLike0 := o.NameLike
		to.NameLike = elemnameLike0
	}
	if o.NameNlike != nil {
		elemnameNlike0 := o.NameNlike
		to.NameNlike = elemnameNlike0
	}
	if o.OwnerIdEq != nil {
		elemownerIdEq0 := o.OwnerIdEq
		to.OwnerIdEq = elemownerIdEq0
	}
	if o.OwnerIdNe != nil {
		elemownerIdNe0 := o.OwnerIdNe
		to.OwnerIdNe = elemownerIdNe0
	}
	if o.OwnerIdGt != nil {
		elemownerIdGt0 := o.OwnerIdGt
		to.OwnerIdGt = elemownerIdGt0
	}
	if o.OwnerIdGte != nil {
		elemownerIdGte0 := o.OwnerIdGte
		to.OwnerIdGte = elemownerIdGte0
	}
	if o.OwnerIdLt != nil {
		elemownerIdLt0 := o.OwnerIdLt
		to.OwnerIdLt = elemownerIdLt0
	}
	if o.OwnerIdLte != nil {
		elemownerIdLte0 := o.OwnerIdLte
		to.OwnerIdLte = elemownerIdLte0
	}
	if o.OwnerIdIn != nil {
		elemownerIdIn0 := make([]string, 0)
		for _, oownerIdIn0 := range *o.OwnerIdIn {
			elemownerIdIn1 := oownerIdIn0
			elemownerIdIn0 = append(elemownerIdIn0, elemownerIdIn1)
		}
		to.OwnerIdIn = &elemownerIdIn0
	}
	if o.OwnerIdNin != nil {
		elemownerIdNin0 := make([]string, 0)
		for _, oownerIdNin0 := range *o.OwnerIdNin {
			elemownerIdNin1 := oownerIdNin0
			elemownerIdNin0 = append(elemownerIdNin0, elemownerIdNin1)
		}
		to.OwnerIdNin = &elemownerIdNin0
	}
	if o.OwnerIdExists != nil {
		elemownerIdExists0 := o.OwnerIdExists
		to.OwnerIdExists = elemownerIdExists0
	}
	if o.OwnerIdLike != nil {
		elemownerIdLike0 := o.OwnerIdLike
		to.OwnerIdLike = elemownerIdLike0
	}
	if o.OwnerIdNlike != nil {
		elemownerIdNlike0 := o.OwnerIdNlike
		to.OwnerIdNlike = elemownerIdNlike0
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
	CreatedAt int8
	UpdatedAt int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.UpdatedAt = s.UpdatedAt
	return to
}
