// This file is auto-generated. DO NOT EDIT.

import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessActorRole<T = ActorRole> = ActorCanAccessFunc<T> & {
    field: {
        role: ActorCanAccessFunc<ActorRole>;
    }
};

export const canReadActorRole = NewCanReadActorRole(
    (actorRoles: ActorRole[], obj?: ActorRole) => {
        return true;
    },
);

export const canWriteActorRole = NewCanWriteActorRole(
    (actorRoles: ActorRole[], obj?: ActorRole) => {
          return true;
    },
);

export function NewCanReadActorRole<T = ActorRole>(canAccessObj: ActorCanAccessFunc<T>): canAccessActorRole<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                role: (_actorRoles: ActorRole[], _obj?: ActorRole) =>  true,
            },
        },
    );
}

export function NewCanWriteActorRole<T = ActorRole>(canAccessObj: ActorCanAccessFunc<T>): canAccessActorRole<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                role: (_actorRoles: ActorRole[], _obj?: ActorRole) =>  true,
            },
        },
    );
}
