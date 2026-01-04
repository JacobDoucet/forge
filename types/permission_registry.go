package types

import (
	"fmt"
	"sort"
)

type PermissionRegistry interface {
	RegisterRbac(permission PermissionsRbacDef) error
	RegisterRoleGroup(roleGroup RoleGroupDef) error
	RegisterAbac(permission PermissionsAbacDef) error
	ListRbac() []PermissionsRbacDef
	ListRoleGroups() []RoleGroupDef
	ListAbac() []PermissionsAbacDef
	ListCustomPermissions() []RegisteredCustomPermission
	HasAbac() bool
	HasRbac() bool
	GetRbacPermission(name string) (PermissionsRbacDef, bool)
	GetAbacPermission(name string) (PermissionsAbacDef, bool)
	ValidateObjectPermissions(name string, permissions ObjectPermissions) error
	RegisterExtensions(permission PermissionsRbacDef) error
	GetExtendedRoles(role string) []string
}

func newPermissionRegistry() PermissionRegistry {
	abac := make(map[string]PermissionsAbacDef)
	abac["ActorId"] = PermissionsAbacDef{
		Name:      "ActorId",
		FieldType: "string",
	}
	rbac := make(map[string]PermissionsRbacDef)
	rbac["Super"] = PermissionsRbacDef{
		Name: "Super",
	}
	return &permissionRegistry{
		rbac:              rbac,
		abac:              abac,
		customPermissions: make(map[string]RegisteredCustomPermission),
		roleGroups:        make(map[string]RoleGroupDef),
		roleExtendedBy:    make(map[string][]string),
	}
}

type permissionRegistry struct {
	rbac              map[string]PermissionsRbacDef
	abac              map[string]PermissionsAbacDef
	customPermissions map[string]RegisteredCustomPermission
	roleGroups        map[string]RoleGroupDef
	roleExtendedBy    map[string][]string
}

type RegisteredCustomPermission struct {
	Name  string
	Roles []string
}

func (r *permissionRegistry) RegisterRbac(permission PermissionsRbacDef) error {
	if _, ok := r.rbac[permission.Name]; ok {
		return fmt.Errorf("permission %s already registered", permission.Name)
	}
	r.rbac[permission.Name] = permission

	for _, customPermission := range permission.CustomPermissions {
		r.addCustomPermission(customPermission, permission.Name)
	}

	return nil
}

func (r *permissionRegistry) RegisterRoleGroup(roleGroup RoleGroupDef) error {
	if _, ok := r.roleGroups[roleGroup.Name]; ok {
		return fmt.Errorf("role group %s already registered", roleGroup.Name)
	}
	for _, role := range roleGroup.Roles {
		if _, ok := r.rbac[role]; !ok {
			return fmt.Errorf("role %s in role group %s is not registered", role, roleGroup.Name)
		}
	}
	r.roleGroups[roleGroup.Name] = roleGroup
	return nil
}

func (r *permissionRegistry) ListRbac() []PermissionsRbacDef {
	var permissions []PermissionsRbacDef
	for _, permission := range r.rbac {
		permissions = append(permissions, permission)
	}
	sort.Slice(permissions, func(i, j int) bool {
		return permissions[i].Name < permissions[j].Name
	})
	return permissions
}

func (r *permissionRegistry) ListRoleGroups() []RoleGroupDef {
	var roleGroups []RoleGroupDef
	for _, roleGroup := range r.roleGroups {
		roleGroups = append(roleGroups, roleGroup)
	}
	sort.Slice(roleGroups, func(i, j int) bool {
		return roleGroups[i].Name < roleGroups[j].Name
	})
	return roleGroups
}

func (r *permissionRegistry) RegisterAbac(permission PermissionsAbacDef) error {
	if _, ok := r.abac[permission.Name]; ok {
		return fmt.Errorf("permission %s already registered", permission.Name)
	}
	r.abac[permission.Name] = permission
	return nil
}

func (r *permissionRegistry) ListAbac() []PermissionsAbacDef {
	var permissions []PermissionsAbacDef
	for _, permission := range r.abac {
		permissions = append(permissions, permission)
	}
	sort.Slice(permissions, func(i, j int) bool {
		return permissions[i].Name < permissions[j].Name
	})
	return permissions
}

