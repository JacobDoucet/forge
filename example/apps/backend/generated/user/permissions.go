package user;

import (
    "github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
    "github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
    "github.com/JacobDoucet/forge/example/apps/backend/generated/coded_error"
    "github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
    "fmt"
)




func (m *Model) Actor() permissions.Actor {
    return m
}

func (m *Model) ActorType() permissions.ActorType {
    return permissions.ActorTypeUser
}

func (m *Model) GetActorLanguage() string {
    return fmt.Sprintf("Unnamed Actor", )
}

func (m *Model) GetActorName() string {
    return fmt.Sprintf("%s %s", m.FirstName, m.LastName)
}

func (m *Model) GetActorUsername() string {
    return fmt.Sprintf("%s", m.Email)
}

func (m *Model) GetActorAdminName() string {
    return fmt.Sprintf("%s", m.Email)
}

func (m *Model) GetActorRoles() []actor_role.Model {
    return m.ActorRoles
}

func (m *Model) GetRoleMap() permissions.RoleMap {
    return permissions.BuildRoleMap(m.GetActorRoles())
}


func (m *Model) GetActorId() string {
    return m.Id
}


func GetAbacProjection(actor permissions.Actor) Projection {
    return Projection{
        Id: true,
    }
}

func HasWritePermissions(m Model, actor permissions.Actor) bool {
    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Super: 
            // Actor has full permissions
            return true
        case enum_role.Admin: 
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
     
    actorIdIn := make(map[string]struct{})

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        case enum_role.Admin: 
            // Actor has full permissions
            return query, nil
        case enum_role.User: 
            // Actor has full permissions
            return query, nil
        }
    }
    if len(actorIdIn) == 0 {
        // If actor does not have granted permissions, return error 
        if query.IdEq != nil {
            if *query.IdEq != actor.GetActorId() {
                return query, coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), *query.IdEq))
            }
            return query, nil
        }
        return query, coded_error.NewUnauthorizedError("actor has no permissions")
    }
     
    if err := func() error {
        // If no query, return
        if len(actorIdIn) == 0 {
            return nil
        }
        // If Eq the query, check if the actor has permissions
        if query.IdEq != nil {
            if _, ok := actorIdIn[*query.IdEq]; !ok {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), *query.IdEq))
            }
            return nil
        }
        // If filtering on a range of values, ensure the actor has permissions for all values
        if query.IdIn != nil {
           l := *query.IdIn
           var noPerms []string
           for _, v := range l {
                if _, ok := actorIdIn[v]; !ok {
                    noPerms = append(noPerms, v)
                }
           }
           if len(noPerms) == 0 {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), noPerms))
           }
           return nil
        }
        // Only include values the actor has permissions for
        l := make([]string, 0, len(actorIdIn))
        for k := range actorIdIn {
            l = append(l, k)
        }
        query.IdIn = &l
        return nil
    }(); err != nil {
        return query, err
    }
    
    return query, nil


}

func ApplyActorWritePermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
     
    actorIdIn := make(map[string]struct{})

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        case enum_role.Admin: 
            // Actor has full permissions
            return query, nil
        }
    }
    if len(actorIdIn) == 0 {
        // If actor does not have granted permissions, return error 
        if query.IdEq != nil {
            if *query.IdEq != actor.GetActorId() {
                return query, coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), *query.IdEq))
            }
            return query, nil
        }
        return query, coded_error.NewUnauthorizedError("actor has no permissions")
    }
     
    if err := func() error {
        // If no query, return
        if len(actorIdIn) == 0 {
            return nil
        }
        // If Eq the query, check if the actor has permissions
        if query.IdEq != nil {
            if _, ok := actorIdIn[*query.IdEq]; !ok {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), *query.IdEq))
            }
            return nil
        }
        // If filtering on a range of values, ensure the actor has permissions for all values
        if query.IdIn != nil {
           l := *query.IdIn
           var noPerms []string
           for _, v := range l {
                if _, ok := actorIdIn[v]; !ok {
                    noPerms = append(noPerms, v)
                }
           }
           if len(noPerms) == 0 {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to Id %v", actor.GetActorId(), noPerms))
           }
           return nil
        }
        // Only include values the actor has permissions for
        l := make([]string, 0, len(actorIdIn))
        for k := range actorIdIn {
            l = append(l, k)
        }
        query.IdIn = &l
        return nil
    }(); err != nil {
        return query, err
    }
    
    return query, nil


}
