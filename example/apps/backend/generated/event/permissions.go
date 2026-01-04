package event;

import (
    "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
)




func GetAbacProjection(actor permissions.Actor) Projection {
    return Projection{
    }
}

func HasWritePermissions(m Model, actor permissions.Actor) bool {
    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        }
    }
    return false
}

func ProjectReadPermissions(p Projection, actor permissions.Actor) Projection {
    return p
}

func ProjectWritePermissions(p Projection, actor permissions.Actor) Projection {
    return p
}

// Permissions on query



func ApplyActorReadPermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
    

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        }
    }
    return query, nil

}

func ApplyActorWritePermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
    

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        }
    }
    return query, nil

}
