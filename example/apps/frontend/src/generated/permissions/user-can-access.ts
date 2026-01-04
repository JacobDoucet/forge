// This file is auto-generated. DO NOT EDIT.

import { User } from '../model/user-model';
import { NewCanReadActorRole, NewCanWriteActorRole } from './actor-role-can-access';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessUser<T = User> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<User>; 
        actorRoles: ReturnType<typeof NewCanReadActorRole<User>>, 
        created: ReturnType<typeof NewCanReadActorTrace<User>>,
        email: ActorCanAccessFunc<User>;
        firstName: ActorCanAccessFunc<User>;
        language: ActorCanAccessFunc<User>;
        lastName: ActorCanAccessFunc<User>;
        role: ActorCanAccessFunc<User>; 
        updated: ReturnType<typeof NewCanReadActorTrace<User>>, 
        updatedByUser: ReturnType<typeof NewCanReadActorTrace<User>>,
    }
};

const getAbacActorId = (obj: User) => obj.id;

export const canReadUser = NewCanReadUser(
    (actorRoles: ActorRole[], obj?: User) => {
        for (const actorRole of actorRoles) {
            switch(actorRole.role) {
            case 'Super':
                return true;
            case 'Admin':
                return true;
            case 'User':
                return true;
            }
        }
        return false;
    },
);

export const canWriteUser = NewCanWriteUser(
    (actorRoles: ActorRole[], obj?: User) => {
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

export function NewCanReadUser<T = User>(canAccessObj: ActorCanAccessFunc<T>): canAccessUser<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                actorRoles:  NewCanReadActorRole( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                email: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                firstName: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                language: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                lastName: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                role: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                updatedByUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
            },
        },
    );
}

export function NewCanWriteUser<T = User>(canAccessObj: ActorCanAccessFunc<T>): canAccessUser<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                actorRoles:  NewCanWriteActorRole( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                email: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                firstName: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                language: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                lastName: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                role: (_actorRoles: ActorRole[], _obj?: User) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
                updatedByUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: User) =>  true),
            },
        },
    );
}
