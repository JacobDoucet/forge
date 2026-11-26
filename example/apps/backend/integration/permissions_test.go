package integration_test

import (
 "context"
 "errors"
 "testing"

 "github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/project"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/task"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/task_api"
)

type testActor struct {
 permissions.SuperActor
 roles []actor_role.Model
}
func (a *testActor) GetActorRoles() []actor_role.Model {return a.roles}
func (a *testActor) GetRoleMap() permissions.RoleMap {return permissions.BuildRoleMap(a.roles)}

func TestGeneratedRBAC(t *testing.T) {
 for _, tc := range []struct{role enum_role.Value; read, taskWrite, projectWrite bool}{
  {enum_role.Admin,true,true,true},
  {enum_role.User,true,true,false},
  {enum_role.Guest,true,false,false},
  {"",false,false,false},
 } {
  a := &testActor{roles:[]actor_role.Model{{Role:tc.role}}}
  if _,err := task.ApplyActorReadPermissionsToWhereClause(a,task.WhereClause{}); (err==nil)!=tc.read {t.Errorf("%s read: %v",tc.role,err)}
  if _,err := task.ApplyActorWritePermissionsToWhereClause(a,task.WhereClause{}); (err==nil)!=tc.taskWrite {t.Errorf("%s task write: %v",tc.role,err)}
  if got:=task.HasWritePermissions(task.Model{},a);got!=tc.taskWrite {t.Errorf("%s task object write: %v",tc.role,got)}
  if got:=project.HasWritePermissions(project.Model{},a);got!=tc.projectWrite {t.Errorf("%s project write: %v",tc.role,got)}
 }
}

type searchSpy struct {
 task_api.UnimplementedClient
 calls int
 query task_api.WhereClause
}
func (s *searchSpy) Search(_ context.Context,q task_api.WhereClause,_ task_api.QueryOptions)(task_api.QueryResult,error){
 s.calls++
 s.query=q
 return task_api.QueryResult{},nil
}

func TestSearchPermissionAndHooks(t *testing.T){
 spy:= &searchSpy{}
 title:="example"
 before,after:=0,0
 c:=task_api.New(spy,task_api.Hooks{
  OnBeforeSearch:func(_ context.Context,_ permissions.Actor,q task.WhereClause,o task_api.QueryOptions)(task.WhereClause,task_api.QueryOptions,error){
   before++;q.TitleEq=&title;return q,o,nil
  },
  OnSearch:func(_ context.Context,_ permissions.Actor,_ task_api.QueryResult,_ task_api.Projection,err error)error{after++;return err},
 })
 if _,_,err:=c.Search(context.Background(),&testActor{},task.WhereClause{},task_api.QueryOptions{});err==nil||spy.calls!=0||after!=0{t.Fatalf("unauthorized search reached backend: %v",err)}
 if _,_,err:=c.Search(context.Background(),permissions.NewSuperActor(),task.WhereClause{},task_api.QueryOptions{});err!=nil{t.Fatal(err)}
 if before!=2||after!=1||spy.calls!=1||spy.query.Task.TitleEq==nil||*spy.query.Task.TitleEq!=title{t.Fatal("hook or query propagation mismatch")}
 sentinel:=errors.New("reject")
 c=task_api.New(spy,task_api.Hooks{OnBeforeSearch:func(_ context.Context,_ permissions.Actor,q task.WhereClause,o task_api.QueryOptions)(task.WhereClause,task_api.QueryOptions,error){return q,o,sentinel}})
 if _,_,err:=c.Search(context.Background(),permissions.NewSuperActor(),task.WhereClause{},task_api.QueryOptions{});!errors.Is(err,sentinel)||spy.calls!=1{t.Fatal("hook rejection did not stop backend")}
}

func TestMongoQuery(t *testing.T){
 title:="example"
 q,err:=(task.MongoWhereClause{TitleEq:&title}).GetLookupQuery()
 if err!=nil{t.Fatal(err)}
 if len(q)!=1||q["$and"]==nil{t.Fatalf("unexpected query: %#v",q)}
}
