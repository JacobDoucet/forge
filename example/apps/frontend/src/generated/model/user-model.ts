// This file is auto-generated. DO NOT EDIT.

import { ActorRole, ActorRoleProjection } from './actor-role-model';
import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { Role } from './role-enum';

export type User = {
  id?: string;
  actorRoles?: ActorRole[];
  created?: ActorTrace;
  email?: string;
  firstName?: string;
  language?: string;
  lastName?: string;
  role?: Role;
  updated?: ActorTrace;
  updatedByUser?: ActorTrace;
}

export function userAsActor(user: User) {

}

export type UserProjection = {
    id?: boolean;
    actorRoles?: boolean;
		actorRolesFields?: ActorRoleProjection;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    email?: boolean;
    firstName?: boolean;
    language?: boolean;
    lastName?: boolean;
    role?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByUser?: boolean;
		updatedByUserFields?: ActorTraceProjection;
}

export type UserSortParams = {
    createdAt?: -1 | 1;
    email?: -1 | 1;
    updatedAt?: -1 | 1;
}
