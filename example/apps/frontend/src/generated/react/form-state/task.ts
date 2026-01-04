// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Task } from '../../model/task-model';
import { useFormState } from './common';
import {
    TaskMutationOptions,
    useCreateTask,
    useUpdateTask,
} from '../tanstack-query/task-queries';
import { MutationResponse } from '../../api/model';

type UseTaskFormStateOptions = {
    initialState: Task;
    onSuccess?: (res: MutationResponse<Task>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: TaskMutationOptions;
};

export function useTaskFormState(options: UseTaskFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateTask = useUpdateTask(options.mutationOptions);
    const createTask = useCreateTask(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateTask.mutate(formState.updates, opts);
        }
        return createTask.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createTask,
        updateTask,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createTask.isLoading || updateTask.isLoading;

    return {
        ...formState,
        save,
        createMutation: createTask,
        updateMutation: updateTask,
        isLoading,
    } as const;
}
