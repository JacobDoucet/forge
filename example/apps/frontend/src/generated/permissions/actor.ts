// This file is auto-generated. DO NOT EDIT.

import { ActorRole } from '../model/actor-role-model';

export type Actor = {
    actorId: string;
    actorLanguage: string;
    actorName: string;
    actorUsername: string;
    actorType: ;
    roles: ActorRole[];
}

type Role = NonNullable<ActorRole['role']>;

export function actorHasRole(a: Actor, roles: Role | Role[]): boolean {
    if (typeof roles === 'string') {
        roles = [roles];
    }
    const roleSet = new Set(roles);
    for (const r of a.roles) {
        if (r.role && roleSet.has(r.role)) {
            return true;
        }
    }
    return false;
}

export type ActorCanAccessFunc<T, R = T> = (actorRoles: ActorRole[], obj?: R) => boolean;
