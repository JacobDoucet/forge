// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Event } from '../../model/event-model';
import { EventWithRefs } from '../../model/event-api';
import {
    searchEvents, SearchEventsParams,
    selectEventById, SelectEventByIdParams,
    createEvent, updateEvent, deleteEvent,
} from '../../api/event-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchEventsProps = Omit<SearchEventsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<EventWithRefs>,
    ApiError,
    SelectManyResponse<EventWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchEvents(
    { queryKey, queryName, ...params }: UseSearchEventsProps,
    queryOptions?: SearchQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const keys = Object.keys(params.query);
        keys.sort();
        const searchKey = keys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof SearchEventsParams['query']])}`
        );
        return ['searchEvents', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchEvents({ baseUrl, ...params }),
    });
}
type UseSelectEventByIdProps = Omit<SelectEventByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectEventByIdOptions = Omit<UseQueryOptions<
    EventWithRefs,
    ApiError,
    EventWithRefs,
    any[]
>, 'initialData'>;

export function useSelectEventById(
    { queryKey, queryName, ...params }: UseSelectEventByIdProps,
    queryOptions?: SelectEventByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectEventById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectEventById({ baseUrl, ...params }),
    });
}

export type EventMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateEvent(options: EventMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Event>, ApiError, Event>(async (event: Event) => {
        const res = await createEvent({ baseUrl, data: event });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateEvent(options: EventMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Event>, ApiError, Event>(async (event: Event) => {
        const res = await updateEvent({ baseUrl, data: event });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteEvent(options: EventMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteEvent({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
