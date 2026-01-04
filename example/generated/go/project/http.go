package project

import (
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
)

type HTTPRecord struct {
	Id          *string                 `json:"id,omitempty"`
	Created     *actor_trace.HTTPRecord `json:"created,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	OwnerId     *string                 `json:"ownerId,omitempty"`
	Updated     *actor_trace.HTTPRecord `json:"updated,omitempty"`
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Description != nil {
		p.Description = true
	}
	if r.Name != nil {
		p.Name = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
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
	// id (Ref<Project>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
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
	// name (string) search options
	NameEq     *string   `json:"nameEq,omitempty"`
	NameNe     *string   `json:"nameNe,omitempty"`
	NameGt     *string   `json:"nameGt,omitempty"`
	NameGte    *string   `json:"nameGte,omitempty"`
	NameLt     *string   `json:"nameLt,omitempty"`
	NameLte    *string   `json:"nameLte,omitempty"`
	NameIn     *[]string `json:"nameIn,omitempty"`
	NameNin    *[]string `json:"nameNin,omitempty"`
	NameExists *bool     `json:"nameExists,omitempty"`
	NameLike   *string   `json:"nameLike,omitempty"`
	NameNlike  *string   `json:"nameNlike,omitempty"`
	// ownerId (string) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdNe     *string   `json:"ownerIdNe,omitempty"`
	OwnerIdGt     *string   `json:"ownerIdGt,omitempty"`
	OwnerIdGte    *string   `json:"ownerIdGte,omitempty"`
	OwnerIdLt     *string   `json:"ownerIdLt,omitempty"`
	OwnerIdLte    *string   `json:"ownerIdLte,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	OwnerIdLike   *string   `json:"ownerIdLike,omitempty"`
	OwnerIdNlike  *string   `json:"ownerIdNlike,omitempty"`
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
		elemupdated0, err := o.Updated.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	return to, nil
}

type HTTPSortParams struct {
	CreatedAt *int8 `json:"createdAt,omitempty"`
	UpdatedAt *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
