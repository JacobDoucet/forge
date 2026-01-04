// This file is auto-generated. DO NOT EDIT.

import { TaskComment } from '../model/task-comment-model';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessTaskComment<T = TaskComment> = ActorCanAccessFunc<T> & {
    field: {
        authorId: ActorCanAccessFunc<TaskComment>;
        createdAt: ActorCanAccessFunc<TaskComment>;
        text: ActorCanAccessFunc<TaskComment>;
    }
};

export const canReadTaskComment = NewCanReadTaskComment(
    (actorRoles: ActorRole[], obj?: TaskComment) => {
        return true;
    },
);

export const canWriteTaskComment = NewCanWriteTaskComment(
    (actorRoles: ActorRole[], obj?: TaskComment) => {
          return true;
    },
);

export function NewCanReadTaskComment<T = TaskComment>(canAccessObj: ActorCanAccessFunc<T>): canAccessTaskComment<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                authorId: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
                createdAt: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
                text: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
            },
        },
    );
}

export function NewCanWriteTaskComment<T = TaskComment>(canAccessObj: ActorCanAccessFunc<T>): canAccessTaskComment<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                authorId: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
                createdAt: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
                text: (_actorRoles: ActorRole[], _obj?: TaskComment) =>  true,
            },
        },
    );
}
