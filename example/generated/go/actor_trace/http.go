package actor_trace

import (
	"time"
)

type HTTPRecord struct {
	ActorId   *string    `json:"actorId,omitempty"`
	ActorName *string    `json:"actorName,omitempty"`
	ActorType *string    `json:"actorType,omitempty"`
	At        *time.Time `json:"at,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.ActorId != nil {
		p.ActorId = true
	}
	if r.ActorName != nil {
		p.ActorName = true
	}
	if r.ActorType != nil {
		p.ActorType = true
	}
	if r.At != nil {
		p.At = true
	}
	return p, nil
}

type HTTPWhereClause struct {
	// actorId (string) search options
	ActorIdEq     *string   `json:"actorIdEq,omitempty"`
	ActorIdNe     *string   `json:"actorIdNe,omitempty"`
	ActorIdGt     *string   `json:"actorIdGt,omitempty"`
	ActorIdGte    *string   `json:"actorIdGte,omitempty"`
	ActorIdLt     *string   `json:"actorIdLt,omitempty"`
	ActorIdLte    *string   `json:"actorIdLte,omitempty"`
	ActorIdIn     *[]string `json:"actorIdIn,omitempty"`
	ActorIdNin    *[]string `json:"actorIdNin,omitempty"`
	ActorIdExists *bool     `json:"actorIdExists,omitempty"`
	ActorIdLike   *string   `json:"actorIdLike,omitempty"`
	ActorIdNlike  *string   `json:"actorIdNlike,omitempty"`
	// actorName (string) search options
	ActorNameEq     *string   `json:"actorNameEq,omitempty"`
	ActorNameNe     *string   `json:"actorNameNe,omitempty"`
	ActorNameGt     *string   `json:"actorNameGt,omitempty"`
	ActorNameGte    *string   `json:"actorNameGte,omitempty"`
	ActorNameLt     *string   `json:"actorNameLt,omitempty"`
	ActorNameLte    *string   `json:"actorNameLte,omitempty"`
	ActorNameIn     *[]string `json:"actorNameIn,omitempty"`
	ActorNameNin    *[]string `json:"actorNameNin,omitempty"`
	ActorNameExists *bool     `json:"actorNameExists,omitempty"`
	ActorNameLike   *string   `json:"actorNameLike,omitempty"`
	ActorNameNlike  *string   `json:"actorNameNlike,omitempty"`
	// actorType (string) search options
	ActorTypeEq     *string   `json:"actorTypeEq,omitempty"`
	ActorTypeNe     *string   `json:"actorTypeNe,omitempty"`
	ActorTypeGt     *string   `json:"actorTypeGt,omitempty"`
	ActorTypeGte    *string   `json:"actorTypeGte,omitempty"`
	ActorTypeLt     *string   `json:"actorTypeLt,omitempty"`
	ActorTypeLte    *string   `json:"actorTypeLte,omitempty"`
	ActorTypeIn     *[]string `json:"actorTypeIn,omitempty"`
	ActorTypeNin    *[]string `json:"actorTypeNin,omitempty"`
	ActorTypeExists *bool     `json:"actorTypeExists,omitempty"`
	ActorTypeLike   *string   `json:"actorTypeLike,omitempty"`
	ActorTypeNlike  *string   `json:"actorTypeNlike,omitempty"`
	// at (timestamp) search options
	AtEq     *time.Time   `json:"atEq,omitempty"`
	AtNe     *time.Time   `json:"atNe,omitempty"`
	AtGt     *time.Time   `json:"atGt,omitempty"`
	AtGte    *time.Time   `json:"atGte,omitempty"`
	AtLt     *time.Time   `json:"atLt,omitempty"`
	AtLte    *time.Time   `json:"atLte,omitempty"`
	AtIn     *[]time.Time `json:"atIn,omitempty"`
	AtNin    *[]time.Time `json:"atNin,omitempty"`
	AtExists *bool        `json:"atExists,omitempty"`
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
	if o.ActorIdEq != nil {
		elemactorIdEq0 := o.ActorIdEq
		to.ActorIdEq = elemactorIdEq0
	}
	if o.ActorIdNe != nil {
		elemactorIdNe0 := o.ActorIdNe
		to.ActorIdNe = elemactorIdNe0
	}
	if o.ActorIdGt != nil {
		elemactorIdGt0 := o.ActorIdGt
		to.ActorIdGt = elemactorIdGt0
	}
	if o.ActorIdGte != nil {
		elemactorIdGte0 := o.ActorIdGte
		to.ActorIdGte = elemactorIdGte0
	}
	if o.ActorIdLt != nil {
		elemactorIdLt0 := o.ActorIdLt
		to.ActorIdLt = elemactorIdLt0
	}
	if o.ActorIdLte != nil {
		elemactorIdLte0 := o.ActorIdLte
		to.ActorIdLte = elemactorIdLte0
	}
	if o.ActorIdIn != nil {
		elemactorIdIn0 := make([]string, 0)
		for _, oactorIdIn0 := range *o.ActorIdIn {
			elemactorIdIn1 := oactorIdIn0
			elemactorIdIn0 = append(elemactorIdIn0, elemactorIdIn1)
		}
		to.ActorIdIn = &elemactorIdIn0
	}
	if o.ActorIdNin != nil {
		elemactorIdNin0 := make([]string, 0)
		for _, oactorIdNin0 := range *o.ActorIdNin {
			elemactorIdNin1 := oactorIdNin0
			elemactorIdNin0 = append(elemactorIdNin0, elemactorIdNin1)
		}
		to.ActorIdNin = &elemactorIdNin0
	}
	if o.ActorIdExists != nil {
		elemactorIdExists0 := o.ActorIdExists
		to.ActorIdExists = elemactorIdExists0
	}
	if o.ActorIdLike != nil {
		elemactorIdLike0 := o.ActorIdLike
		to.ActorIdLike = elemactorIdLike0
	}
	if o.ActorIdNlike != nil {
		elemactorIdNlike0 := o.ActorIdNlike
		to.ActorIdNlike = elemactorIdNlike0
	}
	if o.ActorNameEq != nil {
		elemactorNameEq0 := o.ActorNameEq
		to.ActorNameEq = elemactorNameEq0
	}
	if o.ActorNameNe != nil {
		elemactorNameNe0 := o.ActorNameNe
		to.ActorNameNe = elemactorNameNe0
	}
	if o.ActorNameGt != nil {
		elemactorNameGt0 := o.ActorNameGt
		to.ActorNameGt = elemactorNameGt0
	}
	if o.ActorNameGte != nil {
		elemactorNameGte0 := o.ActorNameGte
		to.ActorNameGte = elemactorNameGte0
	}
	if o.ActorNameLt != nil {
		elemactorNameLt0 := o.ActorNameLt
		to.ActorNameLt = elemactorNameLt0
	}
	if o.ActorNameLte != nil {
		elemactorNameLte0 := o.ActorNameLte
		to.ActorNameLte = elemactorNameLte0
	}
	if o.ActorNameIn != nil {
		elemactorNameIn0 := make([]string, 0)
		for _, oactorNameIn0 := range *o.ActorNameIn {
			elemactorNameIn1 := oactorNameIn0
			elemactorNameIn0 = append(elemactorNameIn0, elemactorNameIn1)
		}
		to.ActorNameIn = &elemactorNameIn0
	}
	if o.ActorNameNin != nil {
		elemactorNameNin0 := make([]string, 0)
		for _, oactorNameNin0 := range *o.ActorNameNin {
			elemactorNameNin1 := oactorNameNin0
			elemactorNameNin0 = append(elemactorNameNin0, elemactorNameNin1)
		}
		to.ActorNameNin = &elemactorNameNin0
	}
	if o.ActorNameExists != nil {
		elemactorNameExists0 := o.ActorNameExists
		to.ActorNameExists = elemactorNameExists0
	}
	if o.ActorNameLike != nil {
		elemactorNameLike0 := o.ActorNameLike
		to.ActorNameLike = elemactorNameLike0
	}
	if o.ActorNameNlike != nil {
		elemactorNameNlike0 := o.ActorNameNlike
		to.ActorNameNlike = elemactorNameNlike0
	}
	if o.ActorTypeEq != nil {
		elemactorTypeEq0 := o.ActorTypeEq
		to.ActorTypeEq = elemactorTypeEq0
	}
	if o.ActorTypeNe != nil {
		elemactorTypeNe0 := o.ActorTypeNe
		to.ActorTypeNe = elemactorTypeNe0
	}
	if o.ActorTypeGt != nil {
		elemactorTypeGt0 := o.ActorTypeGt
		to.ActorTypeGt = elemactorTypeGt0
	}
	if o.ActorTypeGte != nil {
		elemactorTypeGte0 := o.ActorTypeGte
		to.ActorTypeGte = elemactorTypeGte0
	}
	if o.ActorTypeLt != nil {
		elemactorTypeLt0 := o.ActorTypeLt
		to.ActorTypeLt = elemactorTypeLt0
	}
	if o.ActorTypeLte != nil {
		elemactorTypeLte0 := o.ActorTypeLte
		to.ActorTypeLte = elemactorTypeLte0
	}
	if o.ActorTypeIn != nil {
		elemactorTypeIn0 := make([]string, 0)
		for _, oactorTypeIn0 := range *o.ActorTypeIn {
			elemactorTypeIn1 := oactorTypeIn0
			elemactorTypeIn0 = append(elemactorTypeIn0, elemactorTypeIn1)
		}
		to.ActorTypeIn = &elemactorTypeIn0
	}
	if o.ActorTypeNin != nil {
		elemactorTypeNin0 := make([]string, 0)
		for _, oactorTypeNin0 := range *o.ActorTypeNin {
			elemactorTypeNin1 := oactorTypeNin0
			elemactorTypeNin0 = append(elemactorTypeNin0, elemactorTypeNin1)
		}
		to.ActorTypeNin = &elemactorTypeNin0
	}
	if o.ActorTypeExists != nil {
		elemactorTypeExists0 := o.ActorTypeExists
		to.ActorTypeExists = elemactorTypeExists0
	}
	if o.ActorTypeLike != nil {
		elemactorTypeLike0 := o.ActorTypeLike
		to.ActorTypeLike = elemactorTypeLike0
	}
	if o.ActorTypeNlike != nil {
		elemactorTypeNlike0 := o.ActorTypeNlike
		to.ActorTypeNlike = elemactorTypeNlike0
	}
	if o.AtEq != nil {
		elematEq0 := o.AtEq
		to.AtEq = elematEq0
	}
	if o.AtNe != nil {
		elematNe0 := o.AtNe
		to.AtNe = elematNe0
	}
	if o.AtGt != nil {
		elematGt0 := o.AtGt
		to.AtGt = elematGt0
	}
	if o.AtGte != nil {
		elematGte0 := o.AtGte
		to.AtGte = elematGte0
	}
	if o.AtLt != nil {
		elematLt0 := o.AtLt
		to.AtLt = elematLt0
	}
	if o.AtLte != nil {
		elematLte0 := o.AtLte
		to.AtLte = elematLte0
	}
	if o.AtIn != nil {
		elematIn0 := make([]time.Time, 0)
		for _, oatIn0 := range *o.AtIn {
			elematIn1 := oatIn0
			elematIn0 = append(elematIn0, elematIn1)
		}
		to.AtIn = &elematIn0
	}
	if o.AtNin != nil {
		elematNin0 := make([]time.Time, 0)
		for _, oatNin0 := range *o.AtNin {
			elematNin1 := oatNin0
			elematNin0 = append(elematNin0, elematNin1)
		}
		to.AtNin = &elematNin0
	}
	if o.AtExists != nil {
		elematExists0 := o.AtExists
		to.AtExists = elematExists0
	}
	return to, nil
}

type HTTPSortParams struct {
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	return to
}
