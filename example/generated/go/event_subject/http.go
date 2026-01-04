package event_subject

import (
	"github.com/JacobDoucet/forge/example/generated/go/enum_model"
)

type HTTPRecord struct {
	SubjectId   *string           `json:"subjectId,omitempty"`
	SubjectType *enum_model.Value `json:"subjectType,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.SubjectId != nil {
		p.SubjectId = true
	}
	if r.SubjectType != nil {
		p.SubjectType = true
	}
	return p, nil
}

type HTTPWhereClause struct {
	// subjectId (string) search options
	SubjectIdEq     *string   `json:"subjectIdEq,omitempty"`
	SubjectIdNe     *string   `json:"subjectIdNe,omitempty"`
	SubjectIdGt     *string   `json:"subjectIdGt,omitempty"`
	SubjectIdGte    *string   `json:"subjectIdGte,omitempty"`
	SubjectIdLt     *string   `json:"subjectIdLt,omitempty"`
	SubjectIdLte    *string   `json:"subjectIdLte,omitempty"`
	SubjectIdIn     *[]string `json:"subjectIdIn,omitempty"`
	SubjectIdNin    *[]string `json:"subjectIdNin,omitempty"`
	SubjectIdExists *bool     `json:"subjectIdExists,omitempty"`
	SubjectIdLike   *string   `json:"subjectIdLike,omitempty"`
	SubjectIdNlike  *string   `json:"subjectIdNlike,omitempty"`
	// subjectType (Model) search options
	SubjectTypeEq     *enum_model.Value   `json:"subjectTypeEq,omitempty"`
	SubjectTypeNe     *enum_model.Value   `json:"subjectTypeNe,omitempty"`
	SubjectTypeGt     *enum_model.Value   `json:"subjectTypeGt,omitempty"`
	SubjectTypeGte    *enum_model.Value   `json:"subjectTypeGte,omitempty"`
	SubjectTypeLt     *enum_model.Value   `json:"subjectTypeLt,omitempty"`
	SubjectTypeLte    *enum_model.Value   `json:"subjectTypeLte,omitempty"`
	SubjectTypeIn     *[]enum_model.Value `json:"subjectTypeIn,omitempty"`
	SubjectTypeNin    *[]enum_model.Value `json:"subjectTypeNin,omitempty"`
	SubjectTypeExists *bool               `json:"subjectTypeExists,omitempty"`
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
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

type HTTPSortParams struct {
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	return to
}
