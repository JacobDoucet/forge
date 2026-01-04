// This file is auto-generated. DO NOT EDIT.

import { Model } from './model-enum';

export type EventSubjectSearchQuery = {
    // subjectId (string) search options
    subjectIdEq?: string;
    subjectIdNe?: string;
    subjectIdGt?: string;
    subjectIdGte?: string;
    subjectIdLt?: string;
    subjectIdLte?: string;
    subjectIdIn?: string[];
    subjectIdNin?: string[];
    subjectIdExists?: boolean;
    subjectIdLike?: string;
    subjectIdNlike?: string;
    // subjectType (Model) search options
    subjectTypeEq?: Model;
    subjectTypeNe?: Model;
    subjectTypeGt?: Model;
    subjectTypeGte?: Model;
    subjectTypeLt?: Model;
    subjectTypeLte?: Model;
    subjectTypeIn?: Model[];
    subjectTypeNin?: Model[];
    subjectTypeExists?: boolean;
}
