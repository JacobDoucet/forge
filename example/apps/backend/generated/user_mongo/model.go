package user_mongo

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/user"
)

type Model struct {
	user.MongoRecord `bson:",inline"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
