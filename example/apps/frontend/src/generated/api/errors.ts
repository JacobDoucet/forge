// This file is auto-generated. DO NOT EDIT.

export type ErrorCode = 'UNKNOWN'
  | 'INVALID_REQUEST'
  | 'UNAUTHORIZED'
  | 'NOT_FOUND'
  | 'METHOD_NOT_ALLOWED'
  | 'ENTITY_ALREADY_EXISTS'
  | 'TASK_NOT_FOUND'
  | 'INVALID_TASK_STATUS'
  | 'PROJECT_NOT_FOUND'
  | 'UNEXPECTED';

function resolveErrorCode(errorCode: string): ErrorCode {
  errorCode = `${errorCode}`.trim();
  switch (errorCode as ErrorCode) {
    case 'INVALID_REQUEST':
      return errorCode as ErrorCode;
    case 'UNAUTHORIZED':
      return errorCode as ErrorCode;
    case 'NOT_FOUND':
      return errorCode as ErrorCode;
    case 'METHOD_NOT_ALLOWED':
      return errorCode as ErrorCode;
    case 'ENTITY_ALREADY_EXISTS':
      return errorCode as ErrorCode;
    case 'PROJECT_NOT_FOUND':
      return errorCode as ErrorCode;
    case 'UNEXPECTED':
      return errorCode as ErrorCode;
    case 'TASK_NOT_FOUND':
      return errorCode as ErrorCode;
    case 'INVALID_TASK_STATUS':
      return errorCode as ErrorCode;
    default:
      return 'UNKNOWN';
  }
}

export class ApiError extends Error {
  code: ErrorCode;
  constructor(errorCode: string) {
    super(errorCode);
    this.name = 'ApiError';
    this.code = resolveErrorCode(errorCode);
  }
}
