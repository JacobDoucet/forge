package model_template_kotlin

import (
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

func GetKotlinModelFilename(obj types.Object) string {
	return utils.UCC(obj.Name)
}
