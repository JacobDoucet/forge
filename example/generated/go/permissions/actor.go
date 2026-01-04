package permissions

import (
	"errors"
	"github.com/JacobDoucet/forge/example/generated/go/actor_role"
	"github.com/JacobDoucet/forge/example/generated/go/actor_trace"
	"github.com/JacobDoucet/forge/example/generated/go/enum_role"
	"time"
)

type ActorType string

const (
	ActorTypeSuper ActorType = "Super"
)

type Actor interface {
	Actor() Actor
	GetActorLanguage() string
	GetActorName() string
	GetActorUsername() string
	GetActorAdminName() string
	ActorType() ActorType
	GetActorRoles() []actor_role.Model
	GetRoleMap() RoleMap
	GetActorId() string
}

type HTTPActor struct {
	ActorId       string                  `json:"actorId"`
	ActorLanguage string                  `json:"actorLanguage"`
	ActorUsername string                  `json:"actorUsername"`
	ActorName     string                  `json:"actorName"`
	ActorType     string                  `json:"actorType"`
	Roles         []actor_role.HTTPRecord `json:"roles"`
}

func NewHTTPActor(actor Actor) HTTPActor {
	roles := actor.GetActorRoles()
	var httpRoles []actor_role.HTTPRecord

	for _, role := range roles {
		httpRole, _ := role.ToHTTPRecord(actor_role.Projection{
			Role: true,
		})
		httpRoles = append(httpRoles, httpRole)
	}

	return HTTPActor{
		ActorId:       actor.GetActorId(),
		ActorLanguage: actor.GetActorLanguage(),
		ActorUsername: actor.GetActorUsername(),
		ActorName:     actor.GetActorName(),
		ActorType:     string(actor.ActorType()),
		Roles:         httpRoles,
	}
}

func Trace(a Actor) actor_trace.Model {
	return actor_trace.Model{
		ActorId:   a.GetActorId(),
		ActorName: a.GetActorName(),
		ActorType: string(a.ActorType()),
		At:        time.Now(),
	}
}

type ActorRole struct {
	Role    enum_role.Value `json:"role" bson:"role"`
	ActorId *string         `json:"actorId" bson:"actorId"`
}

type RoleMap map[enum_role.Value]bool

func BuildRoleMap(roles []actor_role.Model) RoleMap {
	roleMap := make(RoleMap)
	for _, role := range roles {
		roleMap[role.Role] = true
	}
	return roleMap
}

func ValidateActorRole(ar actor_role.Model) error {
	switch ar.Role {
	case enum_role.Super:
		return nil
	case enum_role.Admin:
		return nil
	case enum_role.Guest:
		return nil
	case enum_role.User:
		return nil
	}
	return errors.New("invalid role " + string(ar.Role))
}

func NewRoleSuper() ActorRole {
	return ActorRole{
		Role: enum_role.Super,
	}
}

func NewRoleAdmin() ActorRole {
	return ActorRole{
		Role: enum_role.Admin,
	}
}

func NewRoleGuest() ActorRole {
	return ActorRole{
		Role: enum_role.Guest,
	}
}

func NewRoleUser() ActorRole {
	return ActorRole{
		Role: enum_role.User,
	}
}