func (r *permissionRegistry) ListCustomPermissions() []RegisteredCustomPermission {
	var permissions []RegisteredCustomPermission
	for _, permission := range r.customPermissions {
		permissions = append(permissions, permission)
	}
	sort.Slice(permissions, func(i, j int) bool {
		return permissions[i].Name < permissions[j].Name
	})
	return permissions
}

func (r *permissionRegistry) GetRbacPermission(name string) (PermissionsRbacDef, bool) {
	permission, ok := r.rbac[name]
	return permission, ok
}

func (r *permissionRegistry) GetAbacPermission(name string) (PermissionsAbacDef, bool) {
	permission, ok := r.abac[name]
	return permission, ok
}

func (r *permissionRegistry) HasAbac() bool {
	return len(r.abac) > 0
}

func (r *permissionRegistry) HasRbac() bool {
	return len(r.rbac) > 0
}

func (r *permissionRegistry) ValidateObjectPermissions(name string, permissions ObjectPermissions) error {
	for _, permission := range permissions.Read {
		if permission.Rbac == "" && len(permission.Abac) == 0 {
			return fmt.Errorf("object has empty permission")
		}
		if _, ok := r.rbac[permission.Rbac]; permission.Rbac != "" && !ok {
			return fmt.Errorf("object has rbac permission %s, but the permission is not defined", permission.Rbac)
		}
		for _, abacPerm := range permission.Abac {
			if _, ok := r.abac[abacPerm]; abacPerm != "" && !ok {
				return fmt.Errorf("object has abac permission %s, but the permission is not defined", permission.Abac)
			}
		}
		if permission.Rbac == "" {
			continue
		}
		//r.addCustomPermission("ReadSome"+name+"s", permission.Rbac)
		//if len(permission.Abac) == 0 {
		//	r.addCustomPermission("ReadAll"+name+"s", permission.Rbac)
		//}
	}
	for _, permission := range permissions.Write {
		if permission.Rbac == "" && len(permission.Abac) == 0 {
			return fmt.Errorf("object has empty permission")
		}
		if _, ok := r.rbac[permission.Rbac]; permission.Rbac != "" && !ok {
			return fmt.Errorf("object has rbac permission %s, but the permission is not defined", permission.Rbac)
		}
		for _, abacPerm := range permission.Abac {
			if _, ok := r.abac[abacPerm]; abacPerm != "" && !ok {
				return fmt.Errorf("object has abac permission %s, but the permission is not defined", permission.Abac)
			}
		}
		if permission.Rbac == "" {
			continue
		}
		//r.addCustomPermission("WriteSome"+name+"s", permission.Rbac)
		//if len(permission.Abac) == 0 {
		//	r.addCustomPermission("WriteAll"+name+"s", permission.Rbac)
		//}
	}
	return nil
}

func (r *permissionRegistry) addCustomPermission(permissionName string, roleName string) {
	fmt.Println("Adding custom permission", permissionName, roleName)
	rcp := r.customPermissions[permissionName]
	rcp.Name = permissionName
	for _, role := range rcp.Roles {
		if role == roleName {
			return
		}
	}
	rcp.Roles = append(rcp.Roles, roleName)
	r.customPermissions[permissionName] = rcp
}

func (r *permissionRegistry) RegisterExtensions(permission PermissionsRbacDef) error {
	if len(permission.Extends) == 0 {
		return nil
	}
	for _, superRole := range permission.Extends {
		if _, ok := r.rbac[superRole]; !ok {
			return fmt.Errorf("role %s in extends of permission %s is not registered", superRole, permission.Name)
		}
		if _, ok := r.rbac[permission.Name]; !ok {
			return fmt.Errorf("cannot extend role %s, because it is not registered", permission.Name)
		}
		r.roleExtendedBy[superRole] = append(r.roleExtendedBy[superRole], permission.Name)
		fmt.Printf("Role %s extends %s\n", permission.Name, superRole)
	}
	return nil
}

func (r *permissionRegistry) GetExtendedRoles(role string) []string {
	extendedRoles, ok := r.roleExtendedBy[role]
	if !ok {
		return nil
	}
	sort.Strings(extendedRoles)
	return extendedRoles
}
