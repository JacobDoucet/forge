// This file is auto-generated. DO NOT EDIT.

import { Task } from '../model/task-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { NewCanReadTaskComment, NewCanWriteTaskComment } from './task-comment-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessTask<T = Task> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Task>;
        assigneeId: ActorCanAccessFunc<Task>; 
        comments: ReturnType<typeof NewCanReadTaskComment<Task>>, 
        created: ReturnType<typeof NewCanReadActorTrace<Task>>,
        description: ActorCanAccessFunc<Task>;
        dueDate: ActorCanAccessFunc<Task>;
        priority: ActorCanAccessFunc<Task>;
        status: ActorCanAccessFunc<Task>;
        tags: ActorCanAccessFunc<Task>;
        title: ActorCanAccessFunc<Task>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Task>>, 
        updatedByUser: ReturnType<typeof NewCanReadActorTrace<Task>>,
    }
};

export const canReadTask = NewCanReadTask(
    (actorRoles: ActorRole[], obj?: Task) => {
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

export const canWriteTask = NewCanWriteTask(
    (actorRoles: ActorRole[], obj?: Task) => {
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

export function NewCanReadTask<T = Task>(canAccessObj: ActorCanAccessFunc<T>): canAccessTask<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                assigneeId: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                comments:  NewCanReadTaskComment( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                description: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                dueDate: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                priority: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                tags: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                updatedByUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
            },
        },
    );
}

export function NewCanWriteTask<T = Task>(canAccessObj: ActorCanAccessFunc<T>): canAccessTask<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                assigneeId: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                comments:  NewCanWriteTaskComment( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                description: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                dueDate: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                priority: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                tags: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: Task) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
                updatedByUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Task) =>  true),
            },
        },
    );
}
