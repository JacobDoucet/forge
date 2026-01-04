package task_comment

import (
	"time"
)

type Model struct {
	AuthorId  string
	CreatedAt time.Time
	Text      string
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if projection.AuthorId {
		elemauthorId0 := m.AuthorId
		r.AuthorId = &elemauthorId0
	}
	if projection.CreatedAt {
		elemcreatedAt0 := m.CreatedAt
		r.CreatedAt = &elemcreatedAt0
	}
	if projection.Text {
		elemtext0 := m.Text
		r.Text = &elemtext0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if projection.AuthorId {
		elemauthorId0 := m.AuthorId
		r.AuthorId = &elemauthorId0
	}
	if projection.CreatedAt {
		elemcreatedAt0 := m.CreatedAt
		r.CreatedAt = &elemcreatedAt0
	}
	if projection.Text {
		elemtext0 := m.Text
		r.Text = &elemtext0
	}
	return r, nil
}

type WhereClause struct {
	// authorId (string) search options
	AuthorIdEq     *string
	AuthorIdNe     *string
	AuthorIdGt     *string
	AuthorIdGte    *string
	AuthorIdLt     *string
	AuthorIdLte    *string
	AuthorIdIn     *[]string
	AuthorIdNin    *[]string
	AuthorIdExists *bool
	AuthorIdLike   *string
	AuthorIdNlike  *string
	// createdAt (timestamp) search options
	CreatedAtEq     *time.Time
	CreatedAtNe     *time.Time
	CreatedAtGt     *time.Time
	CreatedAtGte    *time.Time
	CreatedAtLt     *time.Time
	CreatedAtLte    *time.Time
	CreatedAtIn     *[]time.Time
	CreatedAtNin    *[]time.Time
	CreatedAtExists *bool
	// text (string) search options
	TextEq     *string
	TextNe     *string
	TextGt     *string
	TextGte    *string
	TextLt     *string
	TextLte    *string
	TextIn     *[]string
	TextNin    *[]string
	TextExists *bool
	TextLike   *string
	TextNlike  *string
}

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.AuthorIdEq != nil {
		elemauthorIdEq0 := o.AuthorIdEq
		to.AuthorIdEq = elemauthorIdEq0
	}
	if o.AuthorIdNe != nil {
		elemauthorIdNe0 := o.AuthorIdNe
		to.AuthorIdNe = elemauthorIdNe0
	}
	if o.AuthorIdGt != nil {
		elemauthorIdGt0 := o.AuthorIdGt
		to.AuthorIdGt = elemauthorIdGt0
	}
	if o.AuthorIdGte != nil {
		elemauthorIdGte0 := o.AuthorIdGte
		to.AuthorIdGte = elemauthorIdGte0
	}
	if o.AuthorIdLt != nil {
		elemauthorIdLt0 := o.AuthorIdLt
		to.AuthorIdLt = elemauthorIdLt0
	}
	if o.AuthorIdLte != nil {
		elemauthorIdLte0 := o.AuthorIdLte
		to.AuthorIdLte = elemauthorIdLte0
	}
	if o.AuthorIdIn != nil {
		elemauthorIdIn0 := make([]string, 0)
		for _, oauthorIdIn0 := range *o.AuthorIdIn {
			elemauthorIdIn1 := oauthorIdIn0
			elemauthorIdIn0 = append(elemauthorIdIn0, elemauthorIdIn1)
		}
		to.AuthorIdIn = &elemauthorIdIn0
	}
	if o.AuthorIdNin != nil {
		elemauthorIdNin0 := make([]string, 0)
		for _, oauthorIdNin0 := range *o.AuthorIdNin {
			elemauthorIdNin1 := oauthorIdNin0
			elemauthorIdNin0 = append(elemauthorIdNin0, elemauthorIdNin1)
		}
		to.AuthorIdNin = &elemauthorIdNin0
	}
	if o.AuthorIdExists != nil {
		elemauthorIdExists0 := o.AuthorIdExists
		to.AuthorIdExists = elemauthorIdExists0
	}
	if o.AuthorIdLike != nil {
		elemauthorIdLike0 := o.AuthorIdLike
		to.AuthorIdLike = elemauthorIdLike0
	}
	if o.AuthorIdNlike != nil {
		elemauthorIdNlike0 := o.AuthorIdNlike
		to.AuthorIdNlike = elemauthorIdNlike0
	}
	if o.CreatedAtEq != nil {
		elemcreatedAtEq0 := o.CreatedAtEq
		to.CreatedAtEq = elemcreatedAtEq0
	}
	if o.CreatedAtNe != nil {
		elemcreatedAtNe0 := o.CreatedAtNe
		to.CreatedAtNe = elemcreatedAtNe0
	}
	if o.CreatedAtGt != nil {
		elemcreatedAtGt0 := o.CreatedAtGt
		to.CreatedAtGt = elemcreatedAtGt0
	}
	if o.CreatedAtGte != nil {
		elemcreatedAtGte0 := o.CreatedAtGte
		to.CreatedAtGte = elemcreatedAtGte0
	}
	if o.CreatedAtLt != nil {
		elemcreatedAtLt0 := o.CreatedAtLt
		to.CreatedAtLt = elemcreatedAtLt0
	}
	if o.CreatedAtLte != nil {
		elemcreatedAtLte0 := o.CreatedAtLte
		to.CreatedAtLte = elemcreatedAtLte0
	}
	if o.CreatedAtIn != nil {
		elemcreatedAtIn0 := make([]time.Time, 0)
		for _, ocreatedAtIn0 := range *o.CreatedAtIn {
			elemcreatedAtIn1 := ocreatedAtIn0
			elemcreatedAtIn0 = append(elemcreatedAtIn0, elemcreatedAtIn1)
		}
		to.CreatedAtIn = &elemcreatedAtIn0
	}
	if o.CreatedAtNin != nil {
		elemcreatedAtNin0 := make([]time.Time, 0)
		for _, ocreatedAtNin0 := range *o.CreatedAtNin {
			elemcreatedAtNin1 := ocreatedAtNin0
			elemcreatedAtNin0 = append(elemcreatedAtNin0, elemcreatedAtNin1)
		}
		to.CreatedAtNin = &elemcreatedAtNin0
	}
	if o.CreatedAtExists != nil {
		elemcreatedAtExists0 := o.CreatedAtExists
		to.CreatedAtExists = elemcreatedAtExists0
	}
	if o.TextEq != nil {
		elemtextEq0 := o.TextEq
		to.TextEq = elemtextEq0
	}
	if o.TextNe != nil {
		elemtextNe0 := o.TextNe
		to.TextNe = elemtextNe0
	}
	if o.TextGt != nil {
		elemtextGt0 := o.TextGt
		to.TextGt = elemtextGt0
	}
	if o.TextGte != nil {
		elemtextGte0 := o.TextGte
		to.TextGte = elemtextGte0
	}
	if o.TextLt != nil {
		elemtextLt0 := o.TextLt
		to.TextLt = elemtextLt0
	}
	if o.TextLte != nil {
		elemtextLte0 := o.TextLte
		to.TextLte = elemtextLte0
	}
	if o.TextIn != nil {
		elemtextIn0 := make([]string, 0)
		for _, otextIn0 := range *o.TextIn {
			elemtextIn1 := otextIn0
			elemtextIn0 = append(elemtextIn0, elemtextIn1)
		}
		to.TextIn = &elemtextIn0
	}
	if o.TextNin != nil {
		elemtextNin0 := make([]string, 0)
		for _, otextNin0 := range *o.TextNin {
			elemtextNin1 := otextNin0
			elemtextNin0 = append(elemtextNin0, elemtextNin1)
		}
		to.TextNin = &elemtextNin0
	}
	if o.TextExists != nil {
		elemtextExists0 := o.TextExists
		to.TextExists = elemtextExists0
	}
	if o.TextLike != nil {
		elemtextLike0 := o.TextLike
		to.TextLike = elemtextLike0
	}
	if o.TextNlike != nil {
		elemtextNlike0 := o.TextNlike
		to.TextNlike = elemtextNlike0
	}
	return to, nil
}

type SortParams struct {
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	return to
}
