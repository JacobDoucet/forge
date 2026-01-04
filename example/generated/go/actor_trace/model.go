package actor_trace

import (
	"time"
)

type Model struct {
	ActorId   string
	ActorName string
	ActorType string
	At        time.Time
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if projection.ActorId {
		elemactorId0 := m.ActorId
		r.ActorId = &elemactorId0
	}
	if projection.ActorName {
		elemactorName0 := m.ActorName
		r.ActorName = &elemactorName0
	}
	if projection.ActorType {
		elemactorType0 := m.ActorType
		r.ActorType = &elemactorType0
	}
	if projection.At {
		elemat0 := m.At
		r.At = &elemat0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if projection.ActorId {
		elemactorId0 := m.ActorId
		r.ActorId = &elemactorId0
	}
	if projection.ActorName {
		elemactorName0 := m.ActorName
		r.ActorName = &elemactorName0
	}
	if projection.ActorType {
		elemactorType0 := m.ActorType
		r.ActorType = &elemactorType0
	}
	if projection.At {
		elemat0 := m.At
		r.At = &elemat0
	}
	return r, nil
}

type WhereClause struct {
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

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
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

type SortParams struct {
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	return to
}
