// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Task } from '../../model/task-model';
import { TaskWithRefs } from '../../model/task-api';
import {
    searchTasks, SearchTasksParams,
    selectTaskById, SelectTaskByIdParams,
    createTask, updateTask, deleteTask,
} from '../../api/task-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchTasksProps = Omit<SearchTasksParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<TaskWithRefs>,
    ApiError,
    SelectManyResponse<TaskWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchTasks(
    { queryKey, queryName, ...params }: UseSearchTasksProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchTasksParams['query']])}`
        );
        return ['searchTasks', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchTasks({ baseUrl, ...params }),
    });
}
type UseSelectTaskByIdProps = Omit<SelectTaskByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectTaskByIdOptions = Omit<UseQueryOptions<
    TaskWithRefs,
    ApiError,
    TaskWithRefs,
    any[]
>, 'initialData'>;

export function useSelectTaskById(
    { queryKey, queryName, ...params }: UseSelectTaskByIdProps,
    queryOptions?: SelectTaskByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectTaskById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectTaskById({ baseUrl, ...params }),
    });
}

export type TaskMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateTask(options: TaskMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Task>, ApiError, Task>(async (task: Task) => {
        const res = await createTask({ baseUrl, data: task });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateTask(options: TaskMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Task>, ApiError, Task>(async (task: Task) => {
        const res = await updateTask({ baseUrl, data: task });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteTask(options: TaskMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteTask({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
