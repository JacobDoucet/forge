// This file is auto-generated. DO NOT EDIT.

import { EventSearchQuery, EventWithRefs, EventWithRefsProjection } from '../model/event-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Event, EventSortParams } from '../model/event-model';

export type SearchEventsParams = {
    baseUrl: string;
    query: EventSearchQuery;
    sort?: EventSortParams;
    projection?: EventWithRefsProjection;
    limit?: number;
    skip?: number;
}

async function newApiError(response: Response, defaultText: string): Promise<ApiError> {
    let text = defaultText;
    try {
        text = await response.text();
    } catch(_) {}
    return new ApiError(text);
}

export function searchEvents(params: SearchEventsParams): Promise<SelectManyResponse<EventWithRefs>> {
    return fetch(`${params.baseUrl}/events/search`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            sort: params.sort,
            projection: params.projection,
            limit: params.limit,
            skip: params.skip,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to search Event');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectEventByIdParams = {
    baseUrl: string;
    id: string;
    projection?: EventWithRefsProjection;
}

export function selectEventById(params: SelectEventByIdParams): Promise<EventWithRefs> {
    return fetch(`${params.baseUrl}/events/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Event');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveEventParams = {
    baseUrl: string;
    data: Event;
}

export function createEvent(params: SaveEventParams): Promise<MutationResponse<Event>> {
    return fetch(`${params.baseUrl}/events/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Event');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateEvent(params: SaveEventParams): Promise<MutationResponse<Event>> {
    return fetch(`${params.baseUrl}/events/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Event');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteEventParams = {
    baseUrl: string;
    id: string;
}

export function deleteEvent({ baseUrl, id }: DeleteEventParams): Promise<void> {
    return fetch(`${baseUrl}/events/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Event');
            return Promise.reject(err);
        }
        return;
    });
}
