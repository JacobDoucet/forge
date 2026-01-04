// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { User } from '../../model/user-model';
import { useFormState } from './common';
import {
    UserMutationOptions,
    useCreateUser,
    useUpdateUser,
} from '../tanstack-query/user-queries';
import { MutationResponse } from '../../api/model';

type UseUserFormStateOptions = {
    initialState: User;
    onSuccess?: (res: MutationResponse<User>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: UserMutationOptions;
};

export function useUserFormState(options: UseUserFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateUser = useUpdateUser(options.mutationOptions);
    const createUser = useCreateUser(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateUser.mutate(formState.updates, opts);
        }
        return createUser.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createUser,
        updateUser,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createUser.isLoading || updateUser.isLoading;

    return {
        ...formState,
        save,
        createMutation: createUser,
        updateMutation: updateUser,
        isLoading,
    } as const;
}
