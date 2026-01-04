// This file is auto-generated. DO NOT EDIT.

import { ProjectSearchQuery, ProjectWithRefs, ProjectWithRefsProjection } from '../model/project-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Project, ProjectSortParams } from '../model/project-model';

export type SearchProjectsParams = {
    baseUrl: string;
    query: ProjectSearchQuery;
    sort?: ProjectSortParams;
    projection?: ProjectWithRefsProjection;
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

export function searchProjects(params: SearchProjectsParams): Promise<SelectManyResponse<ProjectWithRefs>> {
    return fetch(`${params.baseUrl}/projects/search`, {
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
            const err = await newApiError(response, 'Failed to search Project');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectProjectByIdParams = {
    baseUrl: string;
    id: string;
    projection?: ProjectWithRefsProjection;
}

export function selectProjectById(params: SelectProjectByIdParams): Promise<ProjectWithRefs> {
    return fetch(`${params.baseUrl}/projects/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Project');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveProjectParams = {
    baseUrl: string;
    data: Project;
}

export function createProject(params: SaveProjectParams): Promise<MutationResponse<Project>> {
    return fetch(`${params.baseUrl}/projects/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Project');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateProject(params: SaveProjectParams): Promise<MutationResponse<Project>> {
    return fetch(`${params.baseUrl}/projects/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Project');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteProjectParams = {
    baseUrl: string;
    id: string;
}

export function deleteProject({ baseUrl, id }: DeleteProjectParams): Promise<void> {
    return fetch(`${baseUrl}/projects/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Project');
            return Promise.reject(err);
        }
        return;
    });
}
