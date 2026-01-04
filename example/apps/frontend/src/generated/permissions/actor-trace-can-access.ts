// This file is auto-generated. DO NOT EDIT.

import { ActorTrace } from '../model/actor-trace-model';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessActorTrace<T = ActorTrace> = ActorCanAccessFunc<T> & {
    field: {
        actorId: ActorCanAccessFunc<ActorTrace>;
        actorName: ActorCanAccessFunc<ActorTrace>;
        actorType: ActorCanAccessFunc<ActorTrace>;
        at: ActorCanAccessFunc<ActorTrace>;
    }
};

export const canReadActorTrace = NewCanReadActorTrace(
    (actorRoles: ActorRole[], obj?: ActorTrace) => {
        return true;
    },
);

export const canWriteActorTrace = NewCanWriteActorTrace(
    (actorRoles: ActorRole[], obj?: ActorTrace) => {
          return true;
    },
);

export function NewCanReadActorTrace<T = ActorTrace>(canAccessObj: ActorCanAccessFunc<T>): canAccessActorTrace<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                actorId: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                actorName: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                actorType: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                at: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
            },
        },
    );
}

export function NewCanWriteActorTrace<T = ActorTrace>(canAccessObj: ActorCanAccessFunc<T>): canAccessActorTrace<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                actorId: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                actorName: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                actorType: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
                at: (_actorRoles: ActorRole[], _obj?: ActorTrace) =>  true,
            },
        },
    );
}
