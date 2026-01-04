// This file is auto-generated. DO NOT EDIT.

import { Event, EventProjection } from './event-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { EventSubjectSearchQuery } from './event-subject-api';
import { EventType } from './event-type-enum';

export type EventWithRefs = {
    event: Event;
}

export type EventWithRefsProjection = EventProjection & {
}

export type SelectEventByIdQuery = {
    id: string;
}

export type EventSearchQuery = {
    // id (Ref<Event>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // subjects (List<EventSubject>) search options
    subjects?: EventSubjectSearchQuery;
    subjectsEmpty?: boolean;
    // type (EventType) search options
    typeEq?: EventType;
    typeNe?: EventType;
    typeGt?: EventType;
    typeGte?: EventType;
    typeLt?: EventType;
    typeLte?: EventType;
    typeIn?: EventType[];
    typeNin?: EventType[];
    typeExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByUser (ActorTrace) search options
    updatedByUser?: ActorTraceSearchQuery;
}
