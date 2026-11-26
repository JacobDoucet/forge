package types

import (
 "testing"
 "gopkg.in/yaml.v3"
)

func TestHTTPMethods(t *testing.T) {
 for _, s := range []string{"GET","post","Put","DELETE","patch"} {
  if _,err := sanitizeHTTPMethod(s); err != nil {t.Errorf("%s: %v",s,err)}
 }
 for _, s := range []string{"","TRACE","invalid"} {
  if _,err := sanitizeHTTPMethod(s); err == nil {t.Errorf("accepted %q",s)}
 }
}

func TestDuplicateEndpointsAndCollections(t *testing.T) {
 e := newEndpointRegistry()
 if err := e.Register(ObjectHTTPDef{Endpoint:"tasks"}); err != nil {t.Fatal(err)}
 if err := e.Register(ObjectHTTPDef{Endpoint:"tasks"}); err == nil {t.Fatal("duplicate endpoint accepted")}
 c := newCollectionRegistry()
 if err := c.Register(Collection{Name:"tasks"}); err != nil {t.Fatal(err)}
 if err := c.Register(Collection{Name:"tasks"}); err == nil {t.Fatal("duplicate collection accepted")}
}

func TestPermissionSpecValidation(t *testing.T) {
 r := newPermissionRegistry()
 if err := r.RegisterRbac(PermissionsRbacDef{Name:"Member"}); err != nil {t.Fatal(err)}
 for _, tc := range []struct{name, spec string; valid bool}{
  {"valid","read: [{rbac: Member, abac: [ActorId]}]",true},
  {"unknown role","read: [{rbac: Missing}]",false},
  {"unknown attribute","write: [{rbac: Member, abac: [Missing]}]",false},
  {"empty rule","write: [{}]",false},
 } {
  t.Run(tc.name,func(t *testing.T){
   var p ObjectPermissions
   if err := yaml.Unmarshal([]byte(tc.spec), &p); err != nil {t.Fatal(err)}
   if err := r.ValidateObjectPermissions("Task",p); (err == nil) != tc.valid {t.Fatalf("valid=%v, error=%v",tc.valid,err)}
  })
 }
}
