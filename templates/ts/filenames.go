package model_template_ts

import (
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

func GetTSModelFilename(obj types.Object) string {
	return utils.KC(obj.Name) + "-model"
}

func GetTSModelApiFilename(obj types.Object) string {
	return utils.KC(obj.Name) + "-api"
}

func GetTSPermissionsCanAccessFilename(obj types.Object) string {
	return utils.KC(obj.Name) + "-can-access"
}

func GetTSFormStateFilename(obj types.Object) string {
	return utils.KC(obj.Name)
}

func GetTSMaterialUITableFilename(obj types.Object) string {
	return utils.KC(obj.Name) + "-table"
}

func GetTSEnumFilename(enum types.Enum) string {
	return utils.KC(enum.Name) + "-enum"
}

func GetTSApiEndpointsFilename(obj types.Object) string {
	return utils.KC(obj.Name) + "-endpoints"
}
