// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';

export type Project = {
  id?: string;
  created?: ActorTrace;
  description?: string;
  name?: string;
  ownerId?: string;
  updated?: ActorTrace;
  updatedByUser?: ActorTrace;
}

export type ProjectProjection = {
    id?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    description?: boolean;
    name?: boolean;
    ownerId?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByUser?: boolean;
		updatedByUserFields?: ActorTraceProjection;
}

export type ProjectSortParams = {
    createdAt?: -1 | 1;
    updatedAt?: -1 | 1;
}
