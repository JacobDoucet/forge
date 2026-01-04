// This file is auto-generated. DO NOT EDIT.

import { UserSearchQuery, UserWithRefs, UserWithRefsProjection } from '../model/user-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { User, UserSortParams } from '../model/user-model';

export type SearchUsersParams = {
    baseUrl: string;
    query: UserSearchQuery;
    sort?: UserSortParams;
    projection?: UserWithRefsProjection;
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

export function searchUsers(params: SearchUsersParams): Promise<SelectManyResponse<UserWithRefs>> {
    return fetch(`${params.baseUrl}/users/search`, {
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
            const err = await newApiError(response, 'Failed to search User');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectUserByIdParams = {
    baseUrl: string;
    id: string;
    projection?: UserWithRefsProjection;
}

export function selectUserById(params: SelectUserByIdParams): Promise<UserWithRefs> {
    return fetch(`${params.baseUrl}/users/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select User');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectUserByEmailIdxParams = {
    baseUrl: string;
    email: string;
    projection?: UserWithRefsProjection;
}

export function selectUserByEmailIdx(params: SelectUserByEmailIdxParams): Promise<UserWithRefs> {
    return fetch(`${params.baseUrl}/users/email/${params.email}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select User');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveUserParams = {
    baseUrl: string;
    data: User;
}

export function createUser(params: SaveUserParams): Promise<MutationResponse<User>> {
    return fetch(`${params.baseUrl}/users/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create User');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateUser(params: SaveUserParams): Promise<MutationResponse<User>> {
    return fetch(`${params.baseUrl}/users/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update User');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteUserParams = {
    baseUrl: string;
    id: string;
}

export function deleteUser({ baseUrl, id }: DeleteUserParams): Promise<void> {
    return fetch(`${baseUrl}/users/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete User');
            return Promise.reject(err);
        }
        return;
    });
}
