package integration_test

import (
 "testing"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/tenant_record"
)

func TestTenantQueryAuthorization(t *testing.T) {
 actor := &testActor{roles: []actor_role.Model{{Role:enum_role.TenantMember,TenantId:"tenant-a"}}}
 allowed, denied := "tenant-a", "tenant-b"
 for _, operation := range []struct{
  name string
  apply func(permissions.Actor,tenant_record.WhereClause)(tenant_record.WhereClause,error)
 }{
  {"read",tenant_record.ApplyActorReadPermissionsToWhereClause},
  {"write",tenant_record.ApplyActorWritePermissionsToWhereClause},
 } {
  t.Run(operation.name,func(t *testing.T){
   for _,tc:=range []struct{name string; query tenant_record.WhereClause; reject bool}{
    {"authorized equality",tenant_record.WhereClause{TenantIdEq:&allowed},false},
    {"unauthorized equality",tenant_record.WhereClause{TenantIdEq:&denied},true},
    {"authorized list",tenant_record.WhereClause{TenantIdIn:&[]string{allowed}},false},
    {"unauthorized list",tenant_record.WhereClause{TenantIdIn:&[]string{denied}},true},
    {"mixed list",tenant_record.WhereClause{TenantIdIn:&[]string{allowed,denied}},true},
   }{
    t.Run(tc.name,func(t *testing.T){
     _,err:=operation.apply(actor,tc.query)
     if (err!=nil)!=tc.reject{t.Fatalf("reject=%v, got error %v",tc.reject,err)}
    })
   }
   q,err:=operation.apply(actor,tenant_record.WhereClause{})
   if err!=nil||q.TenantIdIn==nil||len(*q.TenantIdIn)!=1||(*q.TenantIdIn)[0]!=allowed{t.Fatalf("unfiltered query not tenant scoped: %#v, %v",q,err)}
   if _,err:=operation.apply(&testActor{},tenant_record.WhereClause{});err==nil{t.Fatal("actor without tenant grant accepted")}
   if _,err:=operation.apply(&testActor{roles:[]actor_role.Model{{Role:enum_role.TenantMember}}},tenant_record.WhereClause{});err==nil{t.Fatal("missing tenant attribute accepted")}
  })
 }
}
