// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { User } from '../../model/user-model';
import { UserWithRefs } from '../../model/user-api';
import {
    searchUsers, SearchUsersParams,
    selectUserById, SelectUserByIdParams,
    selectUserByEmailIdx, SelectUserByEmailIdxParams,
    createUser, updateUser, deleteUser,
} from '../../api/user-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchUsersProps = Omit<SearchUsersParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<UserWithRefs>,
    ApiError,
    SelectManyResponse<UserWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchUsers(
    { queryKey, queryName, ...params }: UseSearchUsersProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchUsersParams['query']])}`
        );
        return ['searchUsers', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchUsers({ baseUrl, ...params }),
    });
}
type UseSelectUserByIdProps = Omit<SelectUserByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectUserByIdOptions = Omit<UseQueryOptions<
    UserWithRefs,
    ApiError,
    UserWithRefs,
    any[]
>, 'initialData'>;

export function useSelectUserById(
    { queryKey, queryName, ...params }: UseSelectUserByIdProps,
    queryOptions?: SelectUserByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectUserById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectUserById({ baseUrl, ...params }),
    });
}
type UseSelectUserByEmailIdxProps = Omit<SelectUserByEmailIdxParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectUserByEmailIdxOptions = Omit<UseQueryOptions<
    UserWithRefs,
    ApiError,
    UserWithRefs,
    any[]
>, 'initialData'>;

export function useSelectUserByEmailIdx(
    { queryKey, queryName, ...params }: UseSelectUserByEmailIdxProps,
    queryOptions?: SelectUserByEmailIdxOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectUserByEmailIdx', queryName, params.email];
    }, [queryKey, queryName, params.email]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectUserByEmailIdx({ baseUrl, ...params }),
    });
}

export type UserMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateUser(options: UserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<User>, ApiError, User>(async (user: User) => {
        const res = await createUser({ baseUrl, data: user });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateUser(options: UserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<User>, ApiError, User>(async (user: User) => {
        const res = await updateUser({ baseUrl, data: user });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteUser(options: UserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteUser({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
