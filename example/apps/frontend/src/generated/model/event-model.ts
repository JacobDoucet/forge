// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { EventSubject, EventSubjectProjection } from './event-subject-model';
import { EventType } from './event-type-enum';

export type Event = {
  id?: string;
  created?: ActorTrace;
  subjects?: EventSubject[];
  type?: EventType;
  updated?: ActorTrace;
  updatedByUser?: ActorTrace;
}

export type EventProjection = {
    id?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    subjects?: boolean;
		subjectsFields?: EventSubjectProjection;
    type?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByUser?: boolean;
		updatedByUserFields?: ActorTraceProjection;
}

export type EventSortParams = {
    createdAt?: -1 | 1;
    subjectsSubjectId?: -1 | 1;
    type?: -1 | 1;
    updatedAt?: -1 | 1;
}
