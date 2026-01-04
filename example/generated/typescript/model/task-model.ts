// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { TaskComment, TaskCommentProjection } from './task-comment-model';
import { TaskPriority } from './task-priority-enum';
import { TaskStatus } from './task-status-enum';

export type Task = {
  id?: string;
  assigneeId?: string;
  comments?: TaskComment[];
  created?: ActorTrace;
  description?: string;
  dueDate?: string;
  priority?: TaskPriority;
  status?: TaskStatus;
  tags?: string[];
  title?: string;
  updated?: ActorTrace;
}

export type TaskProjection = {
    id?: boolean;
    assigneeId?: boolean;
    comments?: boolean;
		commentsFields?: TaskCommentProjection;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    description?: boolean;
    dueDate?: boolean;
    priority?: boolean;
    status?: boolean;
    tags?: boolean;
    title?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
}

export type TaskSortParams = {
    assigneeId?: -1 | 1;
    createdAt?: -1 | 1;
    status?: -1 | 1;
    updatedAt?: -1 | 1;
}
