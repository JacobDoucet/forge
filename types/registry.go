package types

import (
	"errors"
	"fmt"
	"sort"

	"github.com/JacobDoucet/forge/utils"
)

type Registry interface {
	GetGoPkgRoot() string
	GetKotlinPkgRoot() string
	Get(name string) (Object, bool)
	RegisterObject(obj Object) error
	OverwriteObject(obj Object) error
	RegisterIfActor(obj Object) error
	ListActorTypes() []string
	RegisterPermissions(actorObj *Object, permission PermissionsDef) error
	HasRbacPermissions() bool
	HasAbacPermissions() bool
	RegisterEnum(enum Enum) error
	GetEnum(name string) (Enum, bool)
	ListEnums() []Enum
	BuildRefs() error
	ListObjects() []Object
	ListActors() []Object
	ListCollections() []Collection
	ListRbacPermissions() []PermissionsRbacDef
	ListRoleGroups() []RoleGroupDef
	ListAbacPermissions() []PermissionsAbacDef
	ListCustomPermissions() []RegisteredCustomPermission
	GetRbacPermission(name string) (PermissionsRbacDef, bool)
	GetAbacPermission(name string) (PermissionsAbacDef, bool)
	Validate() error
	ValidateObjectPermission(name string, permissions ObjectPermissions) error
	RegisterError(customError CustomError) error
	GetError(name string) (CustomError, bool)
	ListErrors() []CustomError
	RegisterEvent(event string) error
	GetEventObjects(permissions *ObjectPermissions) []Object
	GetEventEnums() []Enum
	GetExtendedRoles(role string) []string
	ApplyExtendedRolePermissions() error
}

func NewRegistry(goPkgRoot string, kotlinPkgRoot string) Registry {
	return &registry{
		pkgRoot:            goPkgRoot,
		kotlinPkgRoot:      kotlinPkgRoot,
		objects:            make(map[string]Object),
		actorTypes:         []string{},
		collectionRegistry: newCollectionRegistry(),
		endpointRegistry:   newEndpointRegistry(),
		permissionRegistry: newPermissionRegistry(),
		enumRegistry:       NewEnumRegistry(),
		errorRegistry:      NewErrorRegistry(),
		eventRegistry:      NewEventRegistry(),
	}
}

type registry struct {
	pkgRoot            string
	kotlinPkgRoot      string
	objects            map[string]Object
	actorTypes         []string
	collectionRegistry CollectionRegistry
	endpointRegistry   EndpointRegistry
	permissionRegistry PermissionRegistry
	enumRegistry       EnumRegistry
	errorRegistry      ErrorRegistry
	eventRegistry      EventRegistry
}

func (r *registry) GetGoPkgRoot() string {
	return r.pkgRoot
}

func (r *registry) GetKotlinPkgRoot() string {
	return r.kotlinPkgRoot
}

func (r *registry) Get(name string) (Object, bool) {
	obj, ok := r.objects[name]
	return obj, ok
}

func (r *registry) GetRbacPermission(name string) (PermissionsRbacDef, bool) {
	return r.permissionRegistry.GetRbacPermission(name)
}

func (r *registry) GetAbacPermission(name string) (PermissionsAbacDef, bool) {
	return r.permissionRegistry.GetAbacPermission(name)
}

func (r *registry) RegisterIfActor(obj Object) error {
	if err := obj.parseActor(); err != nil {
		return err
	}
	if obj.IsActor {
		fmt.Println("Registering actor type " + obj.Name)
		r.actorTypes = append(r.actorTypes, obj.Name)
	}
	return nil
}

func (r *registry) ListActorTypes() []string {
	return r.actorTypes
}

func (r *registry) RegisterObject(obj Object) error {
	fmt.Println("Registering object", obj.Name)
	var err error
	if _, ok := r.objects[obj.Name]; ok {
		err = fmt.Errorf("object %s already registered", obj.Name)
	}

	fieldMapErr := obj.buildFieldMap(r)
	if fieldMapErr != nil {
		err = errors.Join(err, fieldMapErr)
	}

	if actorErr := obj.parseActor(); actorErr != nil {
		err = errors.Join(err, actorErr)
	}

	if obj.IsActor {
		if actorFieldErr := obj.appendActorFields(); actorFieldErr != nil {
			err = errors.Join(err, actorFieldErr)
		}
	}

	r.objects[obj.Name] = obj

	for _, collection := range obj.Collection {
		if err := r.collectionRegistry.Register(collection); err != nil {
			return fmt.Errorf("object %s has invalid collection %s", obj.Name, collection.Name)
		}
	}

	if obj.HTTP.Endpoint != "" {
		if err := r.endpointRegistry.Register(obj.HTTP); err != nil {
			return fmt.Errorf("object %s has invalid endpoint %s", obj.Name, obj.HTTP.Endpoint)
		}
	}

	return err
}

func (r *registry) OverwriteObject(obj Object) error {
	r.objects[obj.Name] = obj
	return nil
}

