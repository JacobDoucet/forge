// This file is auto-generated. DO NOT EDIT.

import { EventSubject } from '../model/event-subject-model';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessEventSubject<T = EventSubject> = ActorCanAccessFunc<T> & {
    field: {
        subjectId: ActorCanAccessFunc<EventSubject>;
        subjectType: ActorCanAccessFunc<EventSubject>;
    }
};

export const canReadEventSubject = NewCanReadEventSubject(
    (actorRoles: ActorRole[], obj?: EventSubject) => {
        return true;
    },
);

export const canWriteEventSubject = NewCanWriteEventSubject(
    (actorRoles: ActorRole[], obj?: EventSubject) => {
          return true;
    },
);

export function NewCanReadEventSubject<T = EventSubject>(canAccessObj: ActorCanAccessFunc<T>): canAccessEventSubject<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                subjectId: (_actorRoles: ActorRole[], _obj?: EventSubject) =>  true,
                subjectType: (_actorRoles: ActorRole[], _obj?: EventSubject) =>  true,
            },
        },
    );
}

export function NewCanWriteEventSubject<T = EventSubject>(canAccessObj: ActorCanAccessFunc<T>): canAccessEventSubject<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                subjectId: (_actorRoles: ActorRole[], _obj?: EventSubject) =>  true,
                subjectType: (_actorRoles: ActorRole[], _obj?: EventSubject) =>  true,
            },
        },
    );
}
