// This file is auto-generated. DO NOT EDIT.

import { User, UserProjection } from './user-model';
import { ActorRoleSearchQuery } from './actor-role-api';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { Role } from './role-enum';

export type UserWithRefs = {
    user: User;
}

export type UserWithRefsProjection = UserProjection & {
}

export type SelectUserByIdQuery = {
    id: string;
}

export type SelectUserByEmailIdxQuery = {
    email: string;
}

export type UserSearchQuery = {
    // id (Ref<User>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // actorRoles (List<ActorRole>) search options
    actorRoles?: ActorRoleSearchQuery;
    actorRolesEmpty?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // email (string) search options
    emailEq?: string;
    emailNe?: string;
    emailGt?: string;
    emailGte?: string;
    emailLt?: string;
    emailLte?: string;
    emailIn?: string[];
    emailNin?: string[];
    emailExists?: boolean;
    emailLike?: string;
    emailNlike?: string;
    // firstName (string) search options
    firstNameEq?: string;
    firstNameNe?: string;
    firstNameGt?: string;
    firstNameGte?: string;
    firstNameLt?: string;
    firstNameLte?: string;
    firstNameIn?: string[];
    firstNameNin?: string[];
    firstNameExists?: boolean;
    firstNameLike?: string;
    firstNameNlike?: string;
    // language (string) search options
    languageEq?: string;
    languageNe?: string;
    languageGt?: string;
    languageGte?: string;
    languageLt?: string;
    languageLte?: string;
    languageIn?: string[];
    languageNin?: string[];
    languageExists?: boolean;
    languageLike?: string;
    languageNlike?: string;
    // lastName (string) search options
    lastNameEq?: string;
    lastNameNe?: string;
    lastNameGt?: string;
    lastNameGte?: string;
    lastNameLt?: string;
    lastNameLte?: string;
    lastNameIn?: string[];
    lastNameNin?: string[];
    lastNameExists?: boolean;
    lastNameLike?: string;
    lastNameNlike?: string;
    // role (Role) search options
    roleEq?: Role;
    roleNe?: Role;
    roleGt?: Role;
    roleGte?: Role;
    roleLt?: Role;
    roleLte?: Role;
    roleIn?: Role[];
    roleNin?: Role[];
    roleExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByUser (ActorTrace) search options
    updatedByUser?: ActorTraceSearchQuery;
}
