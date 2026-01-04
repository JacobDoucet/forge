package task_mongo

import (
	"github.com/JacobDoucet/forge/example/generated/go/task"
)

type Model struct {
	task.MongoRecord `bson:",inline"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
