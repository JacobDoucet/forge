// This file is auto-generated. DO NOT EDIT.

import { Project } from '../model/project-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessProject<T = Project> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Project>; 
        created: ReturnType<typeof NewCanReadActorTrace<Project>>,
        description: ActorCanAccessFunc<Project>;
        name: ActorCanAccessFunc<Project>;
        ownerId: ActorCanAccessFunc<Project>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Project>>,
    }
};

export const canReadProject = NewCanReadProject(
    (actorRoles: ActorRole[], obj?: Project) => {
        for (const actorRole of actorRoles) {
            switch(actorRole.role) {
            case 'Super':
                return true;
            case 'Admin':
                return true;
            case 'Guest':
                return true;
            case 'User':
                return true;
            }
        }
        return false;
    },
);

export const canWriteProject = NewCanWriteProject(
    (actorRoles: ActorRole[], obj?: Project) => {
          for (const actorRole of actorRoles) {
              switch(actorRole.role) {
              case 'Super':
                  return true;
              case 'Admin':
                  return true;
              }
          }
          return false;
    },
);

export function NewCanReadProject<T = Project>(canAccessObj: ActorCanAccessFunc<T>): canAccessProject<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Project) =>  true),
                description: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Project) =>  true),
            },
        },
    );
}

export function NewCanWriteProject<T = Project>(canAccessObj: ActorCanAccessFunc<T>): canAccessProject<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Project) =>  true),
                description: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Project) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Project) =>  true),
            },
        },
    );
}
