package event_subject

import (
	"github.com/JacobDoucet/forge/example/generated/go/enum_model"
)

type Model struct {
	SubjectId   string
	SubjectType enum_model.Value
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if projection.SubjectId {
		elemsubjectId0 := m.SubjectId
		r.SubjectId = &elemsubjectId0
	}
	if projection.SubjectType {
		elemsubjectType0 := m.SubjectType
		r.SubjectType = &elemsubjectType0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if projection.SubjectId {
		elemsubjectId0 := m.SubjectId
		r.SubjectId = &elemsubjectId0
	}
	if projection.SubjectType {
		elemsubjectType0 := m.SubjectType
		r.SubjectType = &elemsubjectType0
	}
	return r, nil
}

type WhereClause struct {
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

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.SubjectIdEq != nil {
		elemsubjectIdEq0 := o.SubjectIdEq
		to.SubjectIdEq = elemsubjectIdEq0
	}
	if o.SubjectIdNe != nil {
		elemsubjectIdNe0 := o.SubjectIdNe
		to.SubjectIdNe = elemsubjectIdNe0
	}
	if o.SubjectIdGt != nil {
		elemsubjectIdGt0 := o.SubjectIdGt
		to.SubjectIdGt = elemsubjectIdGt0
	}
	if o.SubjectIdGte != nil {
		elemsubjectIdGte0 := o.SubjectIdGte
		to.SubjectIdGte = elemsubjectIdGte0
	}
	if o.SubjectIdLt != nil {
		elemsubjectIdLt0 := o.SubjectIdLt
		to.SubjectIdLt = elemsubjectIdLt0
	}
	if o.SubjectIdLte != nil {
		elemsubjectIdLte0 := o.SubjectIdLte
		to.SubjectIdLte = elemsubjectIdLte0
	}
	if o.SubjectIdIn != nil {
		elemsubjectIdIn0 := make([]string, 0)
		for _, osubjectIdIn0 := range *o.SubjectIdIn {
			elemsubjectIdIn1 := osubjectIdIn0
			elemsubjectIdIn0 = append(elemsubjectIdIn0, elemsubjectIdIn1)
		}
		to.SubjectIdIn = &elemsubjectIdIn0
	}
	if o.SubjectIdNin != nil {
		elemsubjectIdNin0 := make([]string, 0)
		for _, osubjectIdNin0 := range *o.SubjectIdNin {
			elemsubjectIdNin1 := osubjectIdNin0
			elemsubjectIdNin0 = append(elemsubjectIdNin0, elemsubjectIdNin1)
		}
		to.SubjectIdNin = &elemsubjectIdNin0
	}
	if o.SubjectIdExists != nil {
		elemsubjectIdExists0 := o.SubjectIdExists
		to.SubjectIdExists = elemsubjectIdExists0
	}
	if o.SubjectIdLike != nil {
		elemsubjectIdLike0 := o.SubjectIdLike
		to.SubjectIdLike = elemsubjectIdLike0
	}
	if o.SubjectIdNlike != nil {
		elemsubjectIdNlike0 := o.SubjectIdNlike
		to.SubjectIdNlike = elemsubjectIdNlike0
	}
	if o.SubjectTypeEq != nil {
		elemsubjectTypeEq0 := o.SubjectTypeEq
		to.SubjectTypeEq = elemsubjectTypeEq0
	}
	if o.SubjectTypeNe != nil {
		elemsubjectTypeNe0 := o.SubjectTypeNe
		to.SubjectTypeNe = elemsubjectTypeNe0
	}
	if o.SubjectTypeGt != nil {
		elemsubjectTypeGt0 := o.SubjectTypeGt
		to.SubjectTypeGt = elemsubjectTypeGt0
	}
	if o.SubjectTypeGte != nil {
		elemsubjectTypeGte0 := o.SubjectTypeGte
		to.SubjectTypeGte = elemsubjectTypeGte0
	}
	if o.SubjectTypeLt != nil {
		elemsubjectTypeLt0 := o.SubjectTypeLt
		to.SubjectTypeLt = elemsubjectTypeLt0
	}
	if o.SubjectTypeLte != nil {
		elemsubjectTypeLte0 := o.SubjectTypeLte
		to.SubjectTypeLte = elemsubjectTypeLte0
	}
	if o.SubjectTypeIn != nil {
		elemsubjectTypeIn0 := make([]enum_model.Value, 0)
		for _, osubjectTypeIn0 := range *o.SubjectTypeIn {
			elemsubjectTypeIn1 := osubjectTypeIn0
			elemsubjectTypeIn0 = append(elemsubjectTypeIn0, elemsubjectTypeIn1)
		}
		to.SubjectTypeIn = &elemsubjectTypeIn0
	}
	if o.SubjectTypeNin != nil {
		elemsubjectTypeNin0 := make([]enum_model.Value, 0)
		for _, osubjectTypeNin0 := range *o.SubjectTypeNin {
			elemsubjectTypeNin1 := osubjectTypeNin0
			elemsubjectTypeNin0 = append(elemsubjectTypeNin0, elemsubjectTypeNin1)
		}
		to.SubjectTypeNin = &elemsubjectTypeNin0
	}
	if o.SubjectTypeExists != nil {
		elemsubjectTypeExists0 := o.SubjectTypeExists
		to.SubjectTypeExists = elemsubjectTypeExists0
	}
	return to, nil
}

type SortParams struct {
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	return to
}
