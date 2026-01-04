// This file is auto-generated. DO NOT EDIT.

import { Task, TaskProjection } from './task-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { TaskCommentSearchQuery } from './task-comment-api';
import { TaskPriority } from './task-priority-enum';
import { TaskStatus } from './task-status-enum';

export type TaskWithRefs = {
    task: Task;
}

export type TaskWithRefsProjection = TaskProjection & {
}

export type SelectTaskByIdQuery = {
    id: string;
}

export type TaskSearchQuery = {
    // id (Ref<Task>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // assigneeId (string) search options
    assigneeIdEq?: string;
    assigneeIdNe?: string;
    assigneeIdGt?: string;
    assigneeIdGte?: string;
    assigneeIdLt?: string;
    assigneeIdLte?: string;
    assigneeIdIn?: string[];
    assigneeIdNin?: string[];
    assigneeIdExists?: boolean;
    assigneeIdLike?: string;
    assigneeIdNlike?: string;
    // comments (List<TaskComment>) search options
    comments?: TaskCommentSearchQuery;
    commentsEmpty?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // description (string) search options
    descriptionEq?: string;
    descriptionNe?: string;
    descriptionGt?: string;
    descriptionGte?: string;
    descriptionLt?: string;
    descriptionLte?: string;
    descriptionIn?: string[];
    descriptionNin?: string[];
    descriptionExists?: boolean;
    descriptionLike?: string;
    descriptionNlike?: string;
    // dueDate (timestamp) search options
    dueDateEq?: string;
    dueDateNe?: string;
    dueDateGt?: string;
    dueDateGte?: string;
    dueDateLt?: string;
    dueDateLte?: string;
    dueDateIn?: string[];
    dueDateNin?: string[];
    dueDateExists?: boolean;
    // priority (TaskPriority) search options
    priorityEq?: TaskPriority;
    priorityNe?: TaskPriority;
    priorityGt?: TaskPriority;
    priorityGte?: TaskPriority;
    priorityLt?: TaskPriority;
    priorityLte?: TaskPriority;
    priorityIn?: TaskPriority[];
    priorityNin?: TaskPriority[];
    priorityExists?: boolean;
    // status (TaskStatus) search options
    statusEq?: TaskStatus;
    statusNe?: TaskStatus;
    statusGt?: TaskStatus;
    statusGte?: TaskStatus;
    statusLt?: TaskStatus;
    statusLte?: TaskStatus;
    statusIn?: TaskStatus[];
    statusNin?: TaskStatus[];
    statusExists?: boolean;
    // tags (List<string>) search options
    tagsEq?: string;
    tagsNe?: string;
    tagsGt?: string;
    tagsGte?: string;
    tagsLt?: string;
    tagsLte?: string;
    tagsIn?: string[];
    tagsNin?: string[];
    tagsExists?: boolean;
    tagsLike?: string;
    tagsNlike?: string;
    tagsEmpty?: boolean;
    // title (string) search options
    titleEq?: string;
    titleNe?: string;
    titleGt?: string;
    titleGte?: string;
    titleLt?: string;
    titleLte?: string;
    titleIn?: string[];
    titleNin?: string[];
    titleExists?: boolean;
    titleLike?: string;
    titleNlike?: string;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
}
