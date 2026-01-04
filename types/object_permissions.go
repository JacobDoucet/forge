package types

import "strings"

type ObjectPermissions struct {
	Read  []ObjectPermissionsDef `yaml:"read"`
	Write []ObjectPermissionsDef `yaml:"write"`
}

type ObjectPermissionsDef struct {
	Rbac string   `yaml:"rbac"`
	Abac []string `yaml:"abac"`
}

func (pd *ObjectPermissionsDef) GetName() string {
	return pd.Rbac + strings.Join(pd.Abac, "") + "Permission"
}

type ObjectAbacPermissionsImpl struct {
	Name      string                  `yaml:"name"`
	Interface ObjectAbacInterfaceImpl `yaml:"interface"`
}

type ObjectAbacInterfaceImpl struct {
	Name  string `yaml:"name"`
	Field string `yaml:"field"`
}
