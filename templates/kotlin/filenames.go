package model_template_kotlin

import (
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

func GetKotlinModelFilename(obj types.Object) string {
	return utils.UCC(obj.Name)
}
