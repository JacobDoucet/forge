package permissions

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
)

func NewSuperActor() Actor {
	return &SuperActor{
		Name:      "Super Admin",
		Username:  "super",
		AdminName: "Super Admin",
	}
}

type SuperActor struct {
	Name      string
	Username  string
	AdminName string
}

func (a *SuperActor) GetActorLanguage() string {
	return "en"
}

func (a *SuperActor) ActorType() ActorType {
	return ActorTypeSuper
}

func (a *SuperActor) GetActorName() string {
	return a.Name
}

func (a *SuperActor) GetActorUsername() string {
	return a.Username
}

func (a *SuperActor) GetActorAdminName() string {
	return a.AdminName
}

func (a *SuperActor) Actor() Actor {
	return a
}

func (a *SuperActor) GetActorRoles() []actor_role.Model {
	return []actor_role.Model{
		{Role: enum_role.Super},
	}
}

func (a *SuperActor) GetRoleMap() RoleMap {
	return BuildRoleMap(a.GetActorRoles())
}
func (a *SuperActor) GetActorId() string {
	// TODO
	return ""
}
