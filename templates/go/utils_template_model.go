package model_template_go

import (
	"strings"

	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

type fieldTransformParams struct {
	Field              types.Field
	Obj                types.Object
	Registry           types.Registry
	ReturnVarName      string
	ReceiverVarName    string
	RefStructName      string
	IdTransformType    IdTransformType
	FieldName          string
	TransformFuncArgs  string
	SkipAssign         bool
	Level              int
	ResolveFieldType   func(field types.Field) string
	BuildIterVar       func(field types.Field, n int) string
	BuildKeyVar        func(field types.Field, n int) string
	BuildElemVar       func(field types.Field, n int) string
	IsPointerToValue   bool
	IsValueToPointer   bool
	IsPointerToPointer bool
}

func valueFieldTransform(
	params fieldTransformParams,
) string {
	elemVar := params.BuildElemVar(params.Field, params.Level)
	nextElemVar := params.BuildElemVar(params.Field, params.Level+1)

	transformElemVarAssignment := func(s string) string {
		if params.Field.IsRef() && params.IsPointerToPointer && params.IdTransformType == "" {
			return s
		}
		if params.Field.IsRef() && params.IsPointerToPointer {
			return "&" + s
		}
		if params.Field.IsRef() && params.IdTransformType == "" && params.IsPointerToValue {
			return "*" + s
		}
		if params.Field.IsPrimitive() && params.IsPointerToValue {
			return "*" + s
		}
		if params.Field.IsEnum(params.Registry) && params.IsPointerToValue {
			return "*" + s
		}
		_, isList := params.Field.ParseList()
		if isList && params.IsPointerToPointer {
			return "&" + s
		}
		rootType, _, _ := params.Field.ResolveRootType(params.Registry)
		if _, ok := params.Registry.Get(rootType); ok && params.IsPointerToPointer {
			return "&" + s
		}
		if params.IsValueToPointer {
			return "&" + s
		}
		return s
	}

	formatIterator := func() string {
		if params.IsPointerToValue || params.IsPointerToPointer {
			return "*" + params.FieldName
		}
		return params.FieldName
	}

	concatResult := func(lines []string) string {
		if params.SkipAssign {
			return strings.Join(lines[0:len(lines)-1], "\n")
		}
		return strings.Join(lines, "\n")
	}

	elemType, isList := params.Field.ParseList()
	if isList {
		iterVar := params.BuildIterVar(params.Field, params.Level)
		startLines := []string{
			elemVar + " := make(" + params.ResolveFieldType(params.Field) + ", 0)",
			"for _, " + iterVar + " := range " + formatIterator() + " {",
		}

		transformField := types.Field{Type: elemType, Name: params.Field.Name}
		middleLines := []string{valueFieldTransform(fieldTransformParams{
			Field:             transformField,
			Obj:               params.Obj,
			Registry:          params.Registry,
			ReceiverVarName:   params.ReceiverVarName,
			ReturnVarName:     params.ReturnVarName,
			RefStructName:     params.RefStructName,
			IdTransformType:   params.IdTransformType,
			FieldName:         iterVar,
			TransformFuncArgs: params.TransformFuncArgs,
			SkipAssign:        true,
			Level:             params.Level + 1,
			ResolveFieldType:  params.ResolveFieldType,
			BuildElemVar:      params.BuildElemVar,
			BuildIterVar:      params.BuildIterVar,
			BuildKeyVar:       params.BuildKeyVar,
		})}

		endLines := []string{
			elemVar + " = append(" + elemVar + ", " + nextElemVar + ")",
			"}",
			params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
		}

		return concatResult(append(startLines, append(middleLines, endLines...)...))
	}
	recordType, isKeyVal := params.Field.ParseKeyVal()
	if isKeyVal {
		iterVar := params.BuildIterVar(params.Field, params.Level)
		keyVar := params.BuildKeyVar(params.Field, params.Level)

		startLines := []string{
			elemVar + " := make(" + params.ResolveFieldType(params.Field) + ")",
			"for " + keyVar + ", " + iterVar + " := range " + formatIterator() + " {",
		}

		transformField := types.Field{Type: recordType, Name: params.Field.Name}
		middleLines := []string{valueFieldTransform(fieldTransformParams{
			Field:             transformField,
			Obj:               params.Obj,
			Registry:          params.Registry,
			ReceiverVarName:   params.ReceiverVarName,
			ReturnVarName:     params.ReturnVarName,
			RefStructName:     params.RefStructName,
			IdTransformType:   params.IdTransformType,
			FieldName:         iterVar,
			TransformFuncArgs: params.TransformFuncArgs,
			SkipAssign:        true,
			Level:             params.Level + 1,
			ResolveFieldType:  params.ResolveFieldType,
			BuildElemVar:      params.BuildElemVar,
			BuildIterVar:      params.BuildIterVar,
			BuildKeyVar:       params.BuildKeyVar,
		})}

		endLines := []string{
			elemVar + "[" + keyVar + "] = " + nextElemVar,
			"}",
			params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
		}

		return concatResult(append(startLines, append(middleLines, endLines...)...))
	}

	if params.Field.IsPrimitive() || params.Field.IsEnum(params.Registry) {
		return concatResult([]string{
			elemVar + " := " + params.FieldName,
			params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
		})
	}

	if params.Field.IsRef() {
		return concatResult(refValueTransform(params, elemVar, transformElemVarAssignment))
	}

	return concatResult([]string{
		elemVar + ", err := " + params.FieldName + ".To" + params.RefStructName + "(" + params.TransformFuncArgs + ")",
		"if err != nil {",
		"return " + params.ReturnVarName + ", err",
		"}",
		params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
	})
}

type IdTransformType string

const (
	IdTransformMongoToString IdTransformType = "mongo-to-string"
	IdTransformStringToMongo IdTransformType = "string-to-mongo"
)

func refValueTransform(
	params fieldTransformParams,
	elemVar string,
	transformElemVarAssignment func(string) string,
) []string {
	if params.IdTransformType == IdTransformStringToMongo {
		return transformStringToObjectId(transformStringToObjectIdParams{
			ShouldDerefFromVar:         params.IsPointerToPointer || params.IsPointerToValue,
			ElemVar:                    elemVar,
			FieldName:                  params.FieldName,
			ReturnVarName:              params.ReturnVarName,
			Field:                      params.Field,
			TransformElemVarAssignment: transformElemVarAssignment,
		})
	}

	if params.IdTransformType == IdTransformMongoToString {
		return []string{
			elemVar + " := " + params.FieldName + ".Hex()",
			params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
		}
	}

	return []string{
		elemVar + " := " + params.FieldName,
		params.ReturnVarName + "." + utils.UCC(params.Field.Name) + " = " + transformElemVarAssignment(elemVar),
	}
}

type transformStringToObjectIdParams struct {
	ShouldDerefFromVar         bool
	ShouldDerefToVar           bool
	ElemVar                    string
	FieldName                  string
	ReturnVarName              string
	Field                      types.Field
	TransformElemVarAssignment func(string) string
	TransformReturnField       func(string) string
}

func transformStringToObjectId(params transformStringToObjectIdParams) []string {
	if params.TransformElemVarAssignment == nil {
		params.TransformElemVarAssignment = func(s string) string {
			return s
		}
	}
	if params.TransformReturnField == nil {
		params.TransformReturnField = utils.UCC
	}
	derefFrom := ""
	if params.ShouldDerefFromVar {
		derefFrom = "*"
	}
	derefTo := ""
	if params.ShouldDerefToVar {
		derefTo = "*"
	}
	return []string{
		params.ElemVar + ", err := " + "primitive.ObjectIDFromHex(" + derefFrom + params.FieldName + ")",
		"if err != nil {",
		"return " + params.ReturnVarName + ", errors.Join(errors.New(\"invalid " + params.FieldName + "\"), err)",
		"}",
		derefTo + params.ReturnVarName + "." + params.TransformReturnField(params.Field.Name) + " = " + params.TransformElemVarAssignment(params.ElemVar),
	}
}
