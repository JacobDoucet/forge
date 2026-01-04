package types

type AbacImpl struct {
	Name  string `yaml:"name"`
	Field string `yaml:"field"`
}

type PermissionsDef struct {
	Abac       []PermissionsAbacDef `yaml:"abac"`
	Rbac       []PermissionsRbacDef `yaml:"rbac"`
	RoleGroups []RoleGroupDef       `yaml:"roleGroups"`
}

type PermissionsAbacDef struct {
	Name      string `yaml:"name"`
	FieldType string `yaml:"fieldType"`
}

type PermissionsRbacDef struct {
	Name              string   `yaml:"name"`
	Abac              []string `yaml:"abac"`
	CustomPermissions []string `yaml:"customPermissions"`
	Extends           []string `yaml:"extends"`
}

type RoleGroupDef struct {
	Name  string   `yaml:"name"`
	Roles []string `yaml:"roles"`
}
