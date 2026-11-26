package http_server

import (
 "net/http"
 "net/http/httptest"
 "testing"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/api"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/event_api"
 "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
)

func TestPartialRegistration(t *testing.T) {
 c := &api.CustomClient{}
 c.UseEventClient(event_api.New(&event_api.UnimplementedClient{}))
 props := ServeMuxProps{ResolveActor: func(*http.Request)(permissions.Actor,error){return permissions.Actor{},nil}}
 mux, err := ServeMux(c, props)
 if err != nil {t.Fatal(err)}
 for path, want := range map[string]string{"/events/":"/events/", "/users/":""} {
  _, pattern := mux.Handler(httptest.NewRequest("GET",path,nil))
  if pattern != want {t.Errorf("%s: got %q, want %q",path,pattern,want)}
 }
 props.RequireAllClients = true
 if _,err := ServeMux(c,props); err == nil {t.Fatal("expected missing client error")}
 props.RequireAllClients = false
 if _,err := ServeMux(&api.CustomClient{},props); err != nil {t.Fatal(err)}
 if _,err := ServeMux(&api.CustomClient{},ServeMuxProps{}); err == nil {t.Fatal("expected missing resolver error")}
}
