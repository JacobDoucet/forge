package project;

import (
    "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
    "github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
    "github.com/JacobDoucet/forge/example/apps/backend/generated/coded_error"
)




func GetAbacProjection(actor permissions.Actor) Projection {
    return Projection{
    }
}

func HasWritePermissions(m Model, actor permissions.Actor) bool {
    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Admin: 
            // Actor has full permissions
            return true
        case enum_role.Super: 
            // Actor has full permissions
            return true
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
        case enum_role.Admin: 
            // Actor has full permissions
            return query, nil
        case enum_role.Guest: 
            // Actor has full permissions
            return query, nil
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        case enum_role.User: 
            // Actor has full permissions
            return query, nil
        }
    }
    return query, coded_error.NewUnauthorizedError()

}

func ApplyActorWritePermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
    

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Admin: 
            // Actor has full permissions
            return query, nil
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        }
    }
    return query, coded_error.NewUnauthorizedError()

}
