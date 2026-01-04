// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Event } from '../../model/event-model';
import { useFormState } from './common';
import {
    EventMutationOptions,
    useCreateEvent,
    useUpdateEvent,
} from '../tanstack-query/event-queries';
import { MutationResponse } from '../../api/model';

type UseEventFormStateOptions = {
    initialState: Event;
    onSuccess?: (res: MutationResponse<Event>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: EventMutationOptions;
};

export function useEventFormState(options: UseEventFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateEvent = useUpdateEvent(options.mutationOptions);
    const createEvent = useCreateEvent(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateEvent.mutate(formState.updates, opts);
        }
        return createEvent.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createEvent,
        updateEvent,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createEvent.isLoading || updateEvent.isLoading;

    return {
        ...formState,
        save,
        createMutation: createEvent,
        updateMutation: updateEvent,
        isLoading,
    } as const;
}
