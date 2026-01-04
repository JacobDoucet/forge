// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Project } from '../../model/project-model';
import { useFormState } from './common';
import {
    ProjectMutationOptions,
    useCreateProject,
    useUpdateProject,
} from '../tanstack-query/project-queries';
import { MutationResponse } from '../../api/model';

type UseProjectFormStateOptions = {
    initialState: Project;
    onSuccess?: (res: MutationResponse<Project>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: ProjectMutationOptions;
};

export function useProjectFormState(options: UseProjectFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateProject = useUpdateProject(options.mutationOptions);
    const createProject = useCreateProject(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateProject.mutate(formState.updates, opts);
        }
        return createProject.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createProject,
        updateProject,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createProject.isLoading || updateProject.isLoading;

    return {
        ...formState,
        save,
        createMutation: createProject,
        updateMutation: updateProject,
        isLoading,
    } as const;
}
