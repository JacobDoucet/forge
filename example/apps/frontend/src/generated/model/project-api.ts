// This file is auto-generated. DO NOT EDIT.

import { Project, ProjectProjection } from './project-model';
import { ActorTraceSearchQuery } from './actor-trace-api';

export type ProjectWithRefs = {
    project: Project;
}

export type ProjectWithRefsProjection = ProjectProjection & {
}

export type SelectProjectByIdQuery = {
    id: string;
}

export type ProjectSearchQuery = {
    // id (Ref<Project>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // description (string) search options
    descriptionEq?: string;
    descriptionNe?: string;
    descriptionGt?: string;
    descriptionGte?: string;
    descriptionLt?: string;
    descriptionLte?: string;
    descriptionIn?: string[];
    descriptionNin?: string[];
    descriptionExists?: boolean;
    descriptionLike?: string;
    descriptionNlike?: string;
    // name (string) search options
    nameEq?: string;
    nameNe?: string;
    nameGt?: string;
    nameGte?: string;
    nameLt?: string;
    nameLte?: string;
    nameIn?: string[];
    nameNin?: string[];
    nameExists?: boolean;
    nameLike?: string;
    nameNlike?: string;
    // ownerId (string) search options
    ownerIdEq?: string;
    ownerIdNe?: string;
    ownerIdGt?: string;
    ownerIdGte?: string;
    ownerIdLt?: string;
    ownerIdLte?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    ownerIdLike?: string;
    ownerIdNlike?: string;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
}
