// This file is auto-generated. DO NOT EDIT.

import { TaskSearchQuery, TaskWithRefs, TaskWithRefsProjection } from '../model/task-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Task, TaskSortParams } from '../model/task-model';

export type SearchTasksParams = {
    baseUrl: string;
    query: TaskSearchQuery;
    sort?: TaskSortParams;
    projection?: TaskWithRefsProjection;
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

export function searchTasks(params: SearchTasksParams): Promise<SelectManyResponse<TaskWithRefs>> {
    return fetch(`${params.baseUrl}/tasks/search`, {
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
            const err = await newApiError(response, 'Failed to search Task');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectTaskByIdParams = {
    baseUrl: string;
    id: string;
    projection?: TaskWithRefsProjection;
}

export function selectTaskById(params: SelectTaskByIdParams): Promise<TaskWithRefs> {
    return fetch(`${params.baseUrl}/tasks/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Task');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveTaskParams = {
    baseUrl: string;
    data: Task;
}

export function createTask(params: SaveTaskParams): Promise<MutationResponse<Task>> {
    return fetch(`${params.baseUrl}/tasks/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Task');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateTask(params: SaveTaskParams): Promise<MutationResponse<Task>> {
    return fetch(`${params.baseUrl}/tasks/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Task');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteTaskParams = {
    baseUrl: string;
    id: string;
}

export function deleteTask({ baseUrl, id }: DeleteTaskParams): Promise<void> {
    return fetch(`${baseUrl}/tasks/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Task');
            return Promise.reject(err);
        }
        return;
    });
}
