// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Project } from '../../model/project-model';
import { ProjectWithRefs } from '../../model/project-api';
import {
    searchProjects, SearchProjectsParams,
    selectProjectById, SelectProjectByIdParams,
    createProject, updateProject, deleteProject,
} from '../../api/project-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchProjectsProps = Omit<SearchProjectsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<ProjectWithRefs>,
    ApiError,
    SelectManyResponse<ProjectWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchProjects(
    { queryKey, queryName, ...params }: UseSearchProjectsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchProjectsParams['query']])}`
        );
        return ['searchProjects', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchProjects({ baseUrl, ...params }),
    });
}
type UseSelectProjectByIdProps = Omit<SelectProjectByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectProjectByIdOptions = Omit<UseQueryOptions<
    ProjectWithRefs,
    ApiError,
    ProjectWithRefs,
    any[]
>, 'initialData'>;

export function useSelectProjectById(
    { queryKey, queryName, ...params }: UseSelectProjectByIdProps,
    queryOptions?: SelectProjectByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectProjectById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectProjectById({ baseUrl, ...params }),
    });
}

export type ProjectMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateProject(options: ProjectMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Project>, ApiError, Project>(async (project: Project) => {
        const res = await createProject({ baseUrl, data: project });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateProject(options: ProjectMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Project>, ApiError, Project>(async (project: Project) => {
        const res = await updateProject({ baseUrl, data: project });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteProject(options: ProjectMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteProject({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
