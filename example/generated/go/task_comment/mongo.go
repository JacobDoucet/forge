package task_comment

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MongoRecord struct {
	AuthorId  *string    `bson:"authorId,omitempty"`
	CreatedAt *time.Time `bson:"createdAt,omitempty"`
	Text      *string    `bson:"text,omitempty"`
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.AuthorId != nil {
		elemauthorId0 := r.AuthorId
		m.AuthorId = *elemauthorId0
	}
	if r.CreatedAt != nil {
		elemcreatedAt0 := r.CreatedAt
		m.CreatedAt = *elemcreatedAt0
	}
	if r.Text != nil {
		elemtext0 := r.Text
		m.Text = *elemtext0
	}
	return m, nil
}

type MongoWhereClause struct {
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
	if o.AuthorIdEq != nil {
		query := bson.M{}
		query["authorId"] = o.AuthorIdEq
		and = append(and, query)
	}
	if o.AuthorIdNe != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$ne": o.AuthorIdNe}
		and = append(and, query)
	}
	if o.AuthorIdGt != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$gt": o.AuthorIdGt}
		and = append(and, query)
	}
	if o.AuthorIdGte != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$gte": o.AuthorIdGte}
		and = append(and, query)
	}
	if o.AuthorIdLt != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$lt": o.AuthorIdLt}
		and = append(and, query)
	}
	if o.AuthorIdLte != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$lte": o.AuthorIdLte}
		and = append(and, query)
	}
	if o.AuthorIdIn != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$in": o.AuthorIdIn}
		and = append(and, query)
	}
	if o.AuthorIdNin != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$nin": o.AuthorIdNin}
		and = append(and, query)
	}
	if o.AuthorIdExists != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$exists": *o.AuthorIdExists}
		and = append(and, query)
	}
	if o.AuthorIdLike != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$regex": o.AuthorIdLike, "$options": "i"}
		and = append(and, query)
	}
	if o.AuthorIdNlike != nil {
		query := bson.M{}
		query["authorId"] = bson.M{"$not": bson.M{"$regex": o.AuthorIdNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.CreatedAtEq != nil {
		query := bson.M{}
		query["createdAt"] = o.CreatedAtEq
		and = append(and, query)
	}
	if o.CreatedAtNe != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$ne": o.CreatedAtNe}
		and = append(and, query)
	}
	if o.CreatedAtGt != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$gt": o.CreatedAtGt}
		and = append(and, query)
	}
	if o.CreatedAtGte != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$gte": o.CreatedAtGte}
		and = append(and, query)
	}
	if o.CreatedAtLt != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$lt": o.CreatedAtLt}
		and = append(and, query)
	}
	if o.CreatedAtLte != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$lte": o.CreatedAtLte}
		and = append(and, query)
	}
	if o.CreatedAtIn != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$in": o.CreatedAtIn}
		and = append(and, query)
	}
	if o.CreatedAtNin != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$nin": o.CreatedAtNin}
		and = append(and, query)
	}
	if o.CreatedAtExists != nil {
		query := bson.M{}
		query["createdAt"] = bson.M{"$exists": *o.CreatedAtExists}
		and = append(and, query)
	}
	if o.TextEq != nil {
		query := bson.M{}
		query["text"] = o.TextEq
		and = append(and, query)
	}
	if o.TextNe != nil {
		query := bson.M{}
		query["text"] = bson.M{"$ne": o.TextNe}
		and = append(and, query)
	}
	if o.TextGt != nil {
		query := bson.M{}
		query["text"] = bson.M{"$gt": o.TextGt}
		and = append(and, query)
	}
	if o.TextGte != nil {
		query := bson.M{}
		query["text"] = bson.M{"$gte": o.TextGte}
		and = append(and, query)
	}
	if o.TextLt != nil {
		query := bson.M{}
		query["text"] = bson.M{"$lt": o.TextLt}
		and = append(and, query)
	}
	if o.TextLte != nil {
		query := bson.M{}
		query["text"] = bson.M{"$lte": o.TextLte}
		and = append(and, query)
	}
	if o.TextIn != nil {
		query := bson.M{}
		query["text"] = bson.M{"$in": o.TextIn}
		and = append(and, query)
	}
	if o.TextNin != nil {
		query := bson.M{}
		query["text"] = bson.M{"$nin": o.TextNin}
		and = append(and, query)
	}
	if o.TextExists != nil {
		query := bson.M{}
		query["text"] = bson.M{"$exists": *o.TextExists}
		and = append(and, query)
	}
	if o.TextLike != nil {
		query := bson.M{}
		query["text"] = bson.M{"$regex": o.TextLike, "$options": "i"}
		and = append(and, query)
	}
	if o.TextNlike != nil {
		query := bson.M{}
		query["text"] = bson.M{"$not": bson.M{"$regex": o.TextNlike, "$options": "i"}}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
}
