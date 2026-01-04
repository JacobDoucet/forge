package project_mongo

import (
	"github.com/JacobDoucet/forge/example/generated/go/project"
)

type Model struct {
	project.MongoRecord `bson:",inline"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
