package model_template_go

import _ "embed"

//go:embed obj__model/search_types.tmpl
var objSearchTypesTemplate string

//go:embed obj__model/search_types_transform.tmpl
var objSearchTypesTransformTemplate string

//go:embed obj__model/sort.go.tmpl
var objSortGoTemplate string

//go:embed obj__model/sort_transform.go.tmpl
var objSortTransformGoTemplate string