func (r *registry) RegisterPermissions(actorRoleObj *Object, permission PermissionsDef) error {
	fmt.Println("Registering permissions", permission)
	for _, abac := range permission.Abac {
		if err := r.permissionRegistry.RegisterAbac(abac); err != nil {
			return err
		}

		actorRoleFieldName := utils.LCC(abac.Name)
		if _, ok := actorRoleObj.GetField(actorRoleFieldName); ok {
			continue
		}

		actorRoleObj.Fields = append(actorRoleObj.Fields, Field{
			Name: actorRoleFieldName,
			Type: "string",
		})
	}

	if len(permission.Rbac) == 0 {
		return nil
	}
	rbacEnum := Enum{
		Name: "Role",
		Type: FieldTypeString,
		Values: []string{
			"Super",
		},
	}
	for _, rbac := range permission.Rbac {
		if err := r.permissionRegistry.RegisterRbac(rbac); err != nil {
			return err
		}
		rbacEnum.Values = append(rbacEnum.Values, rbac.Name)
	}
	fmt.Println("Registering rbac enum", rbacEnum)
	err := r.RegisterEnum(rbacEnum)
	if err != nil {
		return errors.Join(errors.New("failed to register rbac enum"), err)
	}
	for _, rbac := range permission.Rbac {
		if err := r.permissionRegistry.RegisterExtensions(rbac); err != nil {
			return errors.Join(errors.New("failed to register rbac extensions"), err)
		}
	}

	for _, roleGroup := range permission.RoleGroups {
		if err := r.RegisterRoleGroup(roleGroup); err != nil {
			return errors.Join(errors.New("failed to register role group"), err)
		}
	}

	err = r.RegisterObject(Object{
		Name: "ActorTrace",
		Fields: []Field{
			{
				Name: "actorType",
				Type: string(FieldTypeString),
			},
			{
				Name: "actorName",
				Type: string(FieldTypeString),
			},
			{
				Name: "actorId",
				Type: string(FieldTypeString),
			},
			{
				Name: "at",
				Type: string(FieldTypeTimestamp),
			},
		},
	})
	if err != nil {
		return errors.Join(errors.New("failed to register ActorTrace object"), err)
	}
	return nil
}

func (r *registry) RegisterRoleGroup(roleGroup RoleGroupDef) error {
	fmt.Println("Registering role group", roleGroup.Name)
	return r.permissionRegistry.RegisterRoleGroup(roleGroup)
}

func (r *registry) HasAbacPermissions() bool {
	return r.permissionRegistry.HasAbac()
}

func (r *registry) HasRbacPermissions() bool {
	return r.permissionRegistry.HasRbac()
}

func (r *registry) ListObjects() []Object {
	var objects []Object
	for _, obj := range r.objects {
		objects = append(objects, obj)
	}
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].Name < objects[j].Name
	})
	return objects
}

func (r *registry) ListActors() []Object {
	var actors []Object
	for _, obj := range r.ListObjects() {
		if obj.IsActor {
			actors = append(actors, obj)
		}
	}
	return actors
}

func (r *registry) ListCollections() []Collection {
	return r.collectionRegistry.List()
}

func (r *registry) ListRbacPermissions() []PermissionsRbacDef {
	return r.permissionRegistry.ListRbac()
}

func (r *registry) ListRoleGroups() []RoleGroupDef {
	return r.permissionRegistry.ListRoleGroups()
}

func (r *registry) ListAbacPermissions() []PermissionsAbacDef {
	return r.permissionRegistry.ListAbac()
}

func (r *registry) ListCustomPermissions() []RegisteredCustomPermission {
	return r.permissionRegistry.ListCustomPermissions()
}

func (r *registry) RegisterEnum(enum Enum) error {
	fmt.Println("Registering enum", enum.Name)
	return r.enumRegistry.Register(enum)
}

func (r *registry) GetEnum(name string) (Enum, bool) {
	return r.enumRegistry.Get(name)
}

func (r *registry) ListEnums() []Enum {
	return r.enumRegistry.List()
}

func (r *registry) BuildRefs() error {
	for key, obj := range r.objects {
		err := obj.BuildRefs(r)
		if err != nil {
			return err
		}
		r.objects[key] = obj
	}
	return nil
}

func (r *registry) Validate() error {
	var err error
	for k, obj := range r.objects {
		if objErr := obj.Validate(r); objErr != nil {
			err = errors.Join(err, objErr)
		}
		r.objects[k] = obj
	}
	return err
}

func (r *registry) ValidateObjectPermission(name string, permissions ObjectPermissions) error {
	return r.permissionRegistry.ValidateObjectPermissions(name, permissions)
}

func (r *registry) RegisterError(customError CustomError) error {
	fmt.Println("Registering error", customError.Code)
	return r.errorRegistry.Register(customError)
}

func (r *registry) GetError(name string) (CustomError, bool) {
	return r.errorRegistry.Get(name)
}

func (r *registry) ListErrors() []CustomError {
	return r.errorRegistry.List()
}

func (r *registry) RegisterEvent(event string) error {
	fmt.Println("Registering event", event)
	return r.eventRegistry.RegisterEvent(event)
}

func (r *registry) GetEventObjects(permissions *ObjectPermissions) []Object {
	return r.eventRegistry.GetObjects(permissions)
}

func (r *registry) GetEventEnums() []Enum {
	return r.eventRegistry.GetEnums()
}

func (r *registry) GetExtendedRoles(role string) []string {
	return r.permissionRegistry.GetExtendedRoles(role)
}

func (r *registry) ApplyExtendedRolePermissions() error {
	for key, obj := range r.objects {
		if err := obj.ApplyExtendedPermissions(r); err != nil {
			return fmt.Errorf("failed to apply extended permissions for object %s: %w", key, err)
		}
		r.objects[key] = obj
	}
	return nil
}
