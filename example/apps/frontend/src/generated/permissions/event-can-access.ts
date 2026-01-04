// This file is auto-generated. DO NOT EDIT.

import { Event } from '../model/event-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { NewCanReadEventSubject, NewCanWriteEventSubject } from './event-subject-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessEvent<T = Event> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Event>; 
        created: ReturnType<typeof NewCanReadActorTrace<Event>>, 
        subjects: ReturnType<typeof NewCanReadEventSubject<Event>>,
        type: ActorCanAccessFunc<Event>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Event>>, 
        updatedByUser: ReturnType<typeof NewCanReadActorTrace<Event>>,
    }
};

export const canReadEvent = NewCanReadEvent(
    (actorRoles: ActorRole[], obj?: Event) => {
        return true;
    },
);

export const canWriteEvent = NewCanWriteEvent(
    (actorRoles: ActorRole[], obj?: Event) => {
          return true;
    },
);

export function NewCanReadEvent<T = Event>(canAccessObj: ActorCanAccessFunc<T>): canAccessEvent<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                subjects:  NewCanReadEventSubject( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                type: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                updatedByUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
            },
        },
    );
}

export function NewCanWriteEvent<T = Event>(canAccessObj: ActorCanAccessFunc<T>): canAccessEvent<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                subjects:  NewCanWriteEventSubject( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                type: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                updatedByUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
            },
        },
    );
}
